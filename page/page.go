package page

import "math"

const (
	defaultPageSize uint64 = 20
	maxPageSize            = 1000
)

// Page 设置page
func Page(pageSize uint64, pageNo uint64) (uint64, uint64, uint64) {
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}
	if pageNo == 0 {
		pageNo = 1
	}

	// pageSize pageNo offset
	return pageSize, pageNo, (pageNo - 1) * pageSize
}

//SlicePage 数组分页
func SlicePage(page, pageSize, total int) (sliceStart, sliceEnd int) {
	if page <= 0 {
		page = 1
	}
	if pageSize < 0 {
		pageSize = 20 //设置一页默认显示的记录数
	}
	if pageSize > total {
		return 0, total
	}
	// 总页数
	pageCount := int(math.Ceil(float64(total) / float64(pageSize)))
	if page > pageCount {
		return 0, 0
	}
	sliceStart = (page - 1) * pageSize
	sliceEnd = sliceStart + pageSize

	if sliceEnd > total {
		sliceEnd = total
	}
	return sliceStart, sliceEnd
}
