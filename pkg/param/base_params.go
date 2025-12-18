// Copyright (c) ZStack.io, Inc.

package param

// BaseParam 基础参数，包含通用标签
type BaseParam struct {
	SystemTags []string `json:"systemTags,omitempty"` // 系统标签
	UserTags   []string `json:"userTags,omitempty"`   // 用户标签
	RequestIp  string   `json:"requestIp,omitempty"`  // 请求IP
}

// QueryParam 查询参数
type QueryParam struct {
	Conditions []QueryCondition `json:"conditions,omitempty"` // 查询条件
	Limit      *int             `json:"limit,omitempty"`      // 返回数量限制
	Start      *int             `json:"start,omitempty"`      // 起始位置
	Count      bool             `json:"count,omitempty"`      // 是否只返回数量
	GroupBy    string           `json:"groupBy,omitempty"`    // 分组字段
	SortBy     string           `json:"sortBy,omitempty"`     // 排序字段
	SortDirection string        `json:"sortDirection,omitempty"` // 排序方向: asc/desc
	Fields     []string         `json:"fields,omitempty"`     // 返回字段
}

// QueryCondition 查询条件
type QueryCondition struct {
	Name  string `json:"name"`  // 字段名
	Op    string `json:"op"`    // 操作符: =, !=, >, <, >=, <=, in, not in, like, not like, is null, is not null
	Value string `json:"value"` // 值
}

// NewQueryParam 创建新的查询参数
func NewQueryParam() *QueryParam {
	return &QueryParam{}
}

// AddQ 添加查询条件
func (p *QueryParam) AddQ(condition string) *QueryParam {
	// 解析条件字符串，格式: "field=value" 或 "field!=value" 等
	p.Conditions = append(p.Conditions, QueryCondition{Value: condition})
	return p
}

// Limit 设置返回数量限制
func (p *QueryParam) Limit(limit int) *QueryParam {
	p.Limit = &limit
	return p
}

// Start 设置起始位置
func (p *QueryParam) Start(start int) *QueryParam {
	p.Start = &start
	return p
}

// DeleteMode 删除模式
type DeleteMode string

const (
	Permissive DeleteMode = "Permissive" // 宽松模式，允许删除有依赖的资源
	Enforcing  DeleteMode = "Enforcing"  // 强制模式，不允许删除有依赖的资源
)
