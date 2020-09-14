# LRP-tool
Learning resources parser tool.

## How to usage?

### Global options
### How to configure logger
For now have only to options that can be used for the logger configuration:
1. Set logger level thought set the parameter `lrpt --log-level=[panic|fatal|error|warn|info|debug|trace] ...` or using `APP_LOG_FORMATTER_JSON` environment variable.
1. Set logger formatter like JSON thought set the parameter `lrpt --json ...` or using `APP_LOG_FORMATTER_JSON` environment variable.

### Parser
#### How to run parser
To run parser with a known configuration `lrpt parser run`

#### How to run parser and store parsing results to the Dummy database 
To use the Dummy database as a storage engine needed to set up a running flag or environment.
- flag example `lrpt parser run --storage.dummy`
- environment example `APP_STORAGE_DUMMY=true`

#### How to run parser and store parsing results to the PostgreSQL database 
To use the PostgreSQL database as a storage engine, a connection URL setup needed.
This can be in two ways as usual.
The first-way is adding a service running option that contains PostgreSQL connection DSN like the example below `lrpt parser run --storage.postgresql=postgresql://postgres:password@localhost:5432/parser`.
The second-way is to set up the environment variable `APP_STORAGE_POSTGRESQL_DSN` that contains PostgreSQL connection DSN to the application running environment.

### How to inspect a known configuration
To inspect a known configuration run tool with the next command`lrpt parser inspect`. 

### How to import a configuration from xlsx file
To import all configuration files from `./business/categories` to `./config/project` run `lrpt config import` command.

## Used libraries
- github.com/urfave/cli/v2 - Console command framework
- github.com/sirupsen/logrus - Power logger library
- github.com/jszwec/csvutil - CSV marshal/unmarshal
- github.com/gocolly/colly/v2 - Lightning Fast and Elegant Scraping Framework for Gophers
- github.com/stretchr/testify - Framework for testing
- github.com/jackc/pgx/v4 - PostgreSQL Driver and Toolkit
- github.com/360EntSecGroup-Skylar/excelize/v2 - Excel processor
_____________
`Writing with love in Golang!`