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

func Test_updateCounterHandler(t *testing.T) {
	type args struct {
		t     domain.Type
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
				t:     domain.Counter,
				name:  "test",
				value: 100,
			},
		},
		{
			name: "positive test #2",
			args: args{
				t:     domain.Counter,
				name:  "test2",
				value: 1232,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest(t, func(mux *chi.Mux, repo repo.Repository, db *db.DB) {
				ctx := context.Background()
				_ = flushDB(ctx, db)

				data, _ := json.Marshal(dto.JsonMetric{
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
				m, err := repo.Get(ctx, domain.Ident{
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
		t     domain.Type
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
				t:     domain.Gauge,
				name:  "test",
				value: 100,
			},
		},
		{
			name: "positive test #2",
			args: args{
				t:     domain.Gauge,
				name:  "test2",
				value: 1232.22,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest(t, func(mux *chi.Mux, repo repo.Repository, db *db.DB) {
				ctx := context.Background()
				_ = flushDB(ctx, db)

				data, _ := json.Marshal(dto.JsonMetric{
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
				m, err := repo.Get(ctx, domain.Ident{
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
			initTest(t, func(mux *chi.Mux, repo repo.Repository, db *db.DB) {
				ctx := context.Background()
				_ = flushDB(ctx, db)

				data, _ := json.Marshal(dto.JsonMetric{
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
		t     domain.Type
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
				t:     domain.Counter,
				name:  "test",
				value: 100,
			},
		},
		{
			name: "negative test #2",
			args: args{
				t:     domain.Counter,
				name:  "test2",
				value: 1232,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest(t, func(mux *chi.Mux, repo repo.Repository, db *db.DB) {
				ctx := context.Background()
				_ = flushDB(ctx, db)

				data, _ := json.Marshal(dto.JsonMetric{
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
