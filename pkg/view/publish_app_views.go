// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PublishAppInventoryView PublishApp
type PublishAppInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"templateContent,omitempty"`
	rest string `json:"appMetaData,omitempty"`
	rest string `json:"preParams,omitempty"`
	rest string `json:"vmRelationShip,omitempty"`
	rest string `json:"buildAppUuid,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"appId,omitempty"`
	rest string `json:"version,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

