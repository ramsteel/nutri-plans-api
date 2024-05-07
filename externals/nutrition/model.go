package nutrition

type Item struct {
	ID    string `json:"tag_id"`
	Name  string `json:"tag_name"`
	Photo Photo  `json:"photo"`
}

type Photo struct {
	Thumb string `json:"thumb"`
}

type SearchItemResponse struct {
	Common *[]Item `json:"common"`
}
