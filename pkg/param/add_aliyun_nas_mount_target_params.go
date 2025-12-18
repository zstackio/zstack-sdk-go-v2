// Copyright (c) ZStack.io, Inc.

package param

// AddAliyunNasMountTargetDetailParam AddAliyunNasMountTarget detail param
type AddAliyunNasMountTargetDetailParam struct {
	NasFSUuid string `json:"nasFSUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	MountDomain string `json:"mountDomain" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAliyunNasMountTargetParam AddAliyunNasMountTarget request param
type AddAliyunNasMountTargetParam struct {
	BaseParam
	Params AddAliyunNasMountTargetDetailParam `json:"params"`
}
