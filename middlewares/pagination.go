package middlewares

import (
	"context"
	"errors"
	"kardashian_api/custom_errors"
	"kardashian_api/database"
	"kardashian_api/models"
	"kardashian_api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func collectionTotalCount(c string, filter bson.D) int64 {
	total, _ := database.Use(c).CountDocuments(context.TODO(), filter)
	return total
}

func validatePagination(skip int, total int64) error {
	if int64(skip) > total {
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
		page, skip, limit := paginationParams(c.Query("page"))
		total := collectionTotalCount(collection, bson.D{})
		err := validatePagination(skip, total)
		if err != nil {
			utils.HandleHttpError(c, custom_errors.BadRequest(err))
		} else {
			paginationOpts := &models.PaginationOpts{
				Page:  page,
				Limit: limit,
				Skip:  skip,
				Total: int(total),
			}
			c.Request = utils.AddToRequest(c.Request, "pagination", paginationOpts)

			c.Next()
		}

	}
}
