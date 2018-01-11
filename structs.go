package main

// Page struct
type Page struct {
	AccessToken   string
	Title         string
	AuthorName    string
	AuthorURL     string
	Content       string
	ReturnContent bool
}

// ContentBlock content struct
type ContentBlock struct {
	Tag      string   `json:"tag"`
	Children []string `json:"children"`
}
