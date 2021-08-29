package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"kardashian_api/database"
	"kardashian_api/models"
	"kardashian_api/utils/http_errors"
	"kardashian_api/utils/pagination"
	"kardashian_api/utils/request"
	"kardashian_api/utils/response"
)

func Pagination(collection string) gin.HandlerFunc {
	return func(c *gin.Context) {
		total := database.CountCollectionDocs(collection, bson.M{})
		if total == 0 {
			response.HttpError(c, http_errors.BadRequest(errors.New("no documents for this collection")))
		}
		page, skip, limit := pagination.Params(c.Query("page"))
		err := pagination.Validate(skip, total)
		if err != nil {
			response.HttpError(c, http_errors.BadRequest(err))
		} else {
			paginationOpts := &models.PaginationOpts{
				Page:  page,
				Limit: limit,
				Skip:  skip,
				Total: total,
			}
			c.Request = request.AddToContext(c.Request, "pagination", paginationOpts)

			c.Next()
		}

	}
}
