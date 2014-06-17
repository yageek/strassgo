package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	. "github.com/yageek/strassgo/tools"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func generateGeoJSON() {
	fmt.Println("Generate JSON...")
	if traffic := NewTraffic(GetSections()[:], GetInformations()[:]); traffic == nil {
		log.Fatal("Could not initialize kml")
	} else {
		traffic.TOGeoJson()
	}
}

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "strassgo", "jeremy")
	})

	m.Get("/traffic", func() (int, string) {

		if fileInfo, err := os.Stat("render/traffic.geojson"); os.IsNotExist(err) {
			fmt.Println("First time creating...")
			generateGeoJSON()
		} else {
			currentTime := time.Now()
			if currentTime.Sub(fileInfo.ModTime()).Seconds() > 180 {
				fmt.Println("File to old...")
				generateGeoJSON()
			}
		}
		data, _ := ioutil.ReadFile("render/traffic.geojson")
		return 200, string(data)
	})

	m.Run()
}
