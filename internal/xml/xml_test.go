package xml

import (
	"encoding/xml"
	"testing"
)

const (
	authorValidFragment = `
	<wp:author>
		<wp:author_id>1</wp:author_id>
		<wp:author_login>example_author</wp:author_login>
		<wp:author_email>example_author@example.com</wp:author_email>
		<wp:author_display_name>ExamplusAuthorius99</wp:author_display_name>
		<wp:author_first_name>Examplus</wp:author_first_name>
		<wp:author_last_name>Authorious</wp:author_last_name>
	</wp:author>`
)

func TestDecodeAuthor(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want Author
	}{
		{"valid author fragment", authorValidFragment, Author{
			XMLName:     xml.Name{Space: "wp", Local: "author"},
			ID:          1,
			Login:       "example_author",
			Email:       "example_author@example.com",
			DisplayName: "ExamplusAuthorius99",
			FirstName:   "Examplus",
			LastName:    "Authorious"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Author
			err := xml.Unmarshal([]byte(tt.in), &got)
			if err != nil {
				t.Errorf("xml.Unmarshal failed for %s: %s", tt.name, err)
				return
			}

			if got != tt.want {
				t.Errorf("xml.Unmarshal got %+v, want %+v", got, tt.want)
			}
		})
	}
}

const (
	categoryValidFragment = `
	<wp:category>
		<wp:term_id>16</wp:term_id>
		<wp:category_nicename>category-nice-name</wp:category_nicename>
		<wp:category_parent></wp:category_parent>
		<wp:cat_name>A Category</wp:cat_name>
	</wp:category>`
)

func TestDecodeCategory(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want Category
	}{
		{"valid category fragment", categoryValidFragment, Category{
			XMLName:  xml.Name{Space: "wp", Local: "category"},
			TermID:   16,
			NiceName: "category-nice-name",
			Parent:   "",
			Name:     "A Category"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Category
			err := xml.Unmarshal([]byte(tt.in), &got)
			if err != nil {
				t.Errorf("xml.Unmarshal failed for %s: %s", tt.name, err)
				return
			}

			if got != tt.want {
				t.Errorf("xml.Unmarshal got %+v, want %+v", got, tt.want)
			}
		})
	}
}
