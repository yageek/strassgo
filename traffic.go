package main

import (
	"encoding/xml"
	"strconv"
	"io/ioutil"
)

type Traffic struct {

	sections map[string]SectionJSON
	informations map[string]InformationJSON
}


func NewTraffic(sections []SectionJSON, informations[]InformationJSON) *Traffic{

	traffic := new(Traffic)

	traffic.sections = make(map[string]SectionJSON)
	traffic.informations = make(map[string]InformationJSON)

	for i := range sections{
		traffic.sections[sections[i].Id] = sections[i]
	}

	for i := range informations{
		traffic.informations[informations[i].Id] = informations[i]
	}

	return traffic
}

func (traffic *Traffic) toKML(){

	kml := &KMLFile{Namespace: "http://earth.google.com/kml/2.1"}
	kml.Document = KMLDocument{Name: "StrassGO",Description: "Traffic en temps réels"}

	redStyle := KMLStyle{Id: "redLine"}
	redStyle.LineStyle = KMLLineStyle{Color: SectionRed, Width: 4}

	yellowStyle := KMLStyle{Id: "yellowLine"}
	yellowStyle.LineStyle = KMLLineStyle{Color: SectionYellow, Width: 4}

	greenStyle := KMLStyle{Id: "greenLine"}
	greenStyle.LineStyle = KMLLineStyle{Color: SectionGreen, Width: 4}

	kml.Document.Styles = []KMLStyle{redStyle,yellowStyle,greenStyle}

	for _, v := range traffic.sections{

		placemark := KMLPlacemark{Name: v.Id, StyleUrl:SectionGreen}

		var coordinates string = ""

		for _, line := range v.SectionPoints{

			x,_ := line.X.Float64()
			y,_ := line.Y.Float64()

			coordinates+= strconv.FormatFloat(x,'f',13,64) + ","
			coordinates+= strconv.FormatFloat(y,'f',13,64) + ",0\n"
		}

		lineString := KMLLineString{Coordinates: coordinates}

		placemark.LineStrings = append(placemark.LineStrings,lineString)

		kml.Document.PlaceMarks = append(kml.Document.PlaceMarks, placemark)

	}

	xmlData, _ := xml.MarshalIndent(kml, ""," ")

	err := ioutil.WriteFile("render/traffic.kml", append([]byte(xml.Header),xmlData...),0666)

	if err != nil {
		panic(err)
	}
}
