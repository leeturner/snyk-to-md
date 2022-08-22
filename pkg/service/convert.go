package service

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"text/template"
)

const (
	defaultTemplate = "markdown.tmpl"
)

var (
	//go:embed template/markdown.tmpl
	content  embed.FS
	snykJson SnykJson
)

type Vulnerability struct {
	Title       string
	ModuleName  string
	PackageName string
	Language    string
	Description string
}

type SnykJson struct {
	DisplayOnlySummary bool
	DisplayRemediation bool
	OK                 bool
	Vulnerabilities    []Vulnerability
	DependencyCount    int
}

func Convert(jsonInput string, displayOnlySummary bool, displayRemediation bool, logger zap.SugaredLogger) (string, error) {
	logger.Debug("Converting json...")

	err := json.Unmarshal([]byte(jsonInput), &snykJson)
	if err != nil || snykJson.Vulnerabilities == nil {
		logger.Error(err)
		return "", errors.New("unable to parse snyk json")
	}
	snykJson.DisplayRemediation = displayRemediation
	snykJson.DisplayOnlySummary = displayOnlySummary

	var markdown bytes.Buffer
	tmpl, err := template.New(defaultTemplate).ParseFS(content, "template/"+defaultTemplate)
	if err != nil {
		logger.Error(err)
		return "", errors.New("unable to parse snyk template file")
	}

	err = tmpl.Execute(&markdown, snykJson)
	if err != nil {
		logger.Error(err)
		return "", errors.New("unable to parse snyk template file")
	}
	return markdown.String(), nil
}
