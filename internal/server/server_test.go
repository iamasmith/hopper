package server

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStartFail(t *testing.T) {
	assert := assert.New(t)
	const binding = ":8000"
	a := New(context.Background(), binding)
	go a.Start()
	defer a.Stop()
	time.Sleep(time.Duration(1) * time.Second)
	b := New(context.Background(), binding)
	assert.Error(b.Start(), "should error")
	b.Stop()
}

func TestServe(t *testing.T) {
	assert := assert.New(t)
	server := New(context.Background(), "127.0.0.1:8000")
	m := server.Mux()
	m.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		},
	)
	go server.Start()
	client := http.Client{}
	resp, err := client.Get("http://127.0.0.1:8000")
	assert.NoError(err, "should not error")
	body, err := io.ReadAll(resp.Body)
	assert.NoError(err, "should not error")
	assert.Equal([]byte("OK"), body, "should match")
	server.Stop()
}
