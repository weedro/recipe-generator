package domain

type Ingredient struct {
	Name     string `json:"name"`
	Quantity uint8  `json:"quantity"`
}

type Recipe struct {
	Hash        string       `json:"hash"`
	Prefix      string       `json:"prefix"`
	Adjective   string       `json:"adjective"`
	Icon        int          `json:"icon"`
	Ingredients []Ingredient `json:"ingredients"`
}
