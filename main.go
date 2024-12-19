package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	_, file, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(file) + "/images"
	entries, e := os.ReadDir(basepath)

	if e != nil {
		fmt.Println(e)
		c.String(400, "", "ОШИБОЧКА ВЫШЛА")
	}

	t := time.Now().YearDay()
	randomEntry := entries[t%len(entries)]

	resFile := filepath.Join("images", randomEntry.Name())

	c.File(resFile)
}
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
