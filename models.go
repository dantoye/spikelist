package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

type Level struct {
	Code     string
	MakerID  string
	Name     string
	Category string
	Votes    map[string]float64
}

func ParseUpload(text string) (*Level, error) {
	pos := strings.LastIndex(text, " ")
	title, code := text[:pos], text[pos:]
	code = codeRegex.FindString(text)
	if code == "" {
		return nil, fmt.Errorf(`upload syntax is !upload The Level Title XXX-XXX-XXX`)
	}
	title = strings.TrimSpace(title)
	category := "-"
	if strings.HasSuffix(title, `(TS)`) {
		category = "TS"
		title = strings.TrimSuffix(title, `(TS)`)
	}
	title = strings.TrimSpace(title)

	if title == "" {
		return nil, fmt.Errorf(`upload syntax is !upload The Level Title XXX-XXX-XXX`)
	}
	return &Level{Code: code, Name: title, Category: category}, nil
}

func ParseQuery(params string) (string, error) {
	parts, err := url.ParseQuery(params)
	if err != nil {
		return "", err
	}
	provider := parts.Get("provider")
	providerID := parts.Get("providerId")
	if provider == "" || providerID == "" {
		return "", fmt.Errorf("invalid request format")
	}
	return provider + ":" + providerID, nil
}

func HandleUpload(text, params string, deps Store) (string, error) {
	nick, err := ParseQuery(params)
	if err != nil {
		return "", err
	}
	level, err := ParseUpload(text)
	if err != nil {
		return "", err
	}
	makerID, err := deps.FindMaker(nick)
	if err != nil {
		return "", err
	}
	level.MakerID = makerID
	if err := deps.UploadLevel(level); err != nil {
		return "", err
	}
	return fmt.Sprintf(`nick "%s" uploaded level "%#v"`, nick, *level), nil
}

func HandleRate(text, params string, deps Store) (string, error) {
	return "", nil
}

func HandleRegister(text, params string, deps Store) (string, error) {
	return "", nil
}

var codeRegex = regexp.MustCompile(`.{3}-.{3}-.{3}`)
