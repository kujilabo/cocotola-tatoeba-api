package entity

import "time"

type TatoebaSentenceFindParameter struct {
	PageNo   int    `json:"pageNo" binding:"required,gte=1"`
	PageSize int    `json:"pageSize" binding:"required,gte=1"`
	Keyword  string `json:"keyword"`
	Random   bool   `json:"random"`
}

type TatoebaSentenceResponse struct {
	SentenceNumber int       `json:"sentenceNumber"`
	Lang           string    `json:"lang"`
	Text           string    `json:"text"`
	Author         string    `json:"author"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type TatoebaSentencePair struct {
	Src TatoebaSentenceResponse `json:"src"`
	Dst TatoebaSentenceResponse `json:"dst"`
}

type TatoebaSentencePairFindResponse struct {
	TotalCount int                   `json:"totalCount"`
	Results    []TatoebaSentencePair `json:"results"`
}
