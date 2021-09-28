package xml

import "encoding/xml"

type Author struct {
	XMLName     xml.Name `xml:"author"`
	ID          int      `xml:"author_id"`
	Login       string   `xml:"author_login"`
	Email       string   `xml:"author_email"`
	DisplayName string   `xml:"author_display_name"`
	FirstName   string   `xml:"author_first_name"`
	LastName    string   `xml:"author_last_name"`
}

type Category struct {
	XMLName  xml.Name `xml:"category"`
	TermID   int      `xml:"term_id"`
	NiceName string   `xml:"category_nicename"`
	Parent   string   `xml:"category_parent"`
	Name     string   `xml:"cat_name"`
}
