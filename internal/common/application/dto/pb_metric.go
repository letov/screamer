package dto

import (
	"screamer/internal/common"
	"screamer/internal/common/domain"
	pb "screamer/proto"
)

func NewPbMetric(m domain.Metric) (*pb.Request, error) {
	var mType pb.MType
	var val float32 = 0
	var delta int64 = 0

	switch m.Ident.Type {
	case domain.Counter:
		mType = pb.MType_COUNTER
		delta = int64(m.Value)
	case domain.Gauge:
		mType = pb.MType_GAUGE
		val = float32(m.Value)
	default:
		return nil, common.ErrTypeNotExists
	}

	return &pb.Request{
		Id:    m.Ident.Name,
		Mtype: mType,
		Value: val,
		Delta: delta,
	}, nil
}
