package test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"screamer/internal/common"
	"screamer/internal/common/domain"
	"screamer/internal/server/application/repo"
	"screamer/internal/server/infrastructure/db"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func Test_updateOldHandler(t *testing.T) {
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
				t:     domain.Counter,
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
		{
			name: "positive test #3",
			args: args{
				t:     domain.Gauge,
				name:  "test3",
				value: 0,
			},
		},
		{
			name: "positive test #4",
			args: args{
				t:     domain.Gauge,
				name:  "test4",
				value: 123122,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest(t, func(mux *chi.Mux, repo repo.Repository, db *db.DB) {
				ctx := context.Background()
				_ = flushDB(ctx, db)
				url := fmt.Sprintf("/update/%v/%v/%v", tt.args.t.String(), tt.args.name, tt.args.value)
				req, err := http.NewRequest("POST", url, nil)
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

func Test_updateOldNegativeHandler(t *testing.T) {
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
				name:  "test1",
				value: 100,
			},
		},
		{
			name: "negative test #1",
			args: args{
				t:     domain.Gauge,
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
				_, err := repo.Get(ctx, domain.Ident{
					Type: tt.args.t,
					Name: tt.args.name,
				})
				assert.ErrorIs(t, err, common.ErrMetricNotExists)
			})
		})
	}
}

func Test_updateOldNegativeTypeHandler(t *testing.T) {
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
				url := fmt.Sprintf("/update/%v/%v/%v", tt.args.t, tt.args.name, tt.args.value)
				req, err := http.NewRequest("POST", url, nil)
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
