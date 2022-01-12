package model

type Poetry struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Dynasty string `json:"dynasty"`
	Content string `json:"content"`
}
