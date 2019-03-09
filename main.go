package main

import (
	"fmt"
	"io/ioutil"
	"log"
	//"os/exec"
	tb "gopkg.in/tucnak/telebot.v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {

	// read token from file
	tok, err := ioutil.ReadFile("token.txt")
	if err != nil {
		fmt.Print(err)
	}
	token := strings.TrimSpace(string(tok))

	// connect with telegram bot api
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world!")
	})
	b.Handle("ip", func(m *tb.Message) {
		myExternalIp := getExternalIp()
		b.Send(m.Sender, myExternalIp)
	})
	b.Handle("cpu", func(m *tb.Message) {
		cpuUsage := getCPUUsage()
		b.Send(m.Sender, cpuUsage)
	})
	b.Handle("memory", func(m *tb.Message) {
		memoryUsage := getMemoryUsage()
		b.Send(m.Sender, memoryUsage)
	})

	// catchall for unknow commands
	b.Handle(tb.OnText, func(m *tb.Message) {
		b.Send(m.Sender, "??")
	})

	b.Start()
}

func getExternalIp() string {
	bodyString := "(empty)"
	resp, err := http.Get("http://ipinfo.io/ip")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			log.Fatal(err2)
		}
		bodyString = string(bodyBytes)
		return bodyString
	}
	return bodyString
}

func getCPUSample() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}
	return
}

func getCPUUsage() string {
	idle0, total0 := getCPUSample()
	time.Sleep(3 * time.Second)
	idle1, total1 := getCPUSample()

	idleTicks := float64(idle1 - idle0)
	totalTicks := float64(total1 - total0)
	cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

	return fmt.Sprintf("CPU usage is %f%% [busy: %f, total: %f]\n", cpuUsage, totalTicks-idleTicks, totalTicks)
}

func getMemoryUsage() string {
	meminfoContents, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		errMsg := "error reading '/proc/meminfo'"
		return errMsg
	}
	lines := strings.Split(string(meminfoContents), "\n")

	MemTotal := 0
	MemFree := 0

	for _, line := range lines {
		if strings.Split(string(line), ":")[0] == "MemTotal" {
			memtotal, err := strconv.Atoi(strings.TrimSpace(strings.Replace((strings.Split(string(line), ":")[1]), "kB", "", -1)))
			if err != nil {
				fmt.Println(err)
			}
			MemTotal = memtotal
		}
		if strings.Split(string(line), ":")[0] == "MemFree" {
			memfree, err := strconv.Atoi(strings.TrimSpace(strings.Replace((strings.Split(string(line), ":")[1]), "kB", "", -1)))
			if err != nil {
				fmt.Println(err)
			}
			MemFree = memfree
		}
	}
	MemUsage := strconv.FormatFloat((float64(MemTotal)-float64(MemFree))/float64(MemTotal), 'f', 6, 64)

	return fmt.Sprintf("%s%% of memory used", MemUsage)
}
