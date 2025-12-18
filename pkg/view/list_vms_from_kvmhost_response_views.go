// Copyright (c) ZStack.io, Inc.

package view

// ListVMsFromKVMHostEventView ListVMsFromKVMHostEvent
type ListVMsFromKVMHostEventView struct {
	Inventories []VmInstanceInventoryView `json:"inventories,omitempty"`
	LibvirtVersion string `json:"libvirtVersion,omitempty"`
	QemuVersion string `json:"qemuVersion,omitempty"`
	V2vCaps map[string]bool `json:"v2vCaps,omitempty"`
	Success bool `json:"success,omitempty"`
}

