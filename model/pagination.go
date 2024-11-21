package model

type Pagination struct {
	// list：返回数据
	List interface{} `json:"list"`
	// current：当前页
	Current int `json:"current"`
	// size：每页数量
	Size int `json:"size"`
	// pages：总页数
	Pages int64 `json:"pages"`
	// total：总记录数
	Total int64 `json:"total"`
}
