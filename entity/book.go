package entity

type Book struct {
	Publishers    []string `json:"publishers"`
	NumberOfPages int      `json:"number_of_pages"`
	Isbn10        []string `json:"isbn_10"`
	Covers        []int    `json:"covers"`
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
