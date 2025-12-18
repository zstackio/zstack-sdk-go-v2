// Copyright (c) ZStack.io, Inc.

package view

import "time"

// FcHbaDeviceInventoryView FcHbaDevice
type FcHbaDeviceInventoryView struct {
	rest string `json:"portName,omitempty"`
	rest string `json:"portState,omitempty"`
	rest string `json:"speed,omitempty"`
	rest string `json:"supportedSpeeds,omitempty"`
	rest string `json:"symbolicName,omitempty"`
	rest string `json:"supportedClasses,omitempty"`
	rest string `json:"nodeName,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"hbaType,omitempty"`
	rest string `json:"createDate,omitempty"`
	rest string `json:"lastOpDate,omitempty"`
}

