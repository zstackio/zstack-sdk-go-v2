// Copyright (c) ZStack.io, Inc.

package param

// CreateAliyunNasMountTargetDetailParam CreateAliyunNasMountTarget detail param
type CreateAliyunNasMountTargetDetailParam struct {
	NasAccessGroupUuid string `json:"nasAccessGroupUuid" validate:"required"`
	VSwitchUuid string `json:"vSwitchUuid,omitempty"`
	NasFSUuid string `json:"nasFSUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAliyunNasMountTargetParam CreateAliyunNasMountTarget request param
type CreateAliyunNasMountTargetParam struct {
	BaseParam
	Params CreateAliyunNasMountTargetDetailParam `json:"params"`
}
