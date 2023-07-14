package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.GET("/komikita-string", func(a echo.Context) error {
		title := a.QueryParam("Nekomata")
		description := a.QueryParam("Seorang Manusia Kucing bernama okayu yang menjadi petualang")
		total_eps := a.QueryParam("150")
		genre := a.QueryParam("Action, Adventure, Comedy")
		return a.String(http.StatusOK, title+" "+description+""+total_eps+""+genre)
	})

	e.GET("komikita-html", func(b echo.Context) error {
		title := b.QueryParam("title")
		description := b.QueryParam("description")
		total_eps := b.QueryParam("total_eps")
		genre := b.QueryParam("genre")
		html := ""
		html += "title : " + title
		html += "<br>"
		html += "description : " + description
		html += "total_eps : " + total_eps
		html += "genre : " + genre
		return b.HTML(http.StatusOK, html)
	})

	e.GET("/komikita-jsom", func(c echo.Context) error {
		titile := c.QueryParam("title")            // string
		description := c.QueryParam("descriptiom") // string
		total_eps := c.QueryParam("total_eps")
		genre := c.QueryParam("genre")
		/*
		   {
		       "name" : "Adi",
		       "nim" : "12344"
		   }
		*/
		// interface bisa nerima data apa aja
		// int
		json := make(map[string]interface{})
		json["title"], _ = strconv.Atoi(titile)
		json["description"], _ = strconv.Atoi(description) // convert to integer
		json["total_eps"], _ = strconv.Atoi(total_eps)
		json["genre"], _ = strconv.Atoi(genre)
		return c.JSON(http.StatusOK, json)
	})

	e.POST("/mahasiswa", func(d echo.Context) error {
		title := d.FormValue("title")
		description := d.FormValue("description")
		return d.String(http.StatusOK, title+" "+description)
	})

	e.Start(":9999")
}
