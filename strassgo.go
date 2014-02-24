package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"os"
	"log"
	"io/ioutil"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
			r.HTML(200, "strassgo", "jeremy")
		})

	m.Get("/traffic",func()(int, string){

			if _, err := os.Stat("render/traffic.kml"); os.IsNotExist(err) {
				if traffic := NewTraffic(GetSections()[:],GetInformations()[:]); traffic == nil{
					log.Fatal("Could not initialize kml")
				}else{
					traffic.toKML()
				}
			}

			data , _ := ioutil.ReadFile("render/traffic.kml")
			return 200 , string(data)
		})

	m.Run()
}
