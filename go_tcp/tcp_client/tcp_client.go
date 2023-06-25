package main

import (
	"net"
	"os"
    "time"
)

const (
	HOST = "localhost"
	PORT = "22222"
	TYPE = "tcp"
)

func main() {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)

	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

    for true {
        _, err = conn.Write([]byte("This is a message\n"))
        if err != nil {
            println("Write data failed:", err.Error())
            os.Exit(1)
        }

        // buffer to get data
        received := make([]byte, 1024)
        _, err = conn.Read(received)
        if err != nil {
            println("Read data failed:", err.Error())
            os.Exit(1)
        }

        println("Received message:", string(received))
        time.Sleep(1 * time.Second)
    }

	conn.Close()
}
