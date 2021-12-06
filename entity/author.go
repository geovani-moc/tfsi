package entity

type AuthorByName struct {
	NumFound      int  `json:"numFound"`
	Start         int  `json:"start"`
	NumFoundExact bool `json:"numFoundExact"`
	Docs          []struct {
		Key            string   `json:"key"`
		Type           string   `json:"type"`
		Name           string   `json:"name"`
		AlternateNames []string `json:"alternate_names,omitempty"`
		BirthDate      string   `json:"birth_date,omitempty"`
		TopWork        string   `json:"top_work"`
		WorkCount      int      `json:"work_count"`
		TopSubjects    []string `json:"top_subjects,omitempty"`
		Version        int64    `json:"_version_"`
	} `json:"docs"`
}

type AuthorByCode struct {
	Name          string   `json:"name"`
	Error         string   `json:"error"`
	Title         string   `json:"title"`
	SourceRecords []string `json:"source_records"`
	RemoteIds     struct {
		Isni     string `json:"isni"`
		Wikidata string `json:"wikidata"`
		Viaf     string `json:"viaf"`
	} `json:"remote_ids"`
	Links []struct {
		URL   string `json:"url"`
		Title string `json:"title"`
		Type  struct {
			Key string `json:"key"`
		} `json:"type"`
	} `json:"links"`
	PersonalName   string   `json:"personal_name"`
	Bio            string   `json:"bio"`
	EntityType     string   `json:"entity_type"`
	Wikipedia      string   `json:"wikipedia"`
	AlternateNames []string `json:"alternate_names"`
	BirthDate      string   `json:"birth_date"`
	Key            string   `json:"key"`
	Photos         []int    `json:"photos"`
	Type           struct {
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

type WorksByAuthor struct {
	Links struct {
		Self   string `json:"self"`
		Author string `json:"author"`
		Next   string `json:"next"`
	} `json:"links"`
	Size    int `json:"size"`
	Entries []struct {
		Title   string `json:"title"`
		Authors []struct {
			Author struct {
				Key string `json:"key"`
			} `json:"author"`
			Type struct {
				Key string `json:"key"`
			} `json:"type"`
		} `json:"authors"`
		Key  string `json:"key"`
		Type struct {
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
		Covers      []int       `json:"covers,omitempty"`
		Description interface{} `json:"description,omitempty"`
		Subjects    []string    `json:"subjects,omitempty"`
		Links       []struct {
			URL   string `json:"url"`
			Title string `json:"title"`
			Type  struct {
				Key string `json:"key"`
			} `json:"type"`
		} `json:"links,omitempty"`
		SubjectPlaces []string `json:"subject_places,omitempty"`
		SubjectPeople []string `json:"subject_people,omitempty"`
		Excerpts      []struct {
			Excerpt string `json:"excerpt"`
			Comment string `json:"comment"`
			Author  struct {
				Key string `json:"key"`
			} `json:"author"`
		} `json:"excerpts,omitempty"`
		SubjectTimes []string `json:"subject_times,omitempty"`
	} `json:"entries"`
}
