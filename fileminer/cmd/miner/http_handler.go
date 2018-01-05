package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/sKudryashov/asec/fileminer/internal/model"
)

func handlerHTTP(data os.FileInfo) {
	url := "fileserver"
	finfo := newHandler(&model.FileInfo{
		Name:    data.Name(),
		Mode:    data.Mode().String(),
		ModTime: data.ModTime(),
		Size:    data.Size(),
	})

	var body io.Reader
	var contentType string

	coin := tossTheCoin()
	if coin == codeRandJSON {
		body = finfo.createJSONBody()
		contentType = "application/json"

	} else if coin == codeRandXML {
		body = finfo.createXMLBody()
		contentType = "text/xml"
	}

	resp, err := http.Post(fmt.Sprintf("%s://%s:%d/", *schema, url, *port), contentType, body)
	if err != nil {
		log.Printf("error sending request %v", err)
	}
	bodyRspBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading response %v", err)
	}
	// debug msg must be deleted
	println("raw response from server", string(bodyRspBytes))
	resp.Body.Close()
	if err != nil {
		log.Printf("error creating request %v", err)
	}
}

func tossTheCoin() int {
	return rand.Intn(1)
}

type bodyHandler struct {
	finfo *model.FileInfo
}

func newHandler(finfo *model.FileInfo) *bodyHandler {
	return &bodyHandler{
		finfo: finfo,
	}
}

func (body bodyHandler) createJSONBody() io.Reader {
	byteBody, err := json.Marshal(body.finfo)
	if err != nil {
		log.Printf("error marshalling json request %v", err)
	}
	bufS := bytes.NewBuffer(byteBody)
	return bufS
}

func (body bodyHandler) createXMLBody() io.Reader {
	byteBody, err := xml.Marshal(body.finfo)
	if err != nil {
		log.Printf("error marshalling xml request %v", err)
	}
	buf := bytes.NewBuffer(byteBody)

	return buf
}
