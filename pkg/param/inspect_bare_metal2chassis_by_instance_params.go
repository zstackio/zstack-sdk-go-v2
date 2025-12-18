// Copyright (c) ZStack.io, Inc.

package param

// InspectBareMetal2ChassisByInstanceDetailParam InspectBareMetal2ChassisByInstance详细参数
type InspectBareMetal2ChassisByInstanceDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// InspectBareMetal2ChassisByInstanceParam InspectBareMetal2ChassisByInstance请求参数
type InspectBareMetal2ChassisByInstanceParam struct {
	BaseParam
	Params InspectBareMetal2ChassisByInstanceDetailParam `json:"params"` // 详细参数
}

