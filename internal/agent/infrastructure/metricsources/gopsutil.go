package metricsources

import (
	"screamer/internal/common/domain"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

const totalMemory = "TotalMemory"
const freeMemory = "FreeMemory"
const cPUutilization1 = "CPUutilization1"

func getGopsutilMetric() []domain.Metric {
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Counts(false)

	m1, _ := domain.NewMetric(totalMemory, float64(v.Total), domain.Gauge)
	m2, _ := domain.NewMetric(freeMemory, float64(v.Available), domain.Gauge)
	m3, _ := domain.NewMetric(cPUutilization1, float64(c), domain.Gauge)

	return []domain.Metric{m1, m2, m3}
}
