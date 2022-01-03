package github

import (
	"html/template"
	"time"
)

const IssueURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   template.HTML `json:"html_url"`
	Title     template.HTML
	State     template.HTML
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      template.HTML
}

type User struct {
	Login   template.HTML
	HTMLURL template.HTML `json:"html_url"`
}
