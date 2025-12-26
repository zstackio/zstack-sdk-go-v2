// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// DataCenterPropertyView DataCenterProperty
type DataCenterPropertyView struct {
	RegionId string `json:"regionId,omitempty"`
	RegionName string `json:"regionName,omitempty"`
}

