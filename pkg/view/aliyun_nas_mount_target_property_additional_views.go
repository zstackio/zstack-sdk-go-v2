// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunNasMountTargetPropertyView AliyunNasMountTargetProperty
type AliyunNasMountTargetPropertyView struct {
	Status *string `json:"status,omitempty"`
	AccessGroupName *string `json:"accessGroupName,omitempty"`
	MountDomain *string `json:"mountDomain,omitempty"`
}

