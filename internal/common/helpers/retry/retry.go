package retry

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
)

// NewRetryJob ретрайбл джоб, tryIntervals - интервалы повторов
func NewRetryJob[T any](
	ctx context.Context,
	name string,
	job func(context.Context) (T, error),
	retryErrors []error,
	tryIntervals []int,
	log *zap.SugaredLogger,
) (res T, err error) {
	for inx, i := range tryIntervals {
		res, err = job(ctx)
		if err == nil {
			return
		}
		d := time.Duration(i) * time.Second
		log.Warnf("%s fail on try %d (%s)", name, inx, err)
		if !isRetryError(err, retryErrors) || inx >= len(tryIntervals) {
			return
		}
		time.Sleep(d)
	}
	log.Warn(name, " fail after all tries")
	return
}

func isRetryError(err error, retryErrors []error) bool {
	if len(retryErrors) == 0 {
		return true
	}
	for _, e := range retryErrors {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}
