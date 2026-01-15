// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// AddAliyunNasMountTargetParamDetail AddAliyunNasMountTarget detail param
type AddAliyunNasMountTargetParamDetail struct {
	NasFSUuid string `json:"nasFSUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	MountDomain string `json:"mountDomain" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAliyunNasMountTargetParam AddAliyunNasMountTarget request param
type AddAliyunNasMountTargetParam struct {
	BaseParam
	AddAliyunNasMountTarget AddAliyunNasMountTargetParamDetail `json:"addAliyunNasMountTarget"`
}
// CreateAliyunNasMountTargetParamDetail CreateAliyunNasMountTarget detail param
type CreateAliyunNasMountTargetParamDetail struct {
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
	CreateAliyunNasMountTarget CreateAliyunNasMountTargetParamDetail `json:"createAliyunNasMountTarget"`
}
