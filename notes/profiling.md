To check coverage rate:

`$ go test -cover ./...`

More detailed breakdown of test coverage by method and function:

`$ go test -coverprofile=/tmp/profile.out ./...`

To view coverage profile:

`$ go tool cover -func=/tmp/profile.out`

Visual way to view coverage profile:

`$ go tool cover -html=/tmp/profile.out`

Instead of just highlighting the statements in green and red, using -covermode=count makes
the coverage profile record the exact number of times that each statement is executed during
the tests:

`$ go test -covermode=count -coverprofile=/tmp/profile.out ./...`
`go tool cover -html=/tmp/profile.out`
