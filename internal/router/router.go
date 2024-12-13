package router

import (
	"net/http"
	"strings"

	"github.com/BorzooMV/how-is-weather/internal/handlers"
)

type Router struct {
}

func (ro Router) WeatherRouter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "wrong method provided", http.StatusBadRequest)
		return
	}
	if hasValidSuffix, suffix := hasAcceptableSuffix(r.URL.Path, "/api/weather/"); hasValidSuffix {
		handlers.GetWeather(w, r, suffix)
	} else {
		http.Error(w, "wrong method provided", http.StatusBadRequest)
		return
	}
}

func hasSuffix(s string, p string) bool {
	return s != strings.TrimPrefix(s, p)
}

func hasAcceptableSuffix(s string, p string) (hasValidSuffix bool, suffix string) {
	if hasSuffix(s, p) {
		suffix := strings.Split(strings.TrimPrefix(s, p), "/")[0]
		hasValidSuffix = suffix != ""
		return hasValidSuffix, suffix
	}

	return hasValidSuffix, suffix
}
