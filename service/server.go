package service

import (
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/unrolled/render"
)

var port string

// NewServer defined
func NewServer(Port string) *martini.ClassicMartini {
	m := martini.Classic()
	//运行程序监听端口
	port = Port
	return m
}

// RunServer run the server
func RunServer(server *martini.ClassicMartini) {
	rd := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html", ".tmpl"},
		IndentJSON: true,
	})
	server.Use(martini.Static("assets"))
	server.Get("/json", func(w http.ResponseWriter, req *http.Request) {
		rd.JSON(w, http.StatusOK, struct {
			ID      string
			Content string
		}{ID: "1331310", Content: "Hello cloudgo!"})
	})

	server.Get("/login", func(res http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		rd.HTML(res, http.StatusOK, "login", " ")
	})
	server.Post("/login", func(res http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		err := rd.HTML(res, http.StatusOK, "login_ok", struct {
			Username string
			Password string
		}{
			Username: r.Form["username"][0],
			Password: r.Form["password"][0]})
		if err != nil {
			panic("render fail!")
		}

		fmt.Println("username:", r.Form["username"][0])
		fmt.Println("password:", r.Form["password"][0])
	})

	server.Get("/unknown", func(rw http.ResponseWriter, r *http.Request) {
		http.Error(rw, "501 Not Implemented", http.StatusNotImplemented)
	})

	server.RunOnAddr(":" + port)
}
