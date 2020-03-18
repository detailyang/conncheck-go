// +build !windows

package conncheck

import (
	"io"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCheckClosedConnnection(t *testing.T) {
	// Listening on any available port
	ln, err := net.Listen("tcp4", "127.0.0.1:0")
	require.Nil(t, err)

	go func() {
		conn, _ := ln.Accept()
		time.Sleep(100 * time.Millisecond)
		conn.Close()
	}()

	conn, err := net.Dial("tcp4", ln.Addr().String())
	require.Nil(t, err)
	time.Sleep(300 * time.Millisecond)
	err = Check(conn)
	require.Equal(t, io.EOF, err)
}

func TestCheckLivedConnection(t *testing.T) {
	// Listening on any available port
	ln, err := net.Listen("tcp4", "127.0.0.1:0")
	require.Nil(t, err)

	w := "abcdefhijklmnopqrsdsds"

	go func() {
		conn, _ := ln.Accept()
		time.Sleep(300 * time.Millisecond)
		conn.Write([]byte(w))
	}()

	conn, err := net.Dial("tcp4", ln.Addr().String())
	require.Nil(t, err)
	err = Check(conn)
	require.Equal(t, nil, err)
	var z [1024]byte
	n, err := conn.Read(z[:])
	require.Equal(t, w, string(z[:n]))
}
