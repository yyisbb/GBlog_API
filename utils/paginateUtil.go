package utils

import (
	"gorm.io/gorm"
	"strconv"
)

// Paginate: gorm page
func Paginate(page, limit string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		page, _ := strconv.Atoi(page)
		if page == 0 {
			page = 1
		}

		// pageSize, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		pageSize, _ := strconv.Atoi(limit)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func ComputeCount(count int64) int64 {
	if count < 10 {
		return 1
	}

	if count%10 == 0 {
		return count / 10
	} else {
		return count/10 + 1
	}
}
