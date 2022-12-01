package ftp

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type Conn struct {
	conn     net.Conn
	dataType dataType
	dataPort *dataPort
	rootDir  string
	workDir  string
}

// NewConn создание нового подключения
func NewConn(conn net.Conn, rootDir string) *Conn {
	return &Conn{
		conn:    conn,
		rootDir: rootDir,
		workDir: "/",
	}
}

// Serve Обработка запросов
func Serve(c *Conn) {
	c.respond(STATUS220)

	s := bufio.NewScanner(c.conn)
	for s.Scan() {
		input := strings.Fields(s.Text())
		if len(input) == 0 {
			continue
		}
		command, args := input[0], input[1:]
		log.Printf("<< %s %v", command, args)

		switch command {
		case "CWD":
			c.cwd(args)
		case "LIST":
			c.list(args)
		case "PORT":
			c.port(args)
		case "USER":
			c.user(args)
		case "QUIT":
			c.respond(STATUS221)
		case "RETR":
			c.retr(args)
		case "TYPE":
			c.setDataType(args)
		default:
			c.respond(STATUS502)
		}
	}
	if s.Err() != nil {
		log.Println(s.Err())
	}
}
