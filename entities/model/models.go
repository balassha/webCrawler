package model

type URL struct {
	URL string `json:"url"`
}

type Links struct {
	ExternalLinks []string
	InternalLinks []string
}
