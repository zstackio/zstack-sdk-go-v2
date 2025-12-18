// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// BuildApplicationInventoryView BuildApplication
type BuildApplicationInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	BuildSystemUuid string `json:"buildSystemUuid,omitempty"`
	TemplateContent string `json:"templateContent,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	AppMetaData string `json:"appMetaData,omitempty"`
	AppId string `json:"appId,omitempty"`
	Version string `json:"version,omitempty"`
	Status string `json:"status,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

