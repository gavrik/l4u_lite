package main

// RESTError - The main Error structure
type RESTError struct {
	ErrorNo  int
	ErrorMsg string
}

// AdminToken - The main token structure
type AdminToken struct {
	Token            string `json:"Token"`
	TokenDescription string `json:"TokenDescription"`
	Domain           string `json:"Domain"`
	ExpireAt         int    `json:"ExpireAt"`
	IsRoot           bool   `json:"IsRoot"`
}

// Link - The Link structure
type Link struct {
	LongLink   string `json:"longLink"`
	ShortLink  string `json:"shortLink"`
	Domain     string `json:"domain"`
	IsEnabled  bool   `json:"isEnabled"`
	CreationOn int    `json:"creationOn"`
}
