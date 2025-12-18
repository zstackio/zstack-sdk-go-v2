// Copyright (c) ZStack.io, Inc.

package param

// BaseParam base parameters
type BaseParam struct {
	SystemTags []string `json:"systemTags,omitempty"`
	UserTags   []string `json:"userTags,omitempty"`
	RequestIp  string   `json:"requestIp,omitempty"`
}

// QueryParam query parameters
type QueryParam struct {
	Conditions     []QueryCondition `json:"conditions,omitempty"`
	LimitNum       *int             `json:"limit,omitempty"`
	StartNum       *int             `json:"start,omitempty"`
	Count          bool             `json:"count,omitempty"`
	GroupBy        string           `json:"groupBy,omitempty"`
	SortBy         string           `json:"sortBy,omitempty"`
	SortDirection  string           `json:"sortDirection,omitempty"`
	Fields         []string         `json:"fields,omitempty"`
	ReplyWithCount bool             `json:"replyWithCount,omitempty"`
}

// QueryCondition query condition
type QueryCondition struct {
	Name  string `json:"name"`
	Op    string `json:"op"`
	Value string `json:"value"`
}

// NewQueryParam creates a new query param
func NewQueryParam() *QueryParam {
	return &QueryParam{}
}

// AddQ adds a query condition
func (p *QueryParam) AddQ(condition string) *QueryParam {
	p.Conditions = append(p.Conditions, QueryCondition{Value: condition})
	return p
}

// Limit sets the limit
func (p *QueryParam) Limit(limit int) *QueryParam {
	p.LimitNum = &limit
	return p
}

// Start sets the start
func (p *QueryParam) Start(start int) *QueryParam {
	p.StartNum = &start
	return p
}

// DeleteMode delete mode
type DeleteMode string

const (
	Permissive DeleteMode = "Permissive"
	Enforcing  DeleteMode = "Enforcing"
)
