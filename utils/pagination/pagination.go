package pagination

import (
	"errors"
	"strconv"
)

func Validate(skip int, total int) error {
	if skip > total {
		return errors.New("page out of range")
	}
	return nil
}

func Params(pageQuery string) (page int, skip int, limit int) {
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
