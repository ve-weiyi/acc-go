package dto

import "acc/app/model/entity"

type TodoReq struct {
	Todo entity.Todo   `json:"todo"`
	Tags []*entity.Tag `json:"tags"`
}
