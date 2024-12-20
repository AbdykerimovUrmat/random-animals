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
	fmt.Println(t)
	i := (t) % len(entries)
	fmt.Println(i)
	randomEntry := entries[i]
	fmt.Println(randomEntry.Name())
	resFile := filepath.Join("images", randomEntry.Name())

	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")
	c.File(resFile)
}
func main() {
	router := gin.Default()
	router.GET("/random-animal", getAlbums)

	router.Run(":8080")
}
