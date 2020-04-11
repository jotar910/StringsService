package api

// Parser struct is the dto model for a parse string request.
type Parser struct {
	Text  string   `json:"text"`
	Flags []string `json:"flags"`
}

// ParserRes struct is the dto model for the result of a parsed string
type ParserRes struct {
	Result string `json:"result"`
}
