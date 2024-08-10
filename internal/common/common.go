package common

var RuntimeMetrics = []string{
	"Alloc",
	"BuckHashSys",
	"Frees",
	"GCCPUFraction",
	"GCSys",
	"HeapAlloc",
	"HeapIdle",
	"HeapInuse",
	"HeapObjects",
	"HeapReleased",
	"HeapSys",
	"LastGC",
	"Lookups",
	"MCacheInuse",
	"MCacheSys",
	"MSpanInuse",
	"MSpanSys",
	"Mallocs",
	"NextGC",
	"NumForcedGC",
	"NumGC",
	"OtherSys",
	"PauseTotalNs",
	"StackInuse",
	"StackSys",
	"Sys",
	"TotalAlloc",
}

const RandomMetric = "RandomValue"
const PollCountMetric = "PollCount"

func GetGaugeInit() *[]string {
	a := append(RuntimeMetrics, RandomMetric)
	return &a
}

func GetCounterInit() *[]string {
	a := []string{PollCountMetric}
	return &a
}

func GetAllInit() *[]string {
	a := append(*GetGaugeInit(), *GetCounterInit()...)
	return &a
}
