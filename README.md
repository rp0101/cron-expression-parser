# cron-expression-parser

cron-expression-parser is a simple cron parser that tells when a cron will be run with a given schedule.

## Local Setup
* Clone the repository and unzip in your local machine
* Install golang: https://golang.org/doc/install
* Download and Install all the dependent packages
```bash
 go mod download
```
* Build the project
```bash
 go build
```
* Run the project, pass the cron string as an argument after the run command
```bash
 go run main.go "cron_string"
 eg: go run main.go "*/15 0 */2 * 1-7 /usr/bin/find"
```
* Run the tests
```bash
 go test cron-expression-parser/tests -run nameOfTextCase
 eg: go test cron-expression-parser/tests -run TestParseCronService_Fail
```
