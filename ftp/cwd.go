package ftp

import (
	"log"
	"os"
	"path/filepath"
)

// TODO: улучшить защиту от доступа к файлам вне public
func (c *Conn) cwd(args []string) {
	if len(args) != 1 {
		c.respond(STATUS501)
		return
	}

	workDir := filepath.Join(c.workDir, args[0])
	absPath := filepath.Join(c.rootDir, workDir)

	_, err := os.Stat(absPath)
	if err != nil {
		log.Print(err)
		c.respond(STATUS550)
		return
	}
	c.workDir = workDir
	c.respond(STATUS200)
}
