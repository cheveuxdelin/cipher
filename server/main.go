package main

import (
	"net/http"

	"github.com/cheveuxdelin/cipher/atbash"
	"github.com/cheveuxdelin/cipher/caesar"
	"github.com/cheveuxdelin/cipher/morse"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type atbashJSON struct {
	Text string `json:"text" binding:"required"`
}
type morseJSON struct {
	Text string
}
type caesarJSON struct {
	Text        string
	N           byte
	OnlyLetters bool
}

func atbashHandler(c *gin.Context) {
	var action string = c.Param("action")
	if action == "" {
		c.String(http.StatusBadRequest, "error")
		return
	}

	var json atbashJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.String(http.StatusBadRequest, "error")
		return
	}

	if action == "encode" {
		result, err := atbash.Encode(json.Text)
		if err != nil {
			c.String(400, "error")
			return
		}
		c.String(200, result)
	} else if action == "decode" {
		result, err := atbash.Encode(json.Text)
		if err != nil {
			c.String(400, "error")
			return
		}
		c.String(200, result)
	} else {
		c.String(400, "error")
		return
	}
}

func morseHandler(c *gin.Context) {
	var action string = c.Param("action")
	if action == "" {
		c.String(http.StatusBadRequest, "error")
		return
	}

	var json morseJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.String(http.StatusBadRequest, "error")
		return
	}

	if action == "encode" {
		result, err := morse.Encode(json.Text)
		if err != nil {
			c.String(400, "error")
			return
		}
		c.String(200, result)
	} else if action == "decode" {
		result, err := morse.Decode(json.Text)
		if err != nil {
			c.String(400, "error")
			return
		}
		c.String(200, result)
	} else {
		c.String(400, "error")
		return
	}
}

func caesarHandler(c *gin.Context) {
	var action string = c.Param("action")
	if action == "" {
		c.String(http.StatusBadRequest, "error")
		return
	}

	var json caesarJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.String(http.StatusBadRequest, "error")
		return
	}

	if action == "encode" {
		result, err := caesar.Encode(json.Text, json.N, json.OnlyLetters)
		if err != nil {
			c.String(400, "error")
			return
		}
		c.String(200, result)
	} else if action == "decode" {
		result, err := morse.Decode(json.Text)
		if err != nil {
			c.String(400, "error")
			return
		}
		c.String(200, result)
	} else {
		c.String(400, "error")
		return
	}
}

func main() {

	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/atbash/:action", atbashHandler)
	r.POST("/morse/:action", morseHandler)
	r.POST("/caesar/:action", morseHandler)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
