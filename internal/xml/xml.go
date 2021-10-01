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

type Channel struct {
	XMLName     xml.Name   `xml:"channel"`
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	PubDate     string     `xml:"pubDate"`
	Language    string     `xml:"language"`
	WxrVersion  string     `xml:"wxr_version"`
	BaseSiteUrl string     `xml:"base_site_url"`
	BaseBlogUrl string     `xml:"base_blog_url"`
	Authors     []Author   `xml:"author"`
	Categories  []Category `xml:"category"`
	Terms       []Term     `xml:"term"`
	Generator   string     `xml:"generator"`
	Site        Site       `xml:"site"`
	Items       []Item     `xml:"item"`
}

type Content struct {
	XMLName xml.Name `xml:"encoded"`
	Data    string   `xml:",cdata"`
}

type Excerpt struct {
	XMLName xml.Name `xml:"encoded"`
	Data    string   `xml:",cdata"`
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
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Creator     string   `xml:"dc creator"`
	GUID        GUID     `xml:"guid"`
	Description string   `xml:"description"`

	// FIXME: Using these URLs here is brittle. If WordPress ever changes
	// the URLs associated with content or excerpt this will break. However,
	// I can't get the namespace shortnames to work here:
	//
	// `xml:"content encoded"`
	// `xml:"excerpt encoded"`
	//
	// Despite these being set in the top-level <rss> tag.
	Content Content `xml:"http://purl.org/rss/1.0/modules/content/ encoded"`
	Excerpt Excerpt `xml:"http://wordpress.org/export/1.2/excerpt/ encoded"`

	PostID          int          `xml:"post_id"`
	PostDate        string       `xml:"post_date"`
	PostDateGMT     string       `xml:"post_date_gmt"`
	PostModified    string       `xml:"post_modified"`
	PostModifiedGMT string       `xml:"post_modified_gmt"`
	CommentStatus   string       `xml:"comment_status"`
	PingStatus      string       `xml:"ping_status"`
	Status          string       `xml:"status"`
	PostName        string       `xml:"post_name"`
	PostParent      int          `xml:"post_parent"`
	MenuOrder       int          `xml:"menu_order"`
	PostType        string       `xml:"post_type"`
	PostPassword    string       `xml:"post_password"`
	IsSticky        int          `xml:"is_sticky"`
	Category        ItemCategory `xml:"category"`
	MetaKVs         []PostMeta   `xml:"postmeta"`
}

type PostMeta struct {
	XMLName xml.Name `xml:"postmeta"`
	Key     string   `xml:"meta_key"`
	Value   string   `xml:"meta_value"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Excerpt string   `xml:"excerpt,attr"`
	Content string   `xml:"content,attr"`
	Wfw     string   `xml:"wfw,attr"`
	Dc      string   `xml:"dc,attr"`
	Wp      string   `xml:"wp,attr"`
	Channel Channel  `xml:"channel"`
}

type Site struct {
	XMLName xml.Name `xml:"site"`
	Xmlns   string   `xml:"xmlns,attr"`
}

type Term struct {
	XMLName  xml.Name `xml:"term"`
	ID       int      `xml:"term_id"`
	Taxonomy string   `xml:"term_taxonomy"`
	Slug     string   `xml:"term_slug"`
	Parent   string   `xml:"term_parent"`
	Name     string   `xml:"term_name"`
}
