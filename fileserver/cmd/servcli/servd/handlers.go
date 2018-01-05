package servd

import (
	"fmt"
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

//AddFileData handler calls service which adds data to the storage
func (h *handler) AddFileData(c echo.Context) error {
	var err error
	r := c.Request()
	defer r.Body.Close()
	ctype := r.Header.Get("Content-Type")
	if ctype == "application/json" {
		err = h.service.SaveFileInfo(r.Body, finfo.TypeJSON)
	} else if ctype == "text/xml" {
		err = h.service.SaveFileInfo(r.Body, finfo.TypeXML)
	} else {
		return fmt.Errorf("header Content-Type application/json or text/xml expected, %s given", ctype)
	}
	status := make(map[string]string)
	if err != nil {
		status["status"] = "error"
		status["message"] = err.Error()
		return c.JSON(http.StatusInternalServerError, status)
	}
	status["status"] = "ok"

	return c.JSON(http.StatusAccepted, status)
}

func (h *handler) GetStatData(c echo.Context) error {
	stat := h.service.GetStat()
	rsp := map[string]int{
		"rps": stat.GetRatePerSecond(),
	}

	return c.JSON(http.StatusOK, rsp)
}

//GetFileData handler returns data which has been previously stored in DB
func (h *handler) GetFileData(c echo.Context) error {
	data, err := h.service.GetFileInfo()
	if err != nil {
		return err
	}

	return c.JSONBlob(http.StatusOK, data)
}
