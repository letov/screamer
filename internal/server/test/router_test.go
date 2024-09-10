package test

import (
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"net/http"
	"net/http/httptest"
	"screamer/internal/server/di"
	"testing"
)

func inttest(t *testing.T, r interface{}) {
	app := fxtest.New(t, di.InjectApp(), fx.Invoke(r))
	defer app.RequireStop()
	app.RequireStart()
}

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inttest(t, func(mux *chi.Mux) {
				req, err := http.NewRequest("POST", tt.args.url, nil)
				if err != nil {
					t.Fatal(err)
				}
				rr := httptest.NewRecorder()
				mux.ServeHTTP(rr, req)
				assert.Equal(t, rr.Code, tt.args.code)
			})
		})
	}
}
