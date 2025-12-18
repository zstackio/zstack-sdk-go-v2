// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IscsiServerInventoryView IscsiServer
type IscsiServerInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"ip,omitempty"`
	rest int `json:"port,omitempty"`
	rest string `json:"chapUserName,omitempty"`
	rest string `json:"chapUserPassword,omitempty"`
	rest string `json:"state,omitempty"`
	rest []IscsiTargetInventoryView `json:"iscsiTargets,omitempty"`
	rest []IscsiServerClusterRefInventoryView `json:"iscsiClusterRefs,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

