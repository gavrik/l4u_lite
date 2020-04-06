package main

// AdminToken - The main token structure
type AdminToken struct {
	Token            string `json:"Token"`
	TokenDescription string `json:"TokenDescription"`
	ExpireAt         int    `json:"ExpireAt"`
}
