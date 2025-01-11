package test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"screamer/internal/common/domain/metric"
	"screamer/internal/server/db"
	"screamer/internal/server/repositories"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func Test_updateCounterHandler(t *testing.T) {
	type args struct {
		t     metric.Type
		name  string
		value int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "positive test #1",
			args: args{
				t:     metric.Counter,
				name:  "test",
				value: 100,
			},
		},
		{
			name: "positive test #2",
			args: args{
				t:     metric.Counter,
				name:  "test2",
				value: 1232,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest(t, func(mux *chi.Mux, repo repositories.Repository, db *db.DB) {
				ctx := context.Background()
				_ = flushDB(ctx, db)

				data, _ := json.Marshal(metric.JSONMetric{
					ID:    tt.args.name,
					MType: tt.args.t.String(),
					Delta: &tt.args.value,
				})

				req, err := http.NewRequest("POST", "/update", bytes.NewBuffer(data))
				if err != nil {
					t.Fatal(err)
				}
				rr := httptest.NewRecorder()
				mux.ServeHTTP(rr, req)
				assert.Equal(t, rr.Code, http.StatusOK)
				m, err := repo.Get(ctx, metric.Ident{
					Type: tt.args.t,
					Name: tt.args.name,
				})
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, m.Value, float64(tt.args.value))
			})
		})
	}
}

func Test_updateGaugeHandler(t *testing.T) {
	type args struct {
		t     metric.Type
		name  string
		value float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "positive test #1",
			args: args{
				t:     metric.Gauge,
				name:  "test",
				value: 100,
			},
		},
		{
			name: "positive test #2",
			args: args{
				t:     metric.Gauge,
				name:  "test2",
				value: 1232.22,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest(t, func(mux *chi.Mux, repo repositories.Repository, db *db.DB) {
				ctx := context.Background()
				_ = flushDB(ctx, db)

				data, _ := json.Marshal(metric.JSONMetric{
					ID:    tt.args.name,
					MType: tt.args.t.String(),
					Value: &tt.args.value,
				})

				req, err := http.NewRequest("POST", "/update", bytes.NewBuffer(data))
				if err != nil {
					t.Fatal(err)
				}
				rr := httptest.NewRecorder()
				mux.ServeHTTP(rr, req)
				assert.Equal(t, rr.Code, http.StatusOK)
				m, err := repo.Get(ctx, metric.Ident{
					Type: tt.args.t,
					Name: tt.args.name,
				})
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, m.Value, tt.args.value)
			})
		})
	}
}

func Test_updateNegativeTypeHandler(t *testing.T) {
	type args struct {
		t     string
		name  string
		value float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "negative test #1",
			args: args{
				t:     "unknownType1",
				name:  "test1",
				value: 100,
			},
		},
		{
			name: "negative test #1",
			args: args{
				t:     "unknownType2",
				name:  "test2",
				value: 123123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest(t, func(mux *chi.Mux, repo repositories.Repository, db *db.DB) {
				ctx := context.Background()
				_ = flushDB(ctx, db)

				data, _ := json.Marshal(metric.JSONMetric{
					ID:    tt.args.name,
					MType: tt.args.t,
					Value: &tt.args.value,
				})

				req, err := http.NewRequest("POST", "/update", bytes.NewBuffer(data))
				if err != nil {
					t.Fatal(err)
				}
				rr := httptest.NewRecorder()
				mux.ServeHTTP(rr, req)
				assert.Equal(t, rr.Code, http.StatusBadRequest)
			})
		})
	}
}

func Test_updateNegativeCounterHandler(t *testing.T) {
	type args struct {
		t     metric.Type
		name  string
		value float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "negative test #1",
			args: args{
				t:     metric.Counter,
				name:  "test",
				value: 100,
			},
		},
		{
			name: "negative test #2",
			args: args{
				t:     metric.Counter,
				name:  "test2",
				value: 1232,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest(t, func(mux *chi.Mux, repo repositories.Repository, db *db.DB) {
				ctx := context.Background()
				_ = flushDB(ctx, db)

				data, _ := json.Marshal(metric.JSONMetric{
					ID:    tt.args.name,
					MType: tt.args.t.String(),
					Value: &tt.args.value,
				})

				req, err := http.NewRequest("POST", "/update", bytes.NewBuffer(data))
				if err != nil {
					t.Fatal(err)
				}
				rr := httptest.NewRecorder()
				mux.ServeHTTP(rr, req)
				assert.Equal(t, rr.Code, http.StatusBadRequest)
			})
		})
	}
}
