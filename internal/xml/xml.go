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
