package internal

import (
	"net/http"
)

// App is the worst.
type App struct {
	Message string
}

func (a App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(a.Message))
}
