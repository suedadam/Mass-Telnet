package main

import (
	"github.com/ziutek/telnet"
	"log"
	"os"
	"time"
	"bufio"
	"strings"
)

const timeout = 10 * time.Second

func checkErr(err error) {
	if err != nil {
		return
	}
}

func expect(t *telnet.Conn, d ...string) {
	checkErr(t.SetReadDeadline(time.Now().Add(timeout)))
	checkErr(t.SkipUntil(d...))
}

func sendln(t *telnet.Conn, s string) {
	checkErr(t.SetWriteDeadline(time.Now().Add(timeout)))
	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = '\n'
	_, err := t.Write(buf)
	checkErr(err)
}

func main() {
	if len(os.Args) != 5 {
		log.Printf("Usage: %s <file name>", os.Args[0])
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")
		host, user, passwd := s[0], s[1], s[2]
		t, err := telnet.Dial("tcp", host)
		checkErr(err)
		t.SetUnixWriteMode(true)

		var data []byte
			expect(t, "login: ")
			sendln(t, user)
			expect(t, "ssword: ")
			sendln(t, passwd)
			expect(t, "$")
			sendln(t, "ls -l")
			data, err = t.ReadBytes('$')
		checkErr(err)
		os.Stdout.Write(data)
		os.Stdout.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
	    log.Fatal(err)
	}
}