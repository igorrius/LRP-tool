# LRP-tool
Learning resources parser tool.

## How to usage?

### Global options
#### Log level
To set logger level use `--log-level=[panic|fatal|error|warn|info|debug|trace]` parameter or `LOG_LEVEL` environment variable.
#### JSON logger formatter
To convert logger output to the JSON format use `--json` parameter or `LOG_FORMATTER_JSON` environment variable.

### Parser
#### How to run parser
To run parser with a known configuration `lrpt [-log-level=debug] parser run`

### How to inspect a known configuration
To inspect a known configuration run tool with the next command`lrpt [--log-level=debug --json] parser inspect`. 

## Used libraries
- github.com/urfave/cli/v2 - console command framework
- github.com/sirupsen/logrus - power logger library
- github.com/jszwec/csvutil - CSV marshal/unmarshal
- github.com/gocolly/colly/v2 - Lightning Fast and Elegant Scraping Framework for Gophers
- github.com/stretchr/testify - Framework for testing
_____________
`Writing with love in Golang!`