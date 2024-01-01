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
	usernames  map[string]string
	accountIds map[string]string

	groupNames map[string]string
	groupIds   map[string]string
}

func NewActorLookupService(actorService *ActorService) *ActorLookupService {
	return &ActorLookupService{
		actorService: actorService,
		usernames:    make(map[string]string),
		accountIds:   make(map[string]string),
		groupNames:   make(map[string]string),
		groupIds:     make(map[string]string),
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
			readUser, err := service.actorService.ReadUser(user)
			if err == nil {
				service.syncUser(readUser)
			}
			defer wg.Done()
		}(user)
	}

	wg.Wait()
}

func (service *ActorLookupService) FindUser(username string) string {
	accountId, ok := service.usernames[username]
	if !ok {
		// the user is not in the system, so we fetch manually
		readUser, err := service.actorService.ReadUser(username)
		if err == nil {
			service.syncUser(readUser)
		}

		return readUser.AccountID
	} else {
		return accountId
	}
}

func (service *ActorLookupService) syncUser(user *jira.User) {
	service.Mutex.Lock()
	username := util.CoalesceString(user.EmailAddress, user.DisplayName)

	service.usernames[username] = user.AccountID
	service.accountIds[user.AccountID] = username

	defer service.Mutex.Unlock()
}

func (service *ActorLookupService) RegisterGroupIds(accountId ...string) {
	if len(accountId) == 0 {
		return
	}

	users, _ := service.actorService.BulkGetGroupsById(accountId)
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

func (service *ActorLookupService) FindGroup(groupName string) string {
	groupId, ok := service.groupNames[groupName]
	if !ok {
		// the user is not in the system, so we fetch manually
		readUser, err := service.actorService.ReadGroup(groupName)
		if err == nil {
			service.syncGroup(readUser)
		}

		return readUser.GroupId
	} else {
		return groupId
	}
}

func (service *ActorLookupService) syncGroup(group *jira.Group) {
	service.groupNames[group.Name] = group.GroupId
	service.groupIds[group.GroupId] = group.Name
}
