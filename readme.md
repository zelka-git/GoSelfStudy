go commands

go mod init example.com/hello
replace example.com/greetings => ../greetings
go mod init example.com/greetings

go mod tidy

return-greetings-for-multiple-people
go test -v