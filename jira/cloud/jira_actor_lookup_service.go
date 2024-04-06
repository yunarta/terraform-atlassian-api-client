package cloud

import (
	"github.com/yunarta/terraform-atlassian-api-client/jira"
	"github.com/yunarta/terraform-atlassian-api-client/util"
	"sync"
)

// ActorLookupService prepare the accountId and username relationship ahead
// The lookup will be used for following use case
//   - Resolving the role user actor back to username
type ActorLookupService struct {
	actorService *ActorService

	sync.Mutex
	usernames  map[string]*jira.User
	accountIds map[string]*jira.User

	groupNames map[string]*jira.Group
	groupIds   map[string]*jira.Group
}

func NewActorLookupService(actorService *ActorService) *ActorLookupService {
	return &ActorLookupService{
		actorService: actorService,
		usernames:    make(map[string]*jira.User),
		accountIds:   make(map[string]*jira.User),
		groupNames:   make(map[string]*jira.Group),
		groupIds:     make(map[string]*jira.Group),
	}
}

func (service *ActorLookupService) RegisterAccountIds(accountId ...string) {
	users, _ := service.actorService.BulkGetUsers(accountId)
	for _, user := range users {
		service.syncUser(&user)
	}
}

func (service *ActorLookupService) RegisterUsernames(username ...string) {
	takeUsers := make([]string, 0)
	for _, user := range username {
		_, ok := service.usernames[user]
		if !ok {
			takeUsers = append(takeUsers, user)
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(takeUsers))

	for _, user := range takeUsers {
		go func(user string) {
			readUser, _ := service.actorService.ReadUser(user)
			if readUser != nil {
				service.syncUser(readUser)
			}
			defer wg.Done()
		}(user)
	}

	wg.Wait()
}

func (service *ActorLookupService) FindUser(username string) *jira.User {
	user, ok := service.usernames[username]
	if !ok {
		// the user is not in the system, so we fetch manually
		readUser, _ := service.actorService.ReadUser(username)
		if readUser != nil {
			service.syncUser(readUser)
			return readUser
		} else {
			return nil
		}
	} else {
		return user
	}
}

func (service *ActorLookupService) FindUserById(accountId string) *jira.User {
	user, ok := service.accountIds[accountId]
	if !ok {
		// the user is not in the system, so we fetch manually
		users, _ := service.actorService.BulkGetUsers([]string{accountId})
		var foundUser *jira.User
		for _, user := range users {
			service.syncUser(&user)
			if user.AccountID == accountId {
				aCopy := user
				foundUser = &aCopy
			}
		}

		return foundUser
	} else {
		return user
	}
}

func (service *ActorLookupService) syncUser(user *jira.User) {
	service.Mutex.Lock()
	username := util.CoalesceString(user.EmailAddress, user.DisplayName)

	var insert = user
	service.usernames[username] = insert
	service.accountIds[user.AccountID] = insert

	defer service.Mutex.Unlock()
}

func (service *ActorLookupService) RegisterGroupIds(groupId ...string) {
	if len(groupId) == 0 {
		return
	}

	users, _ := service.actorService.BulkGetGroupsById(groupId)
	for _, user := range users {
		service.syncGroup(&user)
	}
}

func (service *ActorLookupService) RegisterGroupNames(groupName ...string) {
	if len(groupName) == 0 {
		return
	}

	groups, err := service.actorService.BulkGetGroupsByName(groupName)
	if err != nil {
		return
	}

	for _, group := range groups {
		service.syncGroup(&group)
	}
}

func (service *ActorLookupService) FindGroup(groupName string) *jira.Group {
	group, ok := service.groupNames[groupName]
	if !ok {
		// the user is not in the system, so we fetch manually
		readGroup, _ := service.actorService.ReadGroup(groupName)
		if readGroup != nil {
			service.syncGroup(readGroup)
			return readGroup
		} else {
			return nil
		}
	} else {
		return group
	}
}

func (service *ActorLookupService) FindGroupById(groupId string) *jira.Group {
	group, ok := service.groupIds[groupId]
	if !ok {
		// the user is not in the system, so we fetch manually
		groups, _ := service.actorService.BulkGetGroupsById([]string{groupId})
		var foundGroup *jira.Group
		for _, group := range groups {
			service.syncGroup(&group)
			if group.GroupId == groupId {
				aCopy := group
				foundGroup = &aCopy
			}
		}

		return foundGroup
	} else {
		return group
	}
}

func (service *ActorLookupService) syncGroup(group *jira.Group) {
	var insert = group
	service.groupNames[group.Name] = insert
	service.groupIds[group.GroupId] = insert
}
