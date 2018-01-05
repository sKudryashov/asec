package servd

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

// var getStorage func() *storage.Storage

func TestAddHandler(t *testing.T) {
	// decided not to break this down to smaller unit tests for each layer, but in general that would be better in
	// a real project
	reqData := []byte(`{"name":"TestFile", "mode":"0644", "size":291223, "mod_time":"2006-01-02T15:04:05Z07:00"}`)
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("POST", "/", bytes.NewBuffer(reqData))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	e := echo.New()
	ctx := e.NewContext(request, recorder)
	h := newHandler()
	if err := h.AddFileData(ctx); err != nil {
		t.Fatal(err)
	}
	rsp := recorder.Result()
	if hStatus := rsp.StatusCode; hStatus != 202 {
		t.Fatalf("status 202 expected, got %d", hStatus)
	}
	rspBody, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	println("rspBody", string(rspBody))
}
