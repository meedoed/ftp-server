package ftp

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func (c *Conn) list(args []string) {
	var target string
	if len(args) > 0 {
		target = filepath.Join(c.rootDir, c.workDir, args[0])
	} else {
		target = filepath.Join(c.rootDir, c.workDir)
	}

	files, err := os.ReadDir(target)
	if err != nil {
		log.Print(err)
		c.respond(STATUS550)
		return
	}
	c.respond(STATUS150)

	dataConn, err := c.dataConnect()
	if err != nil {
		log.Print(err)
		c.respond(STATUS425)
		return
	}
	defer dataConn.Close()

	for _, file := range files {
		_, err := fmt.Fprint(dataConn, file.Name(), c.EOL())
		if err != nil {
			log.Print(err)
			c.respond(STATUS426)
		}
	}
	_, err = fmt.Fprintf(dataConn, c.EOL())
	if err != nil {
		log.Print(err)
		c.respond(STATUS426)
	}
	c.respond(STATUS226)
}
