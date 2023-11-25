package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func PaginationFrom(ctx *gin.Context) (int64, int64) {
	pageQuery := ctx.DefaultQuery("page", "1")

	pageSizeQuery := ctx.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageQuery)

	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeQuery)

	if err != nil {
		pageSize = 10
	}

	page = max(page, 1)

	pageSize = max(pageSize, 1)

	return int64((page - 1) * pageSize), int64(pageSize)
}
