package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"port_domain_service_backend/internal/core/domain"
	"port_domain_service_backend/internal/core/services"

	"github.com/sirupsen/logrus"
)

type HTTPHandler struct {
	svc *services.PortDomainsService
}

func NewHTTPHandler(PDService *services.PortDomainsService) *HTTPHandler {
	return &HTTPHandler{
		svc: PDService,
	}
}

func (h *HTTPHandler) HomePage(w http.ResponseWriter, _ *http.Request) {
	logrus.Info("Hello Port Domain Service Backend!")
}

func (h *HTTPHandler) CreatePortDomain(w http.ResponseWriter, r *http.Request) {
	// call services create port domain
	logrus.Info("HTTPHandler CreatePortDomain called!")
	if r.Method == "POST" {
		file, _, err := r.FormFile("jsonFile")
		if err != nil {
			logrus.Error(err.Error())
			defer file.Close()
		}
		defer file.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			logrus.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}

		err = h.svc.CreatePortDomain(buf.Bytes())
		if err != nil {
			logrus.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusCreated)
	} else {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *HTTPHandler) UpdatePortDomain(w http.ResponseWriter, r *http.Request) {
	logrus.Info("UpdatePortDomain called!")
	var ports domain.PortDetails

	if r.Method == "PATCH" {
		if err := json.NewDecoder(r.Body).Decode(&ports); err != nil {
			logrus.Error("unable to parse request body")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logrus.Info(ports)

		details, err := h.svc.UpdatePortDomain(ports)
		if err != nil {
			logrus.Error("unable to update port domain")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if details == nil && err == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(details)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
