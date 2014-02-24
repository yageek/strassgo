package main

import "encoding/xml"

const (
	SectionRed    = "0xBF0000"
	SectionYellow = "0xFFFF00"
	SectionGreen  = "0x00CC33"
)

type KMLFile struct {
	XMLName   xml.Name `xml:"kml"`
	Namespace string `xml:"xmlns,attr"`
	Document  KMLDocument `xml:"Document"`
}

type KMLDocument struct {
	Name        string `xml:"name,omitempty"`
	Description string `xml:"description,omitempty"`
	Styles      []KMLStyle `xml:"Style,omitempty"`
	PlaceMarks  []KMLPlacemark `xml:"Placemark"`

}

type KMLStyle struct {
	Id        string `xml:"id,attr"`
	LineStyle KMLLineStyle `xml:"LineStyle"`

}

type KMLLineStyle struct {
	Color string `xml:"color"`
	Width int `xml:"width"`
}

type KMLPlacemark struct {
	Name        string `xml:"name"`
	StyleUrl    string `xml:"StyleUrl"`
	LineStrings []KMLLineString `xml:"LineString"`
}

type KMLLineString struct {
	Coordinates string `xml:"coordinates"`
}


