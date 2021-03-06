package tools

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	JSONSectionUrl = "http://carto.strasmap.eu/remote.amf.json/TraficInfo.geometry"
	JSONInfoUrl    = "http://carto.strasmap.eu/remote.amf.json/TraficInfo.status"
	JsonObjectsKey = "d"
)

type SectionPointJSON struct {
	X json.Number `json:"x"`
	Y json.Number `json:"y"`
}
type SectionJSON struct {
	Id            string             `json:"id"`
	SectionPoints []SectionPointJSON `json:"go"`
}

type InformationJSON struct {
	Id        string `json:"id"`
	SiracCode string `json:"lw"`
	ColorCode string `json:"lc"`
}

func downloadJSON(url string) ([]byte, error) {

	log.Printf("Downloading data at %s", url)
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode != 200 {
		log.Printf("An error occured during download at %s - err : %v", url, err)
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		return data, nil
	}

}

func unmarshalSectionRawData(raw_data []byte) []SectionJSON {

	log.Printf("Unmarshaling Section data...")
	var objMap map[string][]json.RawMessage
	var sections []SectionJSON

	if err := json.Unmarshal(raw_data, &objMap); err != nil {
		log.Fatal(err)
		return nil
	}

	for _, obj := range objMap[JsonObjectsKey] {

		var dst SectionJSON
		if err := json.Unmarshal(([]byte)(obj), &dst); err != nil {
			log.Fatal(err)
			return nil
		}
		sections = append(sections, dst)
	}
	return sections
}

func unmarshalInformationsRawData(raw_data []byte) []InformationJSON {

	log.Printf("Unmarshaling Informations data...")
	var objMap map[string]json.RawMessage
	var informations []InformationJSON

	if err := json.Unmarshal(raw_data, &objMap); err != nil {
		log.Fatal(err)
		return nil
	}

	if err := json.Unmarshal(objMap[JsonObjectsKey], &informations); err != nil {
		log.Fatal(err)
		return nil
	}

	return informations
}

func GetSections() []SectionJSON {

	if data, error := downloadJSON(JSONSectionUrl); error != nil {
		return nil
	} else {
		return unmarshalSectionRawData(data)
	}
}

func GetInformations() []InformationJSON {

	if data, error := downloadJSON(JSONInfoUrl); error != nil {
		return nil
	} else {
		return unmarshalInformationsRawData(data)
	}
}
