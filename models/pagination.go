package models

type PaginationOpts struct {
	Page  int
	Skip  int
	Limit int
	Total int
}

type Links struct {
	Self string `json:"self"`
	Prev string `json:"prev,omitempty"`
	Next string `json:"next,omitempty"`
}

type Page struct {
	Links   Links       `json:"_links"`
	Limit   int         `json:"limit"`
	Start   int         `json:"start"`
	Size    int         `json:"perPage"`
	Total   int         `json:"total"`
	Results interface{} `json:"results"`
}
