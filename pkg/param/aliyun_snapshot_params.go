// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateAliyunSnapshotParamDetail UpdateAliyunSnapshot detail param
type UpdateAliyunSnapshotParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateAliyunSnapshotParam UpdateAliyunSnapshot request param
type UpdateAliyunSnapshotParam struct {
	BaseParam
	Params UpdateAliyunSnapshotParamDetail `json:"updateAliyunSnapshot"`
}
