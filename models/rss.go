package models

type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title         string   `xml:"title"`
	Description   string   `xml:"description"`
	Link          string   `xml:"link"`
	Image         Image    `xml:"image"`
	Generator     string   `xml:"generator"`
	LastBuildDate string   `xml:"lastBuildDate"`
	AtomLink      AtomLink `xml:"http://www.w3.org/2005/Atom link"`
	Language      string   `xml:"language"`
	Items         []Item   `xml:"item"`
}

type Image struct {
	URL   string `xml:"url"`
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

type AtomLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type Item struct {
	Title       string    `xml:"title" json:"title"`
	Description string    `xml:"description" json:"description"`
	Link        string    `xml:"link" json:"link"`
	GUID        GUID      `xml:"guid" json:"gu_id"`
	PubDate     string    `xml:"pubDate" json:"pub_date"`
	Enclosure   Enclosure `xml:"enclosure" json:"enclosure"`
	Creator     string    `xml:"http://purl.org/dc/elements/1.1/ creator" json:"creator"`
}

type GUID struct {
	IsPermaLink bool   `xml:"isPermaLink,attr" json:"is_perma_link"`
	Value       string `xml:",chardata" json:"value"`
}

type Enclosure struct {
	URL    string `xml:"url,attr" json:"url"`
	Length string `xml:"length,attr" json:"length"`
	Type   string `xml:"type,attr" json:"type"`
}
