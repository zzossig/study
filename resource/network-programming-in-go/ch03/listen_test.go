package ch03

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	t.Logf("bound to %q", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
		}

		go func(conn net.Conn) {
			defer conn.Close()
		}(conn)
	}
}
