// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ModelServiceGroupModelServiceRefInventoryView ModelServiceGroupModelServiceRef
type ModelServiceGroupModelServiceRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	ModelServiceInstanceGroupUuid *string `json:"modelServiceInstanceGroupUuid,omitempty"`
	DependModelServiceInstanceGroupUuid *string `json:"dependModelServiceInstanceGroupUuid,omitempty"`
}

