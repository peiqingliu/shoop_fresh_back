package query

type ListQuery struct {
	PageSize int `json:"pageSize"` // 每页多少条记录
	PageNum int `json:"page"` // 那一页
}
