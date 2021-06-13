package main

import (
	"bufio"
	"fmt"
	"github.com/go-ping/ping"
	"log"
	"os"
	"time"
)

// readLines reads a whole file into memory and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("addresses.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, line := range lines {
		pinger, err := ping.NewPinger(line)
		if err != nil {
			fmt.Println(err)
		}
		pinger.Count = 2
		pinger.Timeout = 3*time.Second
		pinger.Run()
		if pinger.Statistics().PacketsRecv > 0 {
			fmt.Println(line, "-> got it")
		} else  { fmt.Println(line, "-> no response!") }
	}
}
