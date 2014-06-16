package tools

import "testing"

func TestTrafficNew(t *testing.T) {

	if traffic := NewTraffic(GetSections()[:], GetInformations()[:]); traffic == nil {
		t.Errorf("Could not initialize data")
	} else {
		traffic.ToKML()
	}

}
