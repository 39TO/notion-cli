package types

// Blocks
type Block struct {
	Object    string `json:"object"`
	ID        string `json:"id"`
	Type      string `json:"type"`
	ChildPage struct {
		Title string `json:"title"`
	} `json:"child_page"`
	Paragraph struct {
		RichText []struct {
			PlainText string `json:"plain_text"`
		} `json:"rich_text"`
	} `json:"paragraph"`
}

type BlocksResponse struct {
	Results []Block `json:"results"`
}

// Pages
type Page struct {
	ID         string `json:"id"`
	Properties struct {
		Title struct {
			Title []struct {
				PlainText string `json:"plain_text"`
			} `json:"title"`
		} `json:"title"`
	} `json:"properties"`
}
