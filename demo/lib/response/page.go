package response

const (
	pageSize = 5
)

// 页码从0开始

type PageResult struct {
	Datas interface{} `json:"datas"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int64       `json:"total"`
}
