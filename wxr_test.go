package wxr

import (
	"encoding/xml"
	"reflect"
	"testing"
)

const authorValidFragment = `
	<author>
		<author_id>1</author_id>
		<author_login>example_author</author_login>
		<author_email>example_author@example.com</author_email>
		<author_display_name>ExamplusAuthorius99</author_display_name>
		<author_first_name>Examplus</author_first_name>
		<author_last_name>Authorious</author_last_name>
	</author>`

func TestDecodeAuthor(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want Author
	}{
		{"valid author fragment", authorValidFragment, Author{
			XMLName:     xml.Name{Local: "author"},
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
	<category>
		<term_id>16</term_id>
		<category_nicename>category-nice-name</category_nicename>
		<category_parent></category_parent>
		<cat_name>A Category</cat_name>
	</category>`

func TestDecodeCategory(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want Category
	}{
		{"valid category fragment", categoryValidFragment, Category{
			XMLName:  xml.Name{Local: "category"},
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
			<post_id>9</post_id>
			<post_date>2021-08-07 07:56:40</post_date>
			<post_date_gmt>2020-11-29 16:29:33</post_date_gmt>
			<post_modified>2021-08-07 07:56:40</post_modified>
			<post_modified_gmt>2021-08-07 12:56:40</post_modified_gmt>
			<comment_status>closed</comment_status>
			<ping_status>closed</ping_status>
			<post_name>home</post_name>
			<status>publish</status>
			<post_parent>0</post_parent>
			<menu_order>1</menu_order>
			<post_type>nav_menu_item</post_type>
			<post_password>passwd</post_password>
			<is_sticky>0</is_sticky>
			<category domain="nav_menu" nicename="main-menu">Main Menu</category>
			<postmeta>
				<meta_key>_menu_item_type</meta_key>
				<meta_value>custom</meta_value>
			</postmeta>
			<postmeta>
				<meta_key>_menu_item_menu_item_parent</meta_key>
				<meta_value>0</meta_value>
			</postmeta>
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
				{XMLName: xml.Name{Local: "postmeta"}, Key: "_menu_item_type", Value: "custom"},
				{XMLName: xml.Name{Local: "postmeta"}, Key: "_menu_item_menu_item_parent", Value: "0"},
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
	<postmeta>
		<meta_key>fruit</meta_key>
		<meta_value>apple</meta_value>
	</postmeta>
`

func TestDecodePostMeta(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want PostMeta
	}{
		{"valid postmeta fragment", postMetaValidFragment, PostMeta{
			XMLName: xml.Name{Local: "postmeta"},
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
	<term>
		<term_id>3</term_id>
		<term_taxonomy>category</term_taxonomy>
		<term_slug>cat</term_slug>
		<term_parent>none</term_parent>
		<term_name>Category</term_name>
	</term>`

func TestDecodeTerm(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want Term
	}{
		{"valid term fragment", termValidFragment, Term{
			XMLName:  xml.Name{Local: "term"},
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
