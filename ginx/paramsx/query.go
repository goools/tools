package paramsx

func ShiftPage(pageIndex, pageSize int64) (newPageIndex, newPageSize int64) {
	newPageIndex, newPageSize = pageIndex, pageSize
	if pageIndex < 1 {
		newPageIndex = 1
	}
	if pageSize < 10 {
		newPageSize = 10
	}
	return
}
