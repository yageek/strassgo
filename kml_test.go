package main

import (
	"testing"
	"os"
	"encoding/xml"
	"fmt"
)



func TestMarshal(t *testing.T){

	kml := &KMLFile{Namespace: "http://earth.google.com/kml/2.1"}
	kml.Document = KMLDocument{Name: "StrassGO",Description: "Traffic en temps réels"}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent(" ","   ")
	if err := enc.Encode(kml); err != nil {
		fmt.Printf("error:%v\n",err)
	}
}
