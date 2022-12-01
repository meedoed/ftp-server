package ftp

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// TODO: добавить чтение файла фиксированным числом или построчно
func (c *Conn) retr(args []string) {
	if len(args) != 1 {
		c.respond(STATUS501)
		return
	}

	path := filepath.Join(c.rootDir, c.workDir, args[0])
	file, err := os.Open(path)
	if err != nil {
		log.Print(err)
		c.respond(STATUS550)
	}
	c.respond(STATUS150)

	dataConn, err := c.dataConnect()
	if err != nil {
		log.Print(err)
		c.respond(STATUS425)
	}
	defer dataConn.Close()

	_, err = io.Copy(dataConn, file)
	if err != nil {
		log.Print(err)
		c.respond(STATUS426)
		return
	}
	io.WriteString(dataConn, c.EOL())
	c.respond(STATUS226)
}
