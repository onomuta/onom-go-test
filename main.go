package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

type C struct {
  Id int
  Name string
}

func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*.tmpl")
  router.GET("/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "a": "a",
      "b": []string{"b_todo1","b_todo2"},
      "c": []C{{1,"c_mika"},{2,"c_risa"}},
      "d": C{3,"d_mayu"},
      "e": true,
      "f": false,
      "h": true,
    })
  })
  router.Run(":8080")
}