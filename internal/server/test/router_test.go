package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"screamer/internal/server/di"
	"screamer/internal/server/router"
	"testing"
)

func Test_updateHandler(t *testing.T) {
	type args struct {
		url  string
		code int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "positive test #1",
			args: args{
				url:  "/update/counter/testCounter/100",
				code: http.StatusOK,
			},
		},
		{
			name: "positive test #2",
			args: args{
				url:  "/update/gauge/testCounter/100.1",
				code: http.StatusOK,
			},
		},
		{
			name: "positive test #3",
			args: args{
				url:  "/update/counter/testCounter/100.1",
				code: http.StatusBadRequest,
			},
		},
		{
			name: "positive test #4",
			args: args{
				url:  "/update/unknown/testCounter/100",
				code: http.StatusBadRequest,
			},
		},
		{
			name: "positive test #5",
			args: args{
				url:  "/unknown/unknown/testCounter/100",
				code: http.StatusNotFound,
			},
		},
	}
	container := di.BuildContainer()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = container.Invoke(func(router *router.Router) {
				req, err := http.NewRequest("POST", tt.args.url, nil)
				if err != nil {
					t.Fatal(err)
				}
				rr := httptest.NewRecorder()
				handler := router.GetRouter()
				handler.ServeHTTP(rr, req)
				assert.Equal(t, rr.Code, tt.args.code)
			})
		})
	}
}
