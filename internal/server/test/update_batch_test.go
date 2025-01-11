package test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"screamer/internal/common/application/dto"
	"screamer/internal/common/domain"
	"screamer/internal/server/application/repo"
	"screamer/internal/server/infrastructure/db"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func Test_updateBatchGaugeHandler(t *testing.T) {
	type args struct {
		t     domain.Type
		name  string
		value float64
	}
	tests := []struct {
		name string
		args []args
	}{
		{
			name: "positive test #1",
			args: []args{
				{
					t:     domain.Gauge,
					name:  "test1",
					value: 100,
				},
				{
					t:     domain.Gauge,
					name:  "test1",
					value: 400,
				},
			},
		},
		{
			name: "positive test #2",
			args: []args{
				{
					t:     domain.Gauge,
					name:  "test1",
					value: 1232,
				},
				{
					t:     domain.Gauge,
					name:  "test1",
					value: 23123,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest(t, func(mux *chi.Mux, repo repo.Repository, db *db.DB) {
				ctx := context.Background()
				_ = flushDB(ctx, db)

				data, _ := json.Marshal([]dto.JsonMetric{
					{
						ID:    tt.args[0].name,
						MType: tt.args[0].t.String(),
						Value: &tt.args[0].value,
					},
					{
						ID:    tt.args[1].name,
						MType: tt.args[1].t.String(),
						Value: &tt.args[1].value,
					},
				})

				req, err := http.NewRequest("POST", "/updates", bytes.NewBuffer(data))
				if err != nil {
					t.Fatal(err)
				}
				rr := httptest.NewRecorder()
				mux.ServeHTTP(rr, req)
				assert.Equal(t, rr.Code, http.StatusOK)
				m, err := repo.Get(ctx, domain.Ident{
					Type: tt.args[0].t,
					Name: tt.args[0].name,
				})
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, m.Value, float64(tt.args[1].value))
			})
		})
	}
}

func Test_updateBatchCounterHandler(t *testing.T) {
	type args struct {
		t     domain.Type
		name  string
		delta int64
	}
	tests := []struct {
		name string
		args []args
	}{
		{
			name: "positive test #2",
			args: []args{
				{
					t:     domain.Counter,
					name:  "test1",
					delta: 42343,
				},
				{
					t:     domain.Counter,
					name:  "test1",
					delta: 22,
				},
			},
		},
		{
			name: "positive test #2",
			args: []args{
				{
					t:     domain.Counter,
					name:  "test1",
					delta: 4234,
				},
				{
					t:     domain.Counter,
					name:  "test1",
					delta: 3243,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest(t, func(mux *chi.Mux, repo repo.Repository, db *db.DB) {
				ctx := context.Background()
				_ = flushDB(ctx, db)

				data, _ := json.Marshal([]dto.JsonMetric{
					{
						ID:    tt.args[0].name,
						MType: tt.args[0].t.String(),
						Delta: &tt.args[0].delta,
					},
					{
						ID:    tt.args[1].name,
						MType: tt.args[1].t.String(),
						Delta: &tt.args[1].delta,
					},
				})

				req, err := http.NewRequest("POST", "/updateBatch", bytes.NewBuffer(data))
				if err != nil {
					t.Fatal(err)
				}
				rr := httptest.NewRecorder()
				mux.ServeHTTP(rr, req)
				assert.Equal(t, rr.Code, http.StatusOK)
				m, err := repo.Get(ctx, domain.Ident{
					Type: tt.args[0].t,
					Name: tt.args[0].name,
				})
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, int64(m.Value), tt.args[0].delta+tt.args[1].delta)
			})
		})
	}
}

func Test_updateBatchNegativeTypeHandler(t *testing.T) {
	type args struct {
		t     domain.Type
		name  string
		delta int64
	}
	tests := []struct {
		name string
		args []args
	}{
		{
			name: "positive test #2",
			args: []args{
				{
					t:     "unknownType",
					name:  "test1",
					delta: 42343,
				},
				{
					t:     domain.Counter,
					name:  "test1",
					delta: 22,
				},
			},
		},
		{
			name: "positive test #2",
			args: []args{
				{
					t:     "unknownType",
					name:  "test1",
					delta: 4234,
				},
				{
					t:     domain.Counter,
					name:  "test1",
					delta: 3243,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest(t, func(mux *chi.Mux, repo repo.Repository, db *db.DB) {
				ctx := context.Background()
				_ = flushDB(ctx, db)

				data, _ := json.Marshal([]dto.JsonMetric{
					{
						ID:    tt.args[0].name,
						MType: tt.args[0].t.String(),
						Delta: &tt.args[0].delta,
					},
					{
						ID:    tt.args[1].name,
						MType: tt.args[1].t.String(),
						Delta: &tt.args[1].delta,
					},
				})

				req, err := http.NewRequest("POST", "/updateBatch", bytes.NewBuffer(data))
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
