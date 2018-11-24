// A simple tcp server to interact with the k/v store

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	connHost = "localhost"
	connType = "tcp"
)

type Handler struct {
	d *Datastore
}

func (h *Handler) handleSet(msg string) {
	parts := strings.SplitN(msg, " ", 3)
	h.d.Set(parts[1], []byte(parts[2]))
}

func (h *Handler) handleGet(msg string) []byte {
	parts := strings.SplitN(msg, " ", 2)
	return h.d.Get(parts[1])
}

func (h *Handler) Handle(c net.Conn) {
	defer func() {
		c.Write([]byte("END\n"))
		c.Close()
	}()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Printf("received: %s\n", msg)
		switch {
		case strings.ToUpper(msg) == "EXIT":
			return
		case strings.HasPrefix(msg, "SET "):
			h.handleSet(msg)
		case strings.HasPrefix(msg, "GET "):
			c.Write(h.handleGet(msg))
			c.Write([]byte{10})
		default:
			fmt.Println("Unknown message")
		}
	}
}

func NewHandler(d *Datastore) *Handler {
	return &Handler{
		d: d,
	}
}

type Server struct {
	handler *Handler
	port    int
}

func NewServer(port int, handler *Handler) *Server {
	return &Server{
		handler: handler,
		port:    port,
	}
}

func (s *Server) Serve() {
	connStr := fmt.Sprintf("%s:%d", connHost, s.port)
	fmt.Printf("Serving on %s\n", connStr)
	ln, err := net.Listen("tcp", connStr)
	defer ln.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go s.handler.Handle(conn)
	}
}
