package entity

type InfoBook struct {
	Description   string   `json:"description"`
	Title         string   `json:"title"`
	Error         string   `json:"error"`
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

type EditionBook struct {
	Publishers    []string `json:"publishers"`
	NumberOfPages int      `json:"number_of_pages"`
	Isbn10        []string `json:"isbn_10"`
	Covers        []int    `json:"covers"`
	Error         string   `json:"error"`
	Key           string   `json:"key"`
	Authors       []struct {
		Key string `json:"key"`
	} `json:"authors"`
	Ocaid         string   `json:"ocaid"`
	Contributions []string `json:"contributions"`
	Languages     []struct {
		Key string `json:"key"`
	} `json:"languages"`
	Classifications struct {
	} `json:"classifications"`
	SourceRecords []string `json:"source_records"`
	Title         string   `json:"title"`
	Identifiers   struct {
		Goodreads    []string `json:"goodreads"`
		Librarything []string `json:"librarything"`
	} `json:"identifiers"`
	Isbn13      []string `json:"isbn_13"`
	LocalID     []string `json:"local_id"`
	PublishDate string   `json:"publish_date"`
	Works       []struct {
		Key string `json:"key"`
	} `json:"works"`
	Type struct {
		Key string `json:"key"`
	} `json:"type"`
	FirstSentence struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"first_sentence"`
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
