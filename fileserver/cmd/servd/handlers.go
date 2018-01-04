package servd

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sKudryashov/asec/fileserver/internal/finfo"
)

type handler struct {
	service *finfo.Service
}

func newHandler() *handler {
	h := new(handler)
	h.service = finfo.NewFinfoService()
	return h
}

var status map[string]string

//AddFileData handler calls service which adds data to the storage
func (h *handler) AddFileData(c echo.Context) error {
	r := c.Request()
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}
	err = h.service.SaveFileInfo(data)
	if err != nil {
		return err
	}
	status["status"] = "ok"

	return c.JSON(http.StatusAccepted, status)
}

//GetFileData handler returns data which has been previously stored in DB
func (h *handler) GetFileData(c echo.Context) error {
	data, err := h.service.GetFileInfo()
	if err != nil {
		return err
	}

	return c.JSONBlob(http.StatusOK, data)
}
