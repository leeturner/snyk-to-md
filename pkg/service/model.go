package service

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
