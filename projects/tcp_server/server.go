package main

import (
	"fmt"
	"net"
	"bufio"
)

func handleConn(conn *net.TCPConn){
	fmt.Println("connected")
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	defer conn.Close()
	for {
		blah,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ERROR: ",err)
			return
		}
		switch blah {
			case "close\n":
				fmt.Println("closing")
				writer.WriteString("closing\n")
				writer.Flush()
				return
			default:
				fmt.Print("Receiving: ",blah)
				writer.WriteString(string(blah))
				writer.Flush()
		}
	}
}

func main(){
	addr := net.TCPAddr{
		IP: net.IPv4(127,0,0,1),
		Port: 9201,
	}
	ln,err := net.ListenTCP("tcp", &addr)
	if err != nil {
		fmt.Println(err)
	}
	defer ln.Close()
	for {
		fmt.Println("waiting")
		conn,_ := ln.AcceptTCP()
		fmt.Println("accepted")
		go handleConn(conn)
	}
}
