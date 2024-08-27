package service

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type SpellCheckService struct {
	apiURL string
}

func NewSpellCheckService(apiURL string) *SpellCheckService {
	return &SpellCheckService{apiURL: apiURL}
}

type SpellCheckResult struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

func (s *SpellCheckService) GettingErrors(text string) ([][]SpellCheckResult, error) {
	resp, err := http.PostForm(s.apiURL, url.Values{
		"text": {text},
	})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var results [][]SpellCheckResult
	if err = json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *SpellCheckService) SpellChecking(spellCheck [][]SpellCheckResult, text string) (string, error) {
	if len(spellCheck) == 0 || len(spellCheck[0]) == 0 {
		return text, nil
	}

	for _, errRes := range spellCheck[0] {
		if len(errRes.S) > 0 {
			text = strings.Replace(text, errRes.Word, errRes.S[0], -1)
		}
	}

	return text, nil
}
