// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IscsiServerInventoryView IscsiServer
type IscsiServerInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Ip string `json:"ip,omitempty"`
	Port int `json:"port,omitempty"`
	ChapUserName string `json:"chapUserName,omitempty"`
	ChapUserPassword string `json:"chapUserPassword,omitempty"`
	State string `json:"state,omitempty"`
	IscsiTargets []IscsiTargetInventoryView `json:"iscsiTargets,omitempty"`
	IscsiClusterRefs []IscsiServerClusterRefInventoryView `json:"iscsiClusterRefs,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

