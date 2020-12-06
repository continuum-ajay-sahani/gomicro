package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Error contains the message about error
type Error struct {
	Message string `json:"message"`
}

// ErrorMessage contains the error
type ErrorMessage struct {
	Error Error `json:"error"`
}

// MethodHandler stores mapping for each available method-handler combinations
type MethodHandler map[string]http.Handler

// ServeHTTP checks if requested method is allowed and returns 405 if not.
// This middleware was copied from gorilla/handlers and changed according to Continuums REST API requirements
func (h MethodHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.applyHeaders(w.Header())
	if handler, ok := h[req.Method]; ok {
		handler.ServeHTTP(w, req)
		return
	}
	if req.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	SendMethodNotAllowed(w, req, "", nil)
}

func applyDebugHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}

func (h MethodHandler) applyHeaders(header http.Header) {
	header.Set("Access-Control-Allow-Origin", "*")
}

// GetReqRouter return router to handle object
var GetReqRouter = func() *mux.Router {
	return mux.NewRouter()
}

// GetMiddlewareManager return middleware for web
var GetMiddlewareManager = func() *negroni.Negroni {
	middlewareManager := negroni.New()
	middlewareManager.Use(negroni.NewRecovery())
	return middlewareManager
}

// GetHTTPServer return http server object
var GetHTTPServer = func(router *mux.Router, addr string) *http.Server {
	middlewareManager := GetMiddlewareManager()
	middlewareManager.UseHandler(router)
	// Good practice to set timeouts to avoid Slowloris attacks.
	return &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      middlewareManager,
	}
}

// WrapMethodHandler wrapper for MethodHandler map
func WrapMethodHandler(method string, fn func(w http.ResponseWriter, _ *http.Request)) MethodHandler {
	return MethodHandler{
		method: http.HandlerFunc(fn),
	}
}

// SendInternalServerError sends Internal Server Error Status and logs an error if it exists
func SendInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	SendError(w, r, http.StatusInternalServerError, "", err)
}

// SendMethodNotAllowed sends Internal Server Error Status and logs an error if it exists
func SendMethodNotAllowed(w http.ResponseWriter, r *http.Request, message string, err error) {
	SendError(w, r, http.StatusMethodNotAllowed, message, err)
}

// SendError writes a defined string as an error message
// with appropriate headers to the HTTP response
func SendError(w http.ResponseWriter, r *http.Request, code int, message string, errForLog error) {
	var data []byte
	var err error
	if message != "" {
		data, err = json.Marshal(ErrorMessage{Error{message}})
		if err != nil {
			SendInternalServerError(w, r, err)
			return
		}
	}
	if errForLog != nil {
		ctx := r.Context()
		txnValue := GetTransactionContextValue(ctx)
		errMsg := fmt.Sprintf("%v %v %v %v", txnValue, errForLog.Error(), r.Method, r.URL)
		logger.Error(errMsg)
	}
	render(w, code, data)
}

func render(w http.ResponseWriter, code int, response []byte) {
	if response == nil {
		w.WriteHeader(code)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		errMsg := fmt.Sprintf("Cannot write response correlation id = %v", gocql.TimeUUID().String())
		logger.Error(errMsg)
	}
}

// RenderJSON is used for rendering JSON response body with appropriate headers
func RenderJSON(w http.ResponseWriter, r *http.Request, response interface{}) {
	renderJSON(w, r, http.StatusOK, response)
}

// RenderJSONCreated is used for rendering JSON response body when new resource has been created
func RenderJSONCreated(w http.ResponseWriter, r *http.Request, response interface{}) {
	renderJSON(w, r, http.StatusCreated, response)
}

// renderJSON is used for rendering JSON response
func renderJSON(w http.ResponseWriter, r *http.Request, status int, response interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		SendInternalServerError(w, r, fmt.Errorf("common.renderJSON: err:%s", err))
		return
	}
	render(w, status, data)
}

// ParseRequestBody parse request query parameter from request body(POST, PUT, DELETE)
func ParseRequestBody(r *http.Request, m interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	bv := make(map[string]string)
	err = json.Unmarshal(body, &bv)
	if err != nil {
		panic(err)
	}
	q, _ := bv["query"]
	json.Unmarshal([]byte(q), &m)
}
