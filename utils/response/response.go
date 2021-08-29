package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kardashian_api/models"
	"kardashian_api/utils/http_errors"
	"kardashian_api/utils/request"
	"log"
	"net/http"
	"reflect"
)

func mapPage(r *http.Request, data interface{}) *models.Page {
	d := reflect.ValueOf(data)
	size := d.Len()

	p := request.GetContextValue(r, "pagination").(*models.PaginationOpts)
	u := request.GetURI(r)
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

func HttpError(c *gin.Context, err *http_errors.HttpError) {
	log.Printf("Error HandleHttpError: %v", err.Err)
	c.AbortWithStatusJSON(err.Status(), gin.H{"status": err.Status(), "message": err.Message})
}

func SingleResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func PageResponse(c *gin.Context, data interface{}) {
	page := mapPage(c.Request, data)
	c.JSON(http.StatusOK, page)

}
