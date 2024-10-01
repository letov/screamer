package metricsources

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"screamer/internal/common/metric"
)

const totalMemory = "TotalMemory"
const freeMemory = "FreeMemory"
const cPUutilization1 = "CPUutilization1"

func getGopsutilMetric() []*metric.Metric {
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Counts(false)

	return []*metric.Metric{
		metric.NewGauge(totalMemory, float64(v.Total)),
		metric.NewGauge(freeMemory, float64(v.Available)),
		metric.NewGauge(cPUutilization1, float64(c)),
	}
}
