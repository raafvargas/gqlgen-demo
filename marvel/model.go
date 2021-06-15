package marvel

import "encoding/json"

type MarvelResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   *MarvelData `json:"data"`
}

type MarvelData struct {
	Results json.RawMessage `json:"results"`
}

type Character struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Comic struct {
	ID                 int    `json:"id"`
	Title              string `json:"title"`
	VariantDescription string `json:"variantDescription"`
}

type Story struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
