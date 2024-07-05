package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	l "service-hf-client-p5/external/logger"
	ps "service-hf-client-p5/external/strings"
	"service-hf-client-p5/internal/core/application"
	"service-hf-client-p5/internal/core/domain/entity/dto"
	"strings"
)

type Handler interface {
	HandlerClient(rw http.ResponseWriter, req *http.Request)
	HealthCheck(rw http.ResponseWriter, req *http.Request)
}

type handler struct {
	app application.Application
}

func NewHandler(app application.Application) Handler {
	return handler{app: app}
}

func (h handler) HandlerClient(rw http.ResponseWriter, req *http.Request) {

	var routes = map[string]http.HandlerFunc{
		"get hermes_foods/client/{cpf}": h.getByCPF,
		"post hermes_foods/client":      h.saveClient,
	}

	handler, err := router(req.Method, req.URL.Path, routes)

	if err == nil {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.Method + " " + req.URL.Path + ` not found"} `))
}

func (h handler) HealthCheck(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status": "OK"}`))
}

func (h handler) getByCPF(rw http.ResponseWriter, req *http.Request) {
	cpf := getCpf(req.URL.Path)

	c, err := h.app.GetClientByCPF(cpf, "")

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get client by ID: %v"} `, err)
		return
	}

	if c == nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"error": "client not found"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(ps.MarshalString(c)))
}

func (h *handler) saveClient(rw http.ResponseWriter, req *http.Request) {
	msgID := l.MessageID(req.Header.Get(l.MessageIDKey))

	if req.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"} `))
		return
	}

	var buff bytes.Buffer

	var reqClient dto.RequestClient

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqClient); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	c, err := h.app.SaveClient(msgID, reqClient)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save client: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(ps.MarshalString(c)))
}

func getCpf(url string) string {
	indexCpf := strings.Index(url, "client/")

	if indexCpf == -1 {
		return ""
	}

	return strings.ReplaceAll(url[indexCpf:], "client/", "")
}
