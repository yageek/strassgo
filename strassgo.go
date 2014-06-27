package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
	. "github.com/yageek/strassgo/tools"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var lastRefresh time.Time

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
		r.HTML(200, "strassgo", nil)
	})

	m.Get("/traffic", func() (int, string) {

		if fileInfo, err := os.Stat("render/traffic.geojson"); os.IsNotExist(err) {
			fmt.Println("First time creating...")
			generateGeoJSON()
		} else {
			currentTime := time.Now()
			if currentTime.Sub(fileInfo.ModTime()).Seconds() > 180 {
				fmt.Println("File to old...")
				lastRefresh = time.Now()
				generateGeoJSON()
			}
		}
		data, _ := ioutil.ReadFile("render/traffic.geojson")
		return 200, string(data)
	})

	m.Run()
}
