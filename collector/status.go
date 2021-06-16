package collector

import (
	"encoding/json"
	"log"
)

type Status struct {
	Process struct {
		Arch     string `json:"arch"`
		Platform string `json:"platform"`
		Pid      int    `json:"pid"`
	} `json:"process"`
	Epg struct {
		StoredEvents int `json:"storedEvents"`
	} `json:"epg"`
}

func NewStatusCollector() Status {
	api := newAPI()
	body := fetch(&api, "status")

	var status Status
	if err := json.Unmarshal(body, &status); err != nil {
		log.Fatal(err)
	}

	return status
}
