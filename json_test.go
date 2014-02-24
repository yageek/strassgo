package main

import (
	"testing"
	"io/ioutil"
)

func TestDownloadSection(t *testing.T){

	_, err := downloadJSON(JSONSectionUrl)
	if err != nil {
		t.Error(" Error in downloading JSON");
	}
}

func TestUnmarshalSection(t *testing.T){

	json_file ,err := ioutil.ReadFile("section_test.json")

	if err != nil{
		t.Error("Could not open json file")
	}

	var sections []SectionJSON = unmarshalSectionRawData(json_file)
	if len(sections) == 0 {
		t.Error("Sections should not have zero size")
	}
}



