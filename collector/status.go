package collector

import (
	"encoding/json"
	"log"
	"strconv"

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
	time                           *prometheus.Desc
	version                        *prometheus.Desc
	processArch                    *prometheus.Desc
	processPlatform                *prometheus.Desc
	processVersionsNode            *prometheus.Desc
	processVersionsV8              *prometheus.Desc
	processVersionsUv              *prometheus.Desc
	processVersionsZlib            *prometheus.Desc
	processVersionsBrotli          *prometheus.Desc
	processVersionsAres            *prometheus.Desc
	processVersionsModules         *prometheus.Desc
	processVersionsNghttp2         *prometheus.Desc
	processVersionsNapi            *prometheus.Desc
	processVersionsLlhttp          *prometheus.Desc
	processVersionsOpenssl         *prometheus.Desc
	processVersionsCldr            *prometheus.Desc
	processVersionsIcu             *prometheus.Desc
	processVersionsTz              *prometheus.Desc
	processVersionsUnicode         *prometheus.Desc
	processEnvPath                 *prometheus.Desc
	processEnvUsingWinser          *prometheus.Desc
	processEnvNodeEnv              *prometheus.Desc
	processEnvServerConfigPath     *prometheus.Desc
	processEnvTunersConfigPath     *prometheus.Desc
	processEnvChannelsConfigPath   *prometheus.Desc
	processEnvServicesDbPath       *prometheus.Desc
	processEnvProgramsDbPath       *prometheus.Desc
	processPid                     *prometheus.Desc
	processMemoryUsageRss          *prometheus.Desc
	processMemoryUsageHeapTotal    *prometheus.Desc
	processMemoryUsageHeapUsed     *prometheus.Desc
	processMemoryUsageExternal     *prometheus.Desc
	processMemoryUsageArrayBuffers *prometheus.Desc
	epgGatheringNetworks           *prometheus.Desc
	epgStoredEvents                *prometheus.Desc
	streamCountTunerDevice         *prometheus.Desc
	streamCountTsFilter            *prometheus.Desc
	streamCountDecoder             *prometheus.Desc
	errorCountUncaughtException    *prometheus.Desc
	errorCountUnhandledRejection   *prometheus.Desc
	errorCountBufferOverflow       *prometheus.Desc
	errorCountTunerDeviceRespawn   *prometheus.Desc
	errorCountDecoderRespawn       *prometheus.Desc
	timerAccuracyLast              *prometheus.Desc
	timerAccuracyM1Avg             *prometheus.Desc
	timerAccuracyM1Min             *prometheus.Desc
	timerAccuracyM1Max             *prometheus.Desc
	timerAccuracyM5Avg             *prometheus.Desc
	timerAccuracyM5Min             *prometheus.Desc
	timerAccuracyM5Max             *prometheus.Desc
	timerAccuracyM15Avg            *prometheus.Desc
	timerAccuracyM15Min            *prometheus.Desc
	timerAccuracyM15Max            *prometheus.Desc
}

func NewStatusCollector() *statusCollector {
	return &statusCollector{
		time: prometheus.NewDesc(
			"mirakurun_status_time",
			"mirakurun Status Time",
			[]string{"host"},
			nil),
		version: prometheus.NewDesc(
			"mirukurun_status_version",
			"mirakurun Status Version",
			[]string{"host", "version"},
			nil),

		// Process
		processArch: prometheus.NewDesc(
			"mirukurun_status_process_arch",
			"mirukurun Status Process Arch",
			[]string{"host"},
			nil),
		processPlatform: prometheus.NewDesc(
			"mirakurun_status_process_platform",
			"mirukurun Status Process Platform",
			[]string{"host"},
			nil),
		processVersionsNode: prometheus.NewDesc(
			"mirukurun_status_process_versions_node",
			"mirukurun node version",
			[]string{"host", "version"},
			nil),
		processVersionsV8: prometheus.NewDesc(
			"mirukurun_status_process_versions_v8",
			"mirukurun v8 version",
			[]string{"host", "version"},
			nil),
		processVersionsUv: prometheus.NewDesc(
			"mirukurun_status_process_versions_uv",
			"mirukurun uv version",
			[]string{"host", "version"},
			nil),
		processVersionsZlib: prometheus.NewDesc(
			"mirukurun_status_process_versions_zlib",
			"mirukurun zlib version",
			[]string{"host", "version"},
			nil),
		processVersionsBrotli: prometheus.NewDesc(
			"mirukurun_status_process_versions_brotli",
			"mirukurun brotli version",
			[]string{"host", "version"},
			nil),
		processVersionsAres: prometheus.NewDesc(
			"mirukurun_status_process_versions_ares",
			"mirukurun Ares version",
			[]string{"host", "version"},
			nil),
		processVersionsModules: prometheus.NewDesc(
			"mirukurun_status_process_versions_modules",
			"mirukurun modules version",
			[]string{"host", "version"},
			nil),
		processVersionsNghttp2: prometheus.NewDesc(
			"mirukurun_status_process_versions_nghttp2",
			"mirukurun nghttp2 version",
			[]string{"host", "version"},
			nil),
		processVersionsNapi: prometheus.NewDesc(
			"mirukurun_status_process_versions_napi",
			"mirukurun napi version",
			[]string{"host", "version"},
			nil),
		processVersionsLlhttp: prometheus.NewDesc(
			"mirukurun_status_process_versions_llhttp",
			"mirukurun llhttp version",
			[]string{"host", "version"},
			nil),
		processVersionsOpenssl: prometheus.NewDesc(
			"mirukurun_status_process_versions_openssl",
			"mirukurun openssl version",
			[]string{"host", "version"},
			nil),
		processVersionsCldr: prometheus.NewDesc(
			"mirukurun_status_process_versions_cldr",
			"mirukurun cldr version",
			[]string{"host", "version"},
			nil),
		processVersionsIcu: prometheus.NewDesc(
			"mirukurun_status_process_versions_icu",
			"mirukurun icu version",
			[]string{"host", "version"},
			nil),
		processVersionsTz: prometheus.NewDesc(
			"mirukurun_status_process_versions_tz",
			"mirukurun tz version",
			[]string{"host", "version"},
			nil),
		processVersionsUnicode: prometheus.NewDesc(
			"mirukurun_status_process_versions_unicode",
			"mirukurun unicode version",
			[]string{"host", "version"},
			nil),
		processEnvPath: prometheus.NewDesc(
			"mirukurun_status_process_env_path",
			"mirukurun Status Process Env Path",
			[]string{"host", "path"},
			nil),
		processEnvUsingWinser: prometheus.NewDesc(
			"mirukurun_status_process_env_usingwinser",
			"mirukurun Status Process Env UsingWinser",
			[]string{"host"},
			nil),
		processEnvNodeEnv: prometheus.NewDesc(
			"mirukurun_status_process_env_nodeenv",
			"mirukurun Status Process Env NodeEnv",
			[]string{"host", "node_env"},
			nil),
		processEnvTunersConfigPath: prometheus.NewDesc(
			"mirukurun_status_process_env_tunersconfigpath",
			"mirukurun Status Process Env TunersConfigPath",
			[]string{"host", "tuners_config_path"},
			nil),
		processEnvServerConfigPath: prometheus.NewDesc(
			"mirukurun_status_process_env_serverconfigpath",
			"mirukurun Status Process Env ServerConfigPath",
			[]string{"host", "server_config_path"},
			nil),
		processEnvChannelsConfigPath: prometheus.NewDesc(
			"mirukurun_status_process_env_channelsconfigpath",
			"mirukurun Status Process Env ChannelsConfigPath",
			[]string{"host", "channels_config_path"},
			nil),
		processEnvServicesDbPath: prometheus.NewDesc(
			"mirukurun_status_process_env_servicesdbpath",
			"mirukurun Status Process Env ServicesDbPath",
			[]string{"host", "services_db_path"},
			nil),
		processEnvProgramsDbPath: prometheus.NewDesc(
			"mirukurun_status_process_env_programsdbpath",
			"mirukurun Status Process Env ProgramsDbPath",
			[]string{"host", "programs_db_path"},
			nil),
		processPid: prometheus.NewDesc(
			"mirukurun_status_process_pid",
			"mirukurun Status Process Pid",
			[]string{"host"},
			nil),
		processMemoryUsageRss: prometheus.NewDesc(
			"mirukurun_status_process_memoryusage_rss",
			"mirukurun Status Process MemoryUsage Rss",
			[]string{"host"},
			nil),
		processMemoryUsageHeapTotal: prometheus.NewDesc(
			"mirukurun_status_process_memoryusage_heaptotal",
			"mirukurun Status Process MemoryUsage HeapTotal",
			[]string{"host"},
			nil),
		processMemoryUsageHeapUsed: prometheus.NewDesc(
			"mirukurun_status_process_memoryusage_heapused",
			"mirukurun Status Process MemoryUsage HeapUsed",
			[]string{"host"},
			nil),
		processMemoryUsageExternal: prometheus.NewDesc(
			"mirukurun_status_process_memoryusage_external",
			"mirukurun Status Process MemoryUsage External",
			[]string{"host"},
			nil),
		processMemoryUsageArrayBuffers: prometheus.NewDesc(
			"mirukurun_status_process_memoryusage_arraybuffers",
			"mirukurun Status Process MemoryUsage ArrayBuffers",
			[]string{"host"},
			nil),

		// Epg
		epgGatheringNetworks: prometheus.NewDesc(
			"mirukurun_status_epg_gatheringnetworks",
			"mirukurun Status Epg GatheringNetworks",
			[]string{"host"},
			nil),
		epgStoredEvents: prometheus.NewDesc(
			"mirukurun_status_epg_storedevents",
			"mirukurun Status Epg StoredEvents",
			[]string{"host"},
			nil),

		// StreamCount
		streamCountTunerDevice: prometheus.NewDesc(
			"mirukurun_status_streamcount_tunerdevice",
			"mirukurun Status StreamCount TunerDevice",
			[]string{"host"},
			nil),
		streamCountTsFilter: prometheus.NewDesc(
			"mirukurun_status_streamcount_tsfilter",
			"mirukurun Status StreamCount TsFilter",
			[]string{"host"},
			nil),
		streamCountDecoder: prometheus.NewDesc(
			"mirukurun_status_streamcount_decoder",
			"mirukurun Status StreamCount Decoder",
			[]string{"host"},
			nil),

		// ErrorCount
		errorCountUncaughtException: prometheus.NewDesc(
			"mirukurun_status_errorcount_uncaughtexception",
			"mirukurun Status ErrorCount UncaughtException",
			[]string{"host"},
			nil),
		errorCountUnhandledRejection: prometheus.NewDesc(
			"mirukurun_status_errorcount_unhandledrejection",
			"mirukurun Status ErrorCount UnhandledRejection",
			[]string{"host"},
			nil),
		errorCountBufferOverflow: prometheus.NewDesc(
			"mirukurun_status_errorcount_bufferoverflow",
			"mirukurun Status ErrorCount BufferOverflow",
			[]string{"host"},
			nil),
		errorCountTunerDeviceRespawn: prometheus.NewDesc(
			"mirukurun_status_errorcount_tunerdevicerespawn",
			"mirukurun Status ErrorCount TunerDeviceRespawn",
			[]string{"host"},
			nil),
		errorCountDecoderRespawn: prometheus.NewDesc(
			"mirukurun_status_errorcount_decoderrespawn",
			"mirukurun Status ErrorCount DecoderRespawn",
			[]string{"host"},
			nil),

		// TimerAccuracy
		timerAccuracyLast: prometheus.NewDesc(
			"mirukurun_status_timeraccuracy_last",
			"mirukurun Status TimerAccuracy Last",
			[]string{"host"},
			nil),
		timerAccuracyM1Avg: prometheus.NewDesc(
			"mirukurun_status_timeraccuracy_m1_avg",
			"mirukurun Status TimerAccuracy M1 Avg",
			[]string{"host"},
			nil),
		timerAccuracyM1Min: prometheus.NewDesc(
			"mirukurun_status_timeraccuracy_m1_min",
			"mirukurun Status TimerAccuracy M1 Min",
			[]string{"host"},
			nil),
		timerAccuracyM1Max: prometheus.NewDesc(
			"mirukurun_status_timeraccuracy_m1_max",
			"mirukurun Status TimerAccuracy M1 Max",
			[]string{"host"},
			nil),
		timerAccuracyM5Avg: prometheus.NewDesc(
			"mirukurun_status_timeraccuracy_m5_avg",
			"mirukurun Status TimerAccuracy M5 Avg",
			[]string{"host"},
			nil),
		timerAccuracyM5Min: prometheus.NewDesc(
			"mirukurun_status_timeraccuracy_m5_min",
			"mirukurun Status TimerAccuracy M5 Min",
			[]string{"host"},
			nil),
		timerAccuracyM5Max: prometheus.NewDesc(
			"mirukurun_status_timeraccuracy_m5_max",
			"mirukurun Status TimerAccuracy M5 Max",
			[]string{"host"},
			nil),
		timerAccuracyM15Avg: prometheus.NewDesc(
			"mirukurun_status_timeraccuracy_m15_avg",
			"mirukurun Status TimerAccuracy M15 avg",
			[]string{"host"},
			nil),
		timerAccuracyM15Min: prometheus.NewDesc(
			"mirukurun_status_timeraccuracy_m15_min",
			"mirukurun Status TimerAccuracy M15 Min",
			[]string{"host"},
			nil),
		timerAccuracyM15Max: prometheus.NewDesc(
			"mirukurun_status_timeraccuracy_m15_max",
			"mirukurun Status TimerAccuracy M15 Max",
			[]string{"host"},
			nil),
	}
}

func (sc *statusCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- sc.time
	ch <- sc.version
	ch <- sc.processArch
	ch <- sc.processPlatform
	ch <- sc.processVersionsNode
	ch <- sc.processVersionsV8
	ch <- sc.processVersionsUv
	ch <- sc.processVersionsZlib
	ch <- sc.processVersionsBrotli
	ch <- sc.processVersionsAres
	ch <- sc.processVersionsModules
	ch <- sc.processVersionsNghttp2
	ch <- sc.processVersionsNapi
	ch <- sc.processVersionsLlhttp
	ch <- sc.processVersionsOpenssl
	ch <- sc.processVersionsCldr
	ch <- sc.processVersionsIcu
	ch <- sc.processVersionsTz
	ch <- sc.processVersionsUnicode
	ch <- sc.processEnvPath
	ch <- sc.processEnvUsingWinser
	ch <- sc.processEnvNodeEnv
	ch <- sc.processEnvTunersConfigPath
	ch <- sc.processEnvServerConfigPath
	ch <- sc.processEnvChannelsConfigPath
	ch <- sc.processEnvServicesDbPath
	ch <- sc.processEnvProgramsDbPath
	ch <- sc.processPid
	ch <- sc.processMemoryUsageRss
	ch <- sc.processMemoryUsageHeapTotal
	ch <- sc.processMemoryUsageHeapUsed
	ch <- sc.processMemoryUsageExternal
	ch <- sc.processMemoryUsageArrayBuffers
	ch <- sc.epgGatheringNetworks
	ch <- sc.epgStoredEvents
	ch <- sc.streamCountTunerDevice
	ch <- sc.streamCountTsFilter
	ch <- sc.streamCountDecoder
	ch <- sc.errorCountUncaughtException
	ch <- sc.errorCountUnhandledRejection
	ch <- sc.errorCountBufferOverflow
	ch <- sc.errorCountTunerDeviceRespawn
	ch <- sc.errorCountDecoderRespawn
	ch <- sc.timerAccuracyLast
	ch <- sc.timerAccuracyM1Avg
	ch <- sc.timerAccuracyM1Min
	ch <- sc.timerAccuracyM1Max
	ch <- sc.timerAccuracyM5Avg
	ch <- sc.timerAccuracyM5Min
	ch <- sc.timerAccuracyM5Max
	ch <- sc.timerAccuracyM15Avg
	ch <- sc.timerAccuracyM15Min
	ch <- sc.timerAccuracyM15Max
}

func (sc *statusCollector) Collect(ch chan<- prometheus.Metric) {
	api := newAPI()
	body := fetch(&api, "status")

	var status Status
	if err := json.Unmarshal(body, &status); err != nil {
		log.Fatal(err)
	}

	usingWinser, _ := strconv.ParseFloat(status.Process.Env.UsingWinser, 64)

	// メトリクス達
	ch <- prometheus.MustNewConstMetric(
		sc.time,
		prometheus.CounterValue,
		float64(status.Time),
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.version,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Version)

	// Process
	ch <- prometheus.MustNewConstMetric(
		sc.processArch,
		prometheus.GaugeValue,
		1,
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.processPlatform,
		prometheus.GaugeValue,
		1,
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.processVersionsNode,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Versions.Node)
	ch <- prometheus.MustNewConstMetric(
		sc.processVersionsV8,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Versions.V8)
	ch <- prometheus.MustNewConstMetric(
		sc.processVersionsUv,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Versions.Uv)
	ch <- prometheus.MustNewConstMetric(
		sc.processVersionsZlib,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Versions.Zlib)
	ch <- prometheus.MustNewConstMetric(
		sc.processVersionsBrotli,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Versions.Brotli)
	ch <- prometheus.MustNewConstMetric(
		sc.processVersionsAres,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Versions.Ares)
	ch <- prometheus.MustNewConstMetric(
		sc.processVersionsModules,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Versions.Modules)
	ch <- prometheus.MustNewConstMetric(
		sc.processVersionsNghttp2,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Versions.Nghttp2)
	ch <- prometheus.MustNewConstMetric(
		sc.processVersionsNapi,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Versions.Napi)
	ch <- prometheus.MustNewConstMetric(
		sc.processVersionsLlhttp,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Versions.Llhttp)
	ch <- prometheus.MustNewConstMetric(
		sc.processVersionsOpenssl,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Versions.Openssl)
	ch <- prometheus.MustNewConstMetric(
		sc.processVersionsCldr,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Versions.Cldr)
	ch <- prometheus.MustNewConstMetric(
		sc.processVersionsIcu,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Versions.Icu)
	ch <- prometheus.MustNewConstMetric(
		sc.processVersionsTz,
		prometheus.GaugeValue,
		1,
		api.Host, status.Process.Versions.Tz)
	ch <- prometheus.MustNewConstMetric(
		sc.processVersionsUnicode,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Versions.Unicode)
	ch <- prometheus.MustNewConstMetric(
		sc.processEnvPath,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Env.Path)
	ch <- prometheus.MustNewConstMetric(
		sc.processEnvUsingWinser,
		prometheus.GaugeValue,
		usingWinser,
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.processEnvNodeEnv,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Env.NodeEnv)
	ch <- prometheus.MustNewConstMetric(
		sc.processEnvTunersConfigPath,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Env.TunersConfigPath)
	ch <- prometheus.MustNewConstMetric(
		sc.processEnvServerConfigPath,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Env.ServerConfigPath)
	ch <- prometheus.MustNewConstMetric(
		sc.processEnvChannelsConfigPath,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Env.ChannelsConfigPath)
	ch <- prometheus.MustNewConstMetric(
		sc.processEnvServicesDbPath,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Env.ServicesDbPath)
	ch <- prometheus.MustNewConstMetric(
		sc.processEnvProgramsDbPath,
		prometheus.GaugeValue,
		1,
		api.Host,
		status.Process.Env.ProgramsDbPath)
	ch <- prometheus.MustNewConstMetric(
		sc.processPid,
		prometheus.GaugeValue,
		float64(status.Process.Pid),
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.processMemoryUsageRss,
		prometheus.GaugeValue,
		float64(status.Process.MemoryUsage.Rss),
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.processMemoryUsageHeapTotal,
		prometheus.GaugeValue,
		float64(status.Process.MemoryUsage.HeapTotal),
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.processMemoryUsageHeapUsed,
		prometheus.GaugeValue,
		float64(status.Process.MemoryUsage.HeapUsed),
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.processMemoryUsageExternal,
		prometheus.GaugeValue,
		float64(status.Process.MemoryUsage.External),
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.processMemoryUsageArrayBuffers,
		prometheus.GaugeValue,
		float64(status.Process.MemoryUsage.ArrayBuffers),
		api.Host)

	// Epg
	ch <- prometheus.MustNewConstMetric(
		sc.epgGatheringNetworks,
		prometheus.GaugeValue,
		float64(status.Epg.StoredEvents),
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.epgStoredEvents,
		prometheus.GaugeValue,
		float64(status.Epg.StoredEvents),
		api.Host)

	// StreamCount
	ch <- prometheus.MustNewConstMetric(
		sc.streamCountTunerDevice,
		prometheus.GaugeValue,
		float64(status.StreamCount.TunerDevice),
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.streamCountTsFilter,
		prometheus.GaugeValue,
		float64(status.StreamCount.TsFilter),
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.streamCountDecoder,
		prometheus.GaugeValue,
		float64(status.StreamCount.Decoder),
		api.Host)

	// ErrorCount
	ch <- prometheus.MustNewConstMetric(sc.errorCountUncaughtException,
		prometheus.GaugeValue,
		float64(status.ErrorCount.UncaughtException),
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.errorCountUnhandledRejection,
		prometheus.GaugeValue,
		float64(status.ErrorCount.UnhandledRejection),
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.errorCountBufferOverflow,
		prometheus.GaugeValue,
		float64(status.ErrorCount.BufferOverflow),
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.errorCountTunerDeviceRespawn,
		prometheus.GaugeValue,
		float64(status.ErrorCount.TunerDeviceRespawn),
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.errorCountDecoderRespawn,
		prometheus.GaugeValue,
		float64(status.ErrorCount.DecoderRespawn),
		api.Host)

	// TimerAccuracy
	ch <- prometheus.MustNewConstMetric(
		sc.timerAccuracyLast,
		prometheus.GaugeValue,
		status.TimerAccuracy.Last,
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.timerAccuracyM1Avg,
		prometheus.GaugeValue,
		status.TimerAccuracy.M1.Avg,
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.timerAccuracyM1Min,
		prometheus.GaugeValue,
		status.TimerAccuracy.M1.Min,
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.timerAccuracyM1Max,
		prometheus.GaugeValue,
		status.TimerAccuracy.M1.Max,
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.timerAccuracyM5Avg,
		prometheus.GaugeValue,
		status.TimerAccuracy.M5.Avg,
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.timerAccuracyM5Min,
		prometheus.GaugeValue,
		status.TimerAccuracy.M5.Min,
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.timerAccuracyM5Max,
		prometheus.GaugeValue,
		status.TimerAccuracy.M5.Max,
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.timerAccuracyM15Avg,
		prometheus.GaugeValue,
		status.TimerAccuracy.M15.Avg,
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.timerAccuracyM15Min,
		prometheus.GaugeValue,
		status.TimerAccuracy.M15.Min,
		api.Host)
	ch <- prometheus.MustNewConstMetric(
		sc.timerAccuracyM15Max,
		prometheus.GaugeValue,
		status.TimerAccuracy.M15.Max,
		api.Host)
}
