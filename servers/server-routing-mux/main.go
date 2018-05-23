package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"os"
)

type HeroModel struct {
	Name string
	Herotype string
	superpowers bool
}

type PagesModel struct {
	State    bool
	Message string
}

func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/first", first)
	mux.GET("/second", second)
	mux.GET("/user/:hero", hero)
	mux.GET("/blog/:heroType/:name", heroType)
	mux.POST("/blog/:heroType/:name", newHero)
	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	p := PagesModel{
		State: true,
		Message: "Home Page",
	}

	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(bs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)
}

func first(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	p := PagesModel{
		State: true,
		Message: "First Page",
	}

	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(bs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)
}

func second(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	p := PagesModel{
		State: true,
		Message: "Second Page",
	}

	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(bs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)
}

func hero(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	p := HeroModel{
		Name: ps.ByName("hero"),
		Herotype: "good",
		superpowers : true,
	}

	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(bs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)
}

func heroType(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	p := HeroModel{
		Name: ps.ByName("hero"),
		Herotype: ps.ByName("name"),
		superpowers : true,
	}

	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(bs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)
}

func newHero(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	p := HeroModel{
		Name: ps.ByName("hero"),
		Herotype: ps.ByName("name"),
		superpowers : true,
	}

	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(bs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)
}