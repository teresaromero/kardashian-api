package middlewares

import (
	"errors"
	"kardashian_api/database"
	"kardashian_api/models"
	"kardashian_api/utils"
	"kardashian_api/utils/http_errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func validatePagination(skip int, total int) error {
	if skip > total {
		return errors.New("page out of range")
	}
	return nil
}

func paginationParams(pageQuery string) (page int, skip int, limit int) {
	p, _ := strconv.Atoi(pageQuery)
	limit = 10
	if p == 0 {
		page = p + 1
	} else {
		page = p
	}
	skip = (page - 1) * limit
	return page, skip, limit
}

func Pagination(collection string) gin.HandlerFunc {
	return func(c *gin.Context) {
		total := database.CountCollectionDocs(collection, bson.M{})
		if total == 0 {
			utils.HandleHttpError(c, http_errors.BadRequest(errors.New("no documents for this collection")))
		}
		page, skip, limit := paginationParams(c.Query("page"))
		err := validatePagination(skip, total)
		if err != nil {
			utils.HandleHttpError(c, http_errors.BadRequest(err))
		} else {
			paginationOpts := &models.PaginationOpts{
				Page:  page,
				Limit: limit,
				Skip:  skip,
				Total: total,
			}
			c.Request = utils.AddToRequest(c.Request, "pagination", paginationOpts)

			c.Next()
		}

	}
}
