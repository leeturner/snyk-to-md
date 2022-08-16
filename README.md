# Snyk JSON to Markdown Mapper

The Snyk JSON to Markdown Mapper takes the json outputted from `snyk test --json` and creates a local HTML file 
displaying the vulnerabilities discovered.

## Todo:

* Implement all flag:
    * -i 	--input 	Input path from where to read the json. Defaults to stdin
    * -o 	--output 	Output of the resulting HTML. Example: -o snyk.html. Defaults to stdout
    * -s 	--summary 	Generates an HTML with only the summary, instead of the details report. Defaults to details vulnerability report
    * -d 	--debug 	Runs the CLI in debug mode
    * -a 	--actionable-remediation 	Display actionable remediation info if available 
    * -t 	--template 	Template location for generating the html. Defaults to template/test-report.hbs
* Templates - https://github.com/aymerick/raymond
