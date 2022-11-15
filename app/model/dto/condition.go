package dto

type Condition struct {
	Uid       int    `json:"uid"example:""`
	Username  string `json:"username"example:""`
	Keywords  string `json:"keywords" example:""`
	StartTime string `json:"startTime" example:""`
	EndTime   string `json:"endTime" example:""`
	//Page      int    `json:"page" example:"0"`
}
