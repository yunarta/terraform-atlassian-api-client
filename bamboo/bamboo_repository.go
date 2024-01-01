package bamboo

// RepositoryList is a struct that holds information about a list of repositories in Bamboo.
// It is especially useful in managing queries related to multiple repositories at once.
// Not all repositories will be returned when queried; the 'Start' parameter helps us by defining where to start indexing from,
// and 'MaxResult' defines the maximum number of repository results we want to fetch.
type RepositoryList struct {
	Start     int          `json:"start-index,omitempty"`
	MaxResult int          `json:"max-result,omitempty"`
	Results   []Repository `json:"searchResults,omitempty"`
}

// Repository is a struct that holds information related to a single Bamboo repository.
// Each repository is identified by a unique 'ID',
// 'Name' field is the name given to that repository and
// 'RssEnabled' indicates whether RSS feed option is enabled or not for that repository.
type Repository struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	RssEnabled bool   `json:"rssEnabled,omitempty"`
}
