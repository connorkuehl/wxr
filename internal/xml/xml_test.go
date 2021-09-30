package xml

import (
	"encoding/xml"
	"reflect"
	"testing"
)

const authorValidFragment = `
	<wp:author>
		<wp:author_id>1</wp:author_id>
		<wp:author_login>example_author</wp:author_login>
		<wp:author_email>example_author@example.com</wp:author_email>
		<wp:author_display_name>ExamplusAuthorius99</wp:author_display_name>
		<wp:author_first_name>Examplus</wp:author_first_name>
		<wp:author_last_name>Authorious</wp:author_last_name>
	</wp:author>`

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

const categoryValidFragment = `
	<wp:category>
		<wp:term_id>16</wp:term_id>
		<wp:category_nicename>category-nice-name</wp:category_nicename>
		<wp:category_parent></wp:category_parent>
		<wp:cat_name>A Category</wp:cat_name>
	</wp:category>`

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

const itemValidFragment = `
		<item>
			<title>Home</title>
			<link>example.com</link>
			<pubDate>Sun, 29 Nov 2020 16:29:33 +0000</pubDate>
			<dc:creator>CreatorPerson</dc:creator>
			<guid isPermaLink="false">example.com</guid>
			<description>desc</description>
			<content:encoded>empty</content:encoded>
			<excerpt:encoded>empty</excerpt:encoded>
			<wp:post_id>9</wp:post_id>
			<wp:post_date>2021-08-07 07:56:40</wp:post_date>
			<wp:post_date_gmt>2020-11-29 16:29:33</wp:post_date_gmt>
			<wp:post_modified>2021-08-07 07:56:40</wp:post_modified>
			<wp:post_modified_gmt>2021-08-07 12:56:40</wp:post_modified_gmt>
			<wp:comment_status>closed</wp:comment_status>
			<wp:ping_status>closed</wp:ping_status>
			<wp:post_name>home</wp:post_name>
			<wp:status>publish</wp:status>
			<wp:post_parent>0</wp:post_parent>
			<wp:menu_order>1</wp:menu_order>
			<wp:post_type>nav_menu_item</wp:post_type>
			<wp:post_password>passwd</wp:post_password>
			<wp:is_sticky>0</wp:is_sticky>
			<category domain="nav_menu" nicename="main-menu">Main Menu</category>
			<wp:postmeta>
				<wp:meta_key>_menu_item_type</wp:meta_key>
				<wp:meta_value>custom</wp:meta_value>
			</wp:postmeta>
			<wp:postmeta>
				<wp:meta_key>_menu_item_menu_item_parent</wp:meta_key>
				<wp:meta_value>0</wp:meta_value>
			</wp:postmeta>
		</item>
`

func TestItem(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want Item
	}{
		{"valid item fragment", itemValidFragment, Item{
			XMLName:         xml.Name{Local: "item"},
			Title:           "Home",
			Link:            "example.com",
			PubDate:         "Sun, 29 Nov 2020 16:29:33 +0000",
			Creator:         "CreatorPerson",
			GUID:            GUID{XMLName: xml.Name{Local: "guid"}, IsPermaLink: "false"},
			Description:     "desc",
			Content:         "empty",
			Excerpt:         "empty",
			PostID:          9,
			PostDate:        "2021-08-07 07:56:40",
			PostDateGMT:     "2020-11-29 16:29:33",
			PostModified:    "2021-08-07 07:56:40",
			PostModifiedGMT: "2021-08-07 12:56:40",
			CommentStatus:   "closed",
			PingStatus:      "closed",
			PostName:        "home",
			Status:          "publish",
			PostParent:      0,
			MenuOrder:       1,
			PostType:        "nav_menu_item",
			PostPassword:    "passwd",
			IsSticky:        0,
			Category:        ItemCategory{XMLName: xml.Name{Local: "category"}, Domain: "nav_menu", NiceName: "main-menu"},
			MetaKVs: []PostMeta{
				{XMLName: xml.Name{Space: "wp", Local: "postmeta"}, Key: "_menu_item_type", Value: "custom"},
				{XMLName: xml.Name{Space: "wp", Local: "postmeta"}, Key: "_menu_item_menu_item_parent", Value: "0"},
			},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Item

			err := xml.Unmarshal([]byte(tt.in), &got)
			if err != nil {
				t.Errorf("xml.Unmarshal failed for %s: %s", tt.name, err)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("xml.Unmarshal got %+v, want %+v", got, tt.want)
			}
		})
	}
}

const postMetaValidFragment = `
	<wp:postmeta>
		<wp:meta_key>fruit</wp:meta_key>
		<wp:meta_value>apple</wp:meta_value>
	</wp:postmeta>
`

func TestDecodePostMeta(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want PostMeta
	}{
		{"valid postmeta fragment", postMetaValidFragment, PostMeta{
			XMLName: xml.Name{Space: "wp", Local: "postmeta"},
			Key:     "fruit",
			Value:   "apple",
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got PostMeta

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

func TestDecodeSite(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want Site
	}{
		{"valid site with attr", `<site xmlns="com-example:blah:1">1472</site>`, Site{
			// FIXME: not sure why xml.Unmarshal is populating com-example:blah:1 in
			// xml.Name.Space. I don't think it should be. However, it doesn't seem
			// to hurt anything
			XMLName: xml.Name{Space: "com-example:blah:1", Local: "site"},
			Xmlns:   "com-example:blah:1",
		}},
		{"valid site", "<site>1472</site>", Site{
			XMLName: xml.Name{Local: "site"},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Site

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

const termValidFragment = `
	<wp:term>
		<wp:term_id>3</wp:term_id>
		<wp:term_taxonomy>category</wp:term_taxonomy>
		<wp:term_slug>cat</wp:term_slug>
		<wp:term_parent>none</wp:term_parent>
		<wp:term_name>Category</wp:term_name>
	</wp:term>`

func TestDecodeTerm(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want Term
	}{
		{"valid term fragment", termValidFragment, Term{
			XMLName:  xml.Name{Space: "wp", Local: "term"},
			ID:       3,
			Taxonomy: "category",
			Slug:     "cat",
			Parent:   "none",
			Name:     "Category"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Term
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
