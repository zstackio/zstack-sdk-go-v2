// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// PublishAppParamDetail PublishApp detail param
type PublishAppParamDetail struct {
	BuildAppUuid string `json:"buildAppUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Parameters *string `json:"parameters,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// PublishAppParam PublishApp request param
type PublishAppParam struct {
	BaseParam
	Params PublishAppParamDetail `json:"params"`
}
// UpdatePublishAppParamDetail UpdatePublishApp detail param
type UpdatePublishAppParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdatePublishAppParam UpdatePublishApp request param
type UpdatePublishAppParam struct {
	BaseParam
	Params UpdatePublishAppParamDetail `json:"updatePublishApp"`
}
// DeletePublishAppParamDetail DeletePublishApp detail param
type DeletePublishAppParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeletePublishAppParam DeletePublishApp request param
type DeletePublishAppParam struct {
	BaseParam
	Params DeletePublishAppParamDetail `json:"deletePublishApp"`
}
