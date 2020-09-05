package main

import (
	"io"
    "log"
	"net"
    "flag"
    "golang.org/x/net/proxy"
)

var socks5 string

func init() {
    flag.StringVar(&socks5, "socks5", "127.0.0.1:9050", "The Socks5 server")
    flag.Parse()

    if flag.NArg() != 2 {
        log.Fatalln("Usage: CMD [--socks5 SERVER] LocalAddr RemoteAddr")
    }
}

func main() {
    local, remote := flag.Arg(0), flag.Arg(1)
    log.Printf("SockIt: %s -> %s -> %s\n", local, socks5, remote)
	ln, err := net.Listen("tcp", local)
	if err != nil {
        log.Fatalln("Listen")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
            log.Panicln("Accept")
		}

        client := conn.RemoteAddr().String()
        log.Printf("%s: New Client\n", client)
        err = handleConn(conn, remote)
        if err != nil {
            log.Printf("%s: handleConn error: %s\n", client, err)
            conn.Close()
        }
    }
}

func handleConn(conn net.Conn, remote string) error {
    dialer, err := proxy.SOCKS5("tcp", socks5, nil, proxy.Direct)
    if err != nil {
        return err
    }

    proxy, err := dialer.Dial("tcp", remote)
    if err != nil {
        return err
    }

    go copyIO(conn, proxy)
    go copyIO(proxy, conn)
    return nil
}

func copyIO(src, dst net.Conn) {
    defer src.Close()
    defer dst.Close()
    io.Copy(src, dst)
}

