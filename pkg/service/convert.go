package service

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"strings"
	"text/template"
)

const (
	defaultTemplate        = "markdown.tmpl"
	defaultRemediationText = "## Remediation\nThere is no remediation at the moment"
)

var (
	//go:embed template/markdown.tmpl
	content  embed.FS
	snykJson SnykJson
)

type SnykJson struct {
	DisplayOnlySummary bool
	DisplayRemediation bool
	OK                 bool
	Vulnerabilities    []Vulnerability
	DependencyCount    int
}

type Vulnerability struct {
	Id             string
	Title          string
	Credit         []string
	ModuleName     string
	PackageName    string
	Language       string
	PackageManager string
	Description    string
	Summary        string
	Severity       string
	From           []string
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

/*
 * This method trys to find the remediation text in the description and returns it if
 * it is there. If it is return from the remediation title to the end of the description.
 */
func getRemediation(description string) string {
	var index = strings.Index(description, "## Remediation")
	if index > -1 {
		return description[index:]
	}
	// TODO: if no remediation text in the description, try the fixedIn data
	// see - https://github.com/snyk/snyk-to-html/blob/9b8b23702286b1c1e1cdd327659c641b5ed2dbde/src/lib/snyk-to-html.ts#L408
	// if all else fails fall back to the default text
	return defaultRemediationText
}
