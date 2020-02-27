package main

import (
	"bytes"
	"io/ioutil"
	"net"
	"strings"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	go func() {
		main()
	}()

	time.Sleep(1_000_000_000)

	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		t.Fatalf("can't connect to server")
	}
	defer conn.Close()

	var buffer bytes.Buffer

	buffer.Write([]byte("GET / HTTP/1.1\r\nHost: localhost\r\n\r\n"))
	buffer.WriteTo(conn)

	allBytes, err := ioutil.ReadAll(conn)
	if err != nil {
		t.Fatalf("can't read response from server: %v", err)
	}
	response := string(allBytes)
	if !strings.Contains(response, "HTTP/1.1 200 OK") {
		t.Fatalf("incorrect responce: %s", response)
	}



}