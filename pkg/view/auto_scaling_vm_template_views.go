// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AutoScalingVmTemplateInventoryView AutoScalingVmTemplate
type AutoScalingVmTemplateInventoryView struct {
	rest string `json:"vmInstanceName,omitempty"`
	rest string `json:"vmInstanceType,omitempty"`
	rest string `json:"vmInstanceDescription,omitempty"`
	rest string `json:"vmInstanceOfferingUuid,omitempty"`
	rest string `json:"imageUuid,omitempty"`
	rest []string `json:"l3NetworkUuids,omitempty"`
	rest string `json:"rootDiskOfferingUuid,omitempty"`
	rest []string `json:"dataDiskOfferingUuids,omitempty"`
	rest string `json:"vmInstanceZoneUuid,omitempty"`
	rest string `json:"vmInstanceClusterUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"primaryStorageUuidForRootVolume,omitempty"`
	rest string `json:"defaultL3NetworkUuid,omitempty"`
	rest string `json:"strategy,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"state,omitempty"`
	rest []string `json:"systemTags,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

