package collector

import (
	"encoding/json"
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type versionCollector struct {
	currentVersion *prometheus.Desc
	latestVersion  *prometheus.Desc
}

// JSON展開用
type Version struct {
	Current string `json:"current"`
	Latest  string `json:"latest"`
}

func NewVersionCollector() *versionCollector {
	return &versionCollector{
		currentVersion: prometheus.NewDesc(
			"mirukurun_current_version",
			"Current Version of mirakurun.",
			[]string{"host", "version"},
			nil),
		latestVersion: prometheus.NewDesc(
			"mirakurun_latest_version",
			"Latest version of mirakurun.",
			[]string{"host", "version"},
			nil),
	}
}

func (vc *versionCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- vc.currentVersion
	ch <- vc.latestVersion
}

func (vc *versionCollector) Collect(ch chan<- prometheus.Metric) {
	api := newAPI()
	body := fetch(&api, "version")
	var version Version
	if err := json.Unmarshal(body, &version); err != nil {
		log.Fatal(err)
	}

	ch <- prometheus.MustNewConstMetric(vc.currentVersion, prometheus.GaugeValue, 1, api.Host, version.Current)
	ch <- prometheus.MustNewConstMetric(vc.latestVersion, prometheus.GaugeValue, 1, api.Host, version.Latest)
}
