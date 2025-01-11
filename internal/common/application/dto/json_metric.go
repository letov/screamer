package dto

import (
	"screamer/internal/common"
	"screamer/internal/common/domain"
	pb "screamer/proto"
	"strings"
)

// JsonMetric метрика в json (формат структуры из задания)
type JsonMetric struct {
	ID    string   `json:"id"`
	MType string   `json:"type"`
	Delta *int64   `json:"delta,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

func (jm JsonMetric) GetDomainMetric() (domain.Metric, error) {
	switch domain.Type(jm.MType) {
	case domain.Counter:
		if jm.Delta == nil {
			return domain.Metric{}, common.ErrEmptyValue
		}
		return domain.NewMetric(jm.ID, float64(*jm.Delta), domain.Counter)
	case domain.Gauge:
		if jm.Value == nil {
			return domain.Metric{}, common.ErrEmptyValue
		}
		return domain.NewMetric(jm.ID, *jm.Value, domain.Gauge)
	default:
		return domain.Metric{}, common.ErrMetricNotExists
	}
}

func (jm JsonMetric) GetIdent() (domain.Ident, error) {
	m, err := jm.GetDomainMetric()
	if err != nil {
		return domain.Ident{}, err
	}
	return m.Ident, nil
}

func NewJsonMetricFromPb(pb *pb.Request) JsonMetric {
	delta := pb.GetDelta()
	value := float64(pb.GetValue())

	return JsonMetric{
		ID:    pb.GetId(),
		MType: strings.ToLower(pb.GetMtype().String()),
		Delta: &delta,
		Value: &value,
	}
}

func NewJsonMetric(m domain.Metric) (JsonMetric, error) {
	switch m.Ident.Type {
	case domain.Counter:
		v := int64(m.Value)
		return JsonMetric{
			ID:    m.Ident.Name,
			MType: m.Ident.Type.String(),
			Delta: &v,
			Value: nil,
		}, nil
	case domain.Gauge:
		return JsonMetric{
			ID:    m.Ident.Name,
			MType: m.Ident.Type.String(),
			Delta: nil,
			Value: &m.Value,
		}, nil
	default:
		return JsonMetric{}, common.ErrTypeNotExists
	}
}
