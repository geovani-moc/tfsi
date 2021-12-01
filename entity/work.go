package entity

type Work struct {
	Description   string   `json:"description"`
	Title         string   `json:"title"`
	Covers        []int    `json:"covers"`
	SubjectPlaces []string `json:"subject_places"`
	Subjects      []string `json:"subjects"`
	SubjectPeople []string `json:"subject_people"`
	Key           string   `json:"key"`
	Authors       []struct {
		Author struct {
			Key string `json:"key"`
		} `json:"author"`
		Type struct {
			Key string `json:"key"`
		} `json:"type"`
	} `json:"authors"`
	SubjectTimes []string `json:"subject_times"`
	Type         struct {
		Key string `json:"key"`
	} `json:"type"`
	LatestRevision int `json:"latest_revision"`
	Revision       int `json:"revision"`
	Created        struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"created"`
	LastModified struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"last_modified"`
}
