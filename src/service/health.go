package service

import (
	"fmt"
	"net/http"
	"time"
)

// handlerAppVersion used for application version handler
var handlerAppHealth = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "health api: %v\n", time.Now())
}
