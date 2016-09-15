package system

import (
	"net/http"

	"github.com/gorilla/context"
)

func (application *Application) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "application", application)
	next(w, r)
}
