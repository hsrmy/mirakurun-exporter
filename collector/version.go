package collector

import (
	"encoding/json"
	"log"
)

type Version struct {
	Current string `json:"current"`
	Latest  string `json:"latest"`
}

func NewVersionCollector() Version {
	api := newAPI()
	body := fetch(&api, "version")

	var version Version
	if err := json.Unmarshal(body, &version); err != nil {
		log.Fatal(err)
	}

	return version

}
