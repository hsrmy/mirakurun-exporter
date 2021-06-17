package collector

import (
	"encoding/json"
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

// JSON展開用
type Status struct {
	Time    int64  `json:"time"`
	Version string `json:"version"`
	Process struct {
		Arch     string `json:"arch"`
		Platform string `json:"platform"`
		Versions struct {
			Node    string `json:"node"`
			V8      string `json:"v8"`
			Uv      string `json:"uv"`
			Zlib    string `json:"zlib"`
			Brotli  string `json:"brotli"`
			Ares    string `json:"ares"`
			Modules string `json:"modules"`
			Nghttp2 string `json:"nghttp2"`
			Napi    string `json:"napi"`
			Llhttp  string `json:"llhttp"`
			Openssl string `json:"openssl"`
			Cldr    string `json:"cldr"`
			Icu     string `json:"icu"`
			Tz      string `json:"tz"`
			Unicode string `json:"unicode"`
		} `json:"versions"`
		Env struct {
			Path               string `json:"PATH"`
			UsingWinser        string `json:"USING_WINSER"`
			NodeEnv            string `json:"NODE_ENV"`
			ServerConfigPath   string `json:"SERVER_CONFIG_PATH"`
			TunersConfigPath   string `json:"TUNERS_CONFIG_PATH"`
			ChannelsConfigPath string `json:"CHANNELS_CONFIG_PATH"`
			ServicesDbPath     string `json:"SERVICES_DB_PATH"`
			ProgramsDbPath     string `json:"PROGRAMS_DB_PATH"`
		} `json:"env"`
		Pid         int `json:"pid"`
		MemoryUsage struct {
			Rss          int `json:"rss"`
			HeapTotal    int `json:"heapTotal"`
			HeapUsed     int `json:"heapUsed"`
			External     int `json:"external"`
			ArrayBuffers int `json:"arrayBuffers"`
		} `json:"memoryUsage"`
	} `json:"process"`
	Epg struct {
		GatheringNetworks []int `json:"gatheringNetworks"`
		StoredEvents      int   `json:"storedEvents"`
	} `json:"epg"`
	StreamCount struct {
		TunerDevice int `json:"tunerDevice"`
		TsFilter    int `json:"tsFilter"`
		Decoder     int `json:"decoder"`
	} `json:"streamCount"`
	ErrorCount struct {
		UncaughtException  int `json:"uncaughtException"`
		UnhandledRejection int `json:"unhandledRejection"`
		BufferOverflow     int `json:"bufferOverflow"`
		TunerDeviceRespawn int `json:"tunerDeviceRespawn"`
		DecoderRespawn     int `json:"decoderRespawn"`
	} `json:"errorCount"`
	TimerAccuracy struct {
		Last float64 `json:"last"`
		M1   struct {
			Avg float64 `json:"avg"`
			Min float64 `json:"min"`
			Max float64 `json:"max"`
		} `json:"m1"`
		M5 struct {
			Avg float64 `json:"avg"`
			Min float64 `json:"min"`
			Max float64 `json:"max"`
		} `json:"m5"`
		M15 struct {
			Avg float64 `json:"avg"`
			Min float64 `json:"min"`
			Max float64 `json:"max"`
		} `json:"m15"`
	} `json:"timerAccuracy"`
}

type statusCollector struct {
	time            *prometheus.Desc
	version         *prometheus.Desc
	processArch     *prometheus.Desc
	processPlatform *prometheus.Desc
}

func NewStatusCollector() *statusCollector {
	return &statusCollector{
		time:            prometheus.NewDesc("mirakurun_status_time", "Time of status api", nil, nil),
		version:         prometheus.NewDesc("mirukurun_status_version", "Version of status api", []string{"host"}, nil),
		processArch:     prometheus.NewDesc("mirukurun_status_process_arch", "mirukurun process arch", []string{"host"}, nil),
		processPlatform: prometheus.NewDesc("mirakurun_status_process_platform", "mirukurun process platform", []string{"host"}, nil),
	}
}

func (sc *statusCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- sc.time
	ch <- sc.version
	ch <- sc.processArch
	ch <- sc.processPlatform
}

func (sc *statusCollector) Collect(ch chan<- prometheus.Metric) {
	api := newAPI()
	body := fetch(&api, "status")

	var status Status
	if err := json.Unmarshal(body, &status); err != nil {
		log.Fatal(err)
	}

	ch <- prometheus.MustNewConstMetric(sc.time, prometheus.CounterValue, float64(status.Time))
	ch <- prometheus.MustNewConstMetric(sc.version, prometheus.GaugeValue, 1, api.Host)
	ch <- prometheus.MustNewConstMetric(sc.processArch, prometheus.GaugeValue, 1, api.Host)
	ch <- prometheus.MustNewConstMetric(sc.processPlatform, prometheus.GaugeValue, 1, api.Host)
}
