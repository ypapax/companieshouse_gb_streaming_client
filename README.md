Using UK companieshouse streaming API as described here https://developer-specs.companieshouse.gov.uk/streaming-api/guides/overview
```
$ go run main.go
```
```
2020/04/26 13:58:11 main.go:24: listening to https://stream.companieshouse.gov.uk/companies
2020/04/26 13:58:34 main.go:38: reading the body
2020/04/26 13:58:34 main.go:43: error: EOF
2020/04/26 13:58:34 main.go:19: exiting with error: EOF
exit status 1
```