package model

type Url struct{
	Id string `json:"id"`
	ShortCode string `json:"short-code"`
	OriginalURL string `json:"original-url"`
}