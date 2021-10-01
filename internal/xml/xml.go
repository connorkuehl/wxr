package xml

import "encoding/xml"

type Author struct {
	XMLName     xml.Name `xml:"wp author"`
	ID          int      `xml:"wp author_id"`
	Login       string   `xml:"wp author_login"`
	Email       string   `xml:"wp author_email"`
	DisplayName string   `xml:"wp author_display_name"`
	FirstName   string   `xml:"wp author_first_name"`
	LastName    string   `xml:"wp author_last_name"`
}

type Category struct {
	XMLName  xml.Name `xml:"wp category"`
	TermID   int      `xml:"wp term_id"`
	NiceName string   `xml:"wp category_nicename"`
	Parent   string   `xml:"wp category_parent"`
	Name     string   `xml:"wp cat_name"`
}

type Channel struct {
	XMLName     xml.Name   `xml:"channel"`
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	PubDate     string     `xml:"pubDate"`
	Language    string     `xml:"language"`
	WxrVersion  string     `xml:"wp wxr_version"`
	BaseSiteUrl string     `xml:"wp base_site_url"`
	BaseBlogUrl string     `xml:"wp base_blog_url"`
	Authors     []Author   `xml:"wp author"`
	Categories  []Category `xml:"wp category"`
	Terms       []Term     `xml:"wp term"`
	Generator   string     `xml:"generator"`
	Site        Site       `xml:"site"`
	Items       []Item     `xml:"item"`
}

type GUID struct {
	XMLName     xml.Name `xml:"guid"`
	IsPermaLink string   `xml:"isPermaLink,attr"`
}

type ItemCategory struct {
	XMLName  xml.Name `xml:"category"`
	Domain   string   `xml:"domain,attr"`
	NiceName string   `xml:"nicename,attr"`
}

type Item struct {
	XMLName         xml.Name     `xml:"item"`
	Title           string       `xml:"title"`
	Link            string       `xml:"link"`
	PubDate         string       `xml:"pubDate"`
	Creator         string       `xml:"dc creator"`
	GUID            GUID         `xml:"guid"`
	Description     string       `xml:"description"`
	Content         string       `xml:"content encoded"`
	Excerpt         string       `xml:"excerpt encoded"`
	PostID          int          `xml:"wp post_id"`
	PostDate        string       `xml:"wp post_date"`
	PostDateGMT     string       `xml:"wp post_date_gmt"`
	PostModified    string       `xml:"wp post_modified"`
	PostModifiedGMT string       `xml:"wp post_modified_gmt"`
	CommentStatus   string       `xml:"wp comment_status"`
	PingStatus      string       `xml:"wp ping_status"`
	Status          string       `xml:"wp status"`
	PostName        string       `xml:"wp post_name"`
	PostParent      int          `xml:"wp post_parent"`
	MenuOrder       int          `xml:"wp menu_order"`
	PostType        string       `xml:"wp post_type"`
	PostPassword    string       `xml:"wp post_password"`
	IsSticky        int          `xml:"wp is_sticky"`
	Category        ItemCategory `xml:"category"`
	MetaKVs         []PostMeta   `xml:"postmeta"`
}

type PostMeta struct {
	XMLName xml.Name `xml:"wp postmeta"`
	Key     string   `xml:"wp meta_key"`
	Value   string   `xml:"wp meta_value"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Excerpt string   `xml:"xmlns excerpt,attr"`
	Content string   `xml:"xmlns content,attr"`
	Wfw     string   `xml:"xmlns wfw,attr"`
	Dc      string   `xml:"xmlns dc,attr"`
	Wp      string   `xml:"xmlns wp,attr"`
	Channel Channel  `xml:"channel"`
}

type Site struct {
	XMLName xml.Name `xml:"site"`
	Xmlns   string   `xml:"xmlns,attr"`
}

type Term struct {
	XMLName  xml.Name `xml:"wp term"`
	ID       int      `xml:"wp term_id"`
	Taxonomy string   `xml:"wp term_taxonomy"`
	Slug     string   `xml:"wp term_slug"`
	Parent   string   `xml:"wp term_parent"`
	Name     string   `xml:"wp term_name"`
}
