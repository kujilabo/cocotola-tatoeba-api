//go:generate mockery --output mock --name TatoebaSentence
//go:generate mockery --output mock --name TatoebaSentencePair
//go:generate mockery --output mock --name TatoebaSentenceAddParameter
//go:generate mockery --output mock --name TatoebaSentenceSearchCondition
//go:generate mockery --output mock --name TatoebaSentencePairSearchResult
//go:generate mockery --output mock --name TatoebaSentenceRepository
package service

import (
	"context"
	"errors"
	"time"

	"github.com/kujilabo/cocotola-tatoeba-api/src/app/domain"
	libD "github.com/kujilabo/cocotola-tatoeba-api/src/lib/domain"
)

var ErrTatoebaSentenceAlreadyExists = errors.New("tatoebaSentence already exists")
var ErrTatoebaSentenceNotFound = errors.New("tatoebaSentence not found")

type TatoebaSentence interface {
	GetSentenceNumber() int
	GetLang3() domain.Lang3
	GetText() string
	GetAuthor() string
	GetUpdatedAt() time.Time
}

type tatoebaSentence struct {
	SentenceNumber int
	Lang3          domain.Lang3
	Text           string
	Author         string
	UpdatedAt      time.Time
}

func NewTatoebaSentence(sentenceNumber int, lang3 domain.Lang3, text, author string, updatedAt time.Time) (TatoebaSentence, error) {
	m := &tatoebaSentence{
		SentenceNumber: sentenceNumber,
		Lang3:          lang3,
		Text:           text,
		Author:         author,
		UpdatedAt:      updatedAt,
	}

	return m, libD.Validator.Struct(m)
}

func (m *tatoebaSentence) GetSentenceNumber() int {
	return m.SentenceNumber
}

func (m *tatoebaSentence) GetLang3() domain.Lang3 {
	return m.Lang3
}

func (m *tatoebaSentence) GetText() string {
	return m.Text
}

func (m *tatoebaSentence) GetAuthor() string {
	return m.Author
}

func (m *tatoebaSentence) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

type TatoebaSentencePair interface {
	GetSrc() TatoebaSentence
	GetDst() TatoebaSentence
}

type tatoebaSentencePair struct {
	Src TatoebaSentence
	Dst TatoebaSentence
}

func NewTatoebaSentencePair(src, dst TatoebaSentence) (TatoebaSentencePair, error) {
	m := &tatoebaSentencePair{
		Src: src,
		Dst: dst,
	}

	return m, libD.Validator.Struct(m)
}

func (m *tatoebaSentencePair) GetSrc() TatoebaSentence {
	return m.Src
}

func (m *tatoebaSentencePair) GetDst() TatoebaSentence {
	return m.Dst
}

type TatoebaSentenceAddParameter interface {
	GetSentenceNumber() int
	GetLang3() domain.Lang3
	GetText() string
	GetAuthor() string
	GetUpdatedAt() time.Time
}

type tatoebaSentenceAddParameter struct {
	SentenceNumber int `validate:"required"`
	Lang3          domain.Lang3
	Text           string `validate:"required"`
	Author         string `validate:"required"`
	UpdatedAt      time.Time
}

func NewTatoebaSentenceAddParameter(sentenceNumber int, lang3 domain.Lang3, text, author string, updatedAt time.Time) (TatoebaSentenceAddParameter, error) {
	m := &tatoebaSentenceAddParameter{
		SentenceNumber: sentenceNumber,
		Lang3:          lang3,
		Text:           text,
		Author:         author,
		UpdatedAt:      updatedAt,
	}

	return m, libD.Validator.Struct(m)
}

func (p *tatoebaSentenceAddParameter) GetSentenceNumber() int {
	return p.SentenceNumber
}

func (p *tatoebaSentenceAddParameter) GetLang3() domain.Lang3 {
	return p.Lang3
}

func (p *tatoebaSentenceAddParameter) GetText() string {
	return p.Text
}

func (p *tatoebaSentenceAddParameter) GetAuthor() string {
	return p.Author
}

func (p *tatoebaSentenceAddParameter) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

type TatoebaSentenceSearchCondition interface {
	GetPageNo() int
	GetPageSize() int
	GetKeyword() string
	IsRandom() bool
}

type tatoebaSentenceSearchCondition struct {
	PageNo   int `validate:"required,gte=1"`
	PageSize int `validate:"required,gte=1,lte=100"`
	Keyword  string
	Random   bool
}

func NewTatoebaSentenceSearchCondition(pageNo, pageSize int, keyword string, random bool) (TatoebaSentenceSearchCondition, error) {
	m := &tatoebaSentenceSearchCondition{
		PageNo:   pageNo,
		PageSize: pageSize,
		Keyword:  keyword,
		Random:   random,
	}

	return m, libD.Validator.Struct(m)
}

func (c *tatoebaSentenceSearchCondition) GetPageNo() int {
	return c.PageNo
}

func (c *tatoebaSentenceSearchCondition) GetPageSize() int {
	return c.PageSize
}

func (c *tatoebaSentenceSearchCondition) GetKeyword() string {
	return c.Keyword
}

func (c *tatoebaSentenceSearchCondition) IsRandom() bool {
	return c.Random
}

type TatoebaSentencePairSearchResult interface {
	GetTotalCount() int
	GetResults() []TatoebaSentencePair
}

type tatoebaSentencePairSearchResult struct {
	TotalCount int
	Results    []TatoebaSentencePair
}

func NewTatoebaSentencePairSearchResult(totalCount int, results []TatoebaSentencePair) TatoebaSentencePairSearchResult {
	return &tatoebaSentencePairSearchResult{
		TotalCount: totalCount,
		Results:    results,
	}
}

func (r *tatoebaSentencePairSearchResult) GetTotalCount() int {
	return r.TotalCount
}

func (r *tatoebaSentencePairSearchResult) GetResults() []TatoebaSentencePair {
	return r.Results
}

type TatoebaSentenceRepository interface {
	FindTatoebaSentencePairs(ctx context.Context, param TatoebaSentenceSearchCondition) (TatoebaSentencePairSearchResult, error)

	FindTatoebaSentenceBySentenceNumber(ctx context.Context, sentenceNumber int) (TatoebaSentence, error)

	Add(ctx context.Context, param TatoebaSentenceAddParameter) error

	ContainsSentenceBySentenceNumber(ctx context.Context, sentenceNumber int) (bool, error)
}
