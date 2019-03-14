package view

import (
	"html/template"
	"io"

	"golang-demos/crawler-multi-thread/frontend/model"
)

// SearchResultView struct
type SearchResultView struct {
	template *template.Template
}

// CreateSearchResultView create a search result view
func CreateSearchResultView(
	filename string) SearchResultView {
	return SearchResultView{
		template: template.Must(
			template.ParseFiles(filename)),
	}
}

// Render render the template
func (s SearchResultView) Render(
	w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}
