package request

import (
	"context"
	"net/http"
	"strings"
)

func AddToContext(r *http.Request, key interface{}, value interface{}) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), key, value))
}

func GetContextValue(r *http.Request, name string) interface{} {
	return r.Context().Value(name)
}

func GetURI(r *http.Request) string {
	u := strings.Split(r.RequestURI, "?")
	return strings.TrimSuffix(u[0], "/")
}
