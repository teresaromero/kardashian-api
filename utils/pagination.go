package utils

import (
	"context"
	"fmt"
	"kardashian_api/models"
	"net/http"
	"reflect"
	"strings"
)

func pageResponse(r *http.Request, data interface{}) *models.Page {

	d := reflect.ValueOf(data)
	size := d.Len()

	p := GetPaginationContext(r)
	u := getRequestURI(r)
	self := fmt.Sprintf("%s?page=%d", u, p.Page)

	var links models.Links
	nx := p.Page + 1
	pv := p.Page - 1
	if pv == 0 {
		next := fmt.Sprintf("%s?page=%d", u, nx)

		links = models.Links{
			Self: self,
			Next: next,
		}
	} else if size < p.Limit {
		prev := fmt.Sprintf("%s?page=%d", u, pv)

		links = models.Links{
			Self: self,
			Prev: prev,
		}
	} else {
		next := fmt.Sprintf("%s?page=%d", u, nx)
		prev := fmt.Sprintf("%s?page=%d", u, pv)
		links = models.Links{
			Self: self,
			Next: next,
			Prev: prev,
		}
	}

	page := &models.Page{
		Results: data,
		Size:    size,
		Start:   p.Skip,
		Limit:   p.Limit,
		Links:   links,
		Total:   p.Total,
	}
	return page

}

func AddToRequest(r *http.Request, key interface{}, value interface{}) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), key, value))
}

func getValueFromRequest(r *http.Request, key string) interface{} {
	return r.Context().Value(key)
}

func GetPaginationContext(r *http.Request) *models.PaginationOpts {
	return getValueFromRequest(r, "pagination").(*models.PaginationOpts)
}

func getRequestURI(r *http.Request) string {
	u := strings.Split(r.RequestURI, "?")
	return strings.TrimSuffix(u[0], "/")
}
