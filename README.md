## telemon

Chat with your server

![Example](https://raw.githubusercontent.com/elhenro/telemon/master/img.png)



### description

a simple telegram bot that runs on your linux server to provide perfomance stats, general information and notification alerts

### setup

`git clone https://github.com/elhenro/telemon`

`cd telemon`

get go dependencies 

`go get -d ./...`

place your token in the `token.txt` file for the bot (obtain it from [botfather](https://telegram.me/BotFather))

`echo "1234567890123:examplePasteYourTokenHere12345" >> token.txt`

run

`go run main.go`

### commands

introduce yourself with `hello`. after that the bot will send you alerts if something is down. 

`cpu` returns current cpu usage

`ip` returns current external ip of your server

`memory` returns current memory usage in percent

more to come..
