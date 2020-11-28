package service

import (
	"fmt"
	"net/http"
	"time"
)

// handlerAppVersion used for application version handler
var handlerAppVersion = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "version api: %v\n", time.Now())
}
