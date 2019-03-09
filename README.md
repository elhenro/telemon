## telemon

### description

a simple telegram bot that runs on your linux server to provide perfomance stats, general information and notification alerts

### setup

`git clone (this)`

`cd telemon`

get go dependencies 

`go get -d ./...`

place your token in the `token.txt` file for the bot (obtain it from botfather)

`echo "1234567890123:examplePasteYourTokenHere12345" >> token.txt`

run

`go run main.go`

### commands

`hello` returns "hello world"

`cpu` returns current cpu usage, read from `/proc/stat` file

`ip` returns current external ip of your server

more to come..
