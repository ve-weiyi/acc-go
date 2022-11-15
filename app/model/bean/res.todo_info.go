package bean

import "acc/app/model/entity"

type TodoInfoResp struct {
	Username *string       `json:"username"`
	Tags     []*entity.Tag `json:"tags"`
	Details  *entity.Todo  `json:"details"`
}
