package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"port_domain_service_backend/internal/adapters/repository"
	"port_domain_service_backend/internal/core/domain"
	"port_domain_service_backend/internal/core/services"
	"syscall"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHomePage(t *testing.T) {
	Convey("TestCreatePortDomain Should create port domains", t, func() {
		t.Logf("Starting TestCreatePortDomain")

		sigdone := make(chan os.Signal, 1)
		signal.Notify(sigdone, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		portDomains := make(map[string]map[string]interface{})
		repo := repository.PortDomainsRepository{
			PortDomains: portDomains,
		}
		pds := services.NewPortDomainsService(repo)

		svr, err := NewServer(ctx, cancel, pds)
		if err != nil {
			logrus.WithError(err).Fatal("unable to create the API server")
		}

		go func() {
			svr.Run()
			cancel()
		}()

		time.Sleep(2 * time.Second)

		Convey("TestHomePage home  ", func() {
			setReq := "/home"
			req := httptest.NewRequest(http.MethodGet, setReq, nil)
			rec := httptest.NewRecorder()
			svr.httpHandler.HomePage(rec, req)
			resp := rec.Result()
			//So(rec.Code, ShouldEqual, http.StatusOK)
			t.Log(resp.StatusCode)
		})

		cancel()

		select {
		case <-ctx.Done():
			logrus.WithError(ctx.Err()).Info("main got cancel")
		case <-sigdone:
			logrus.Info("got sigdone, sending cancel")
			cancel()
		}

		t.Logf("exiting TestCreatePortDomain")
	})
}

func TestCreateUpdatePortDomain(t *testing.T) {
	Convey("TestCreatePortDomain Should create port domains", t, func() {
		t.Logf("Starting TestCreatePortDomain")

		sigdone := make(chan os.Signal, 1)
		signal.Notify(sigdone, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		portDomains := make(map[string]map[string]interface{})
		repo := repository.PortDomainsRepository{
			PortDomains: portDomains,
		}
		pds := services.NewPortDomainsService(repo)

		svr, err := NewServer(ctx, cancel, pds)
		if err != nil {
			logrus.WithError(err).Fatal("unable to create the API server")
		}

		go func() {
			svr.Run()
			cancel()
		}()

		time.Sleep(2 * time.Second)

		Convey("TestCreateUpdatePortDomain create ", func() {
			setReq := "/create"

			// Create a buffer to store the request body
			var buf bytes.Buffer

			// Create a new multipart writer with the buffer
			w := multipart.NewWriter(&buf)
			// Add a file to the request
			file, err := os.Open("/home/rajashrijadhav/RajashriJadhavData/RAJASHRI_ASSIGNMENTS/port_domain_service_backend/testdata/ports.json")
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()
			// Create a new form field
			fw, err := w.CreateFormFile("jsonFile", "/home/rajashrijadhav/RajashriJadhavData/RAJASHRI_ASSIGNMENTS/port_domain_service_backend/testdata/ports.json")
			if err != nil {
				t.Fatal(err)
			}
			// Copy the contents of the file to the form field
			if _, err := io.Copy(fw, file); err != nil {
				t.Fatal(err)
			}
			// Close the multipart writer to finalize the request
			w.Close()

			//f, _ := os.Open("/home/rajashrijadhav/RajashriJadhavData/RAJASHRI_ASSIGNMENTS/port_domain_service_backend/testdata/ports.json")
			req := httptest.NewRequest(http.MethodPost, setReq, &buf)
			req.Header.Set("Content-Type", w.FormDataContentType())

			rec := httptest.NewRecorder()
			svr.httpHandler.CreatePortDomain(rec, req)
			resp := rec.Result()
			So(rec.Code, ShouldEqual, http.StatusCreated)
			t.Log(resp.StatusCode)

			Convey("TestCreateUpdatePortDomain update ", func() {
				setReq := "/update"

				//payload := []byte(`{"USSEA":{"name":"Cape Romanzof","city":"Cape Romanzof","province":"Alaska","country":"United Arab Emirates","alias":["Tacoma"],"regions":[],"coordinates":[55.2756505, 25.284755],"timezone":"America/Anchorage","unlocs":["AEPRA"],"code":"3001"}}`)

				//  -- this way to supply data also works
				jd := make(map[string]map[string]interface{})
				jd["USSEA"] = map[string]interface{}{
					"name":        "Cape Romanzof",
					"city":        "Cape Romanzof",
					"province":    "Alaska",
					"country":     "United Arab Emirates",
					"alias":       []string{"Tacoma"},
					"regions":     []string{},
					"coordinates": []float32{55.2756505, 25.284755},
					"timezone":    "America/Anchorage",
					"unlocs":      []string{"AEPRA"},
					"code":        "3001",
				}
				payload, _ := json.Marshal(jd)

				logrus.Info(payload)

				req := httptest.NewRequest(http.MethodPatch, setReq, bytes.NewBuffer(payload))
				req.Header.Set("Content-Type", "application/json")
				rec := httptest.NewRecorder()
				svr.httpHandler.UpdatePortDomain(rec, req)
				resp := rec.Result()
				defer resp.Body.Close()
				So(rec.Code, ShouldEqual, http.StatusOK)
				var pd domain.PortDetail
				err := json.NewDecoder(resp.Body).Decode(&pd)
				if err != nil {
					t.Log(err.Error())
				}
				t.Log("update response : ")
				t.Log(pd)
				t.Log(resp.StatusCode)
			})
		})

		cancel()

		select {
		case <-ctx.Done():
			logrus.WithError(ctx.Err()).Info("main got cancel")
		case <-sigdone:
			logrus.Info("got sigdone, sending cancel")
			cancel()
		}

		t.Logf("exiting TestCreateUpdatePortDomain")
	})
}

func TestCreateUpdatePortDomainIncorrectMethod(t *testing.T) {
	Convey("TestCreatePortDomain Should create port domains", t, func() {
		t.Logf("Starting TestCreatePortDomain")

		sigdone := make(chan os.Signal, 1)
		signal.Notify(sigdone, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		portDomains := make(map[string]map[string]interface{})
		repo := repository.PortDomainsRepository{
			PortDomains: portDomains,
		}
		pds := services.NewPortDomainsService(repo)

		svr, err := NewServer(ctx, cancel, pds)
		if err != nil {
			logrus.WithError(err).Fatal("unable to create the API server")
		}

		go func() {
			svr.Run()
			cancel()
		}()

		time.Sleep(2 * time.Second)

		Convey("TestCreateUpdatePortDomainIncorrectMethod create ", func() {
			setReq := "/create"
			req := httptest.NewRequest(http.MethodGet, setReq, nil)
			rec := httptest.NewRecorder()
			svr.httpHandler.CreatePortDomain(rec, req)
			resp := rec.Result()
			So(rec.Code, ShouldEqual, http.StatusMethodNotAllowed)
			t.Log(resp.StatusCode)

			Convey("TestCreateUpdatePortDomainIncorrectMethod update ", func() {
				setReq := "/update"
				req := httptest.NewRequest(http.MethodGet, setReq, nil)
				rec := httptest.NewRecorder()
				svr.httpHandler.UpdatePortDomain(rec, req)
				resp := rec.Result()
				defer resp.Body.Close()
				So(rec.Code, ShouldEqual, http.StatusMethodNotAllowed)
				t.Log(resp.StatusCode)
			})
		})

		cancel()

		select {
		case <-ctx.Done():
			logrus.WithError(ctx.Err()).Info("main got cancel")
		case <-sigdone:
			logrus.Info("got sigdone, sending cancel")
			cancel()
		}

		t.Logf("exiting TestCreateUpdatePortDomain")
	})
}
