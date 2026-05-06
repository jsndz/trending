package techcrunch

type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title       string   `xml:"title"`
	PublishedAt string   `xml:"pubDate"`
	Link        string   `xml:"link"`
	Author      string   `xml:"creator"`
	Categories  []string `xml:"category"`
	Description string   `xml:"description"`
}
