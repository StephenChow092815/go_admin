package model

type Pagination struct {
	// list：返回数据
	List interface{}
	// current：当前页
	Current int
	// size：每页数量
	Size int
	// pages：总页数
	Pages int64
	// total：总记录数
	Total int64
}
