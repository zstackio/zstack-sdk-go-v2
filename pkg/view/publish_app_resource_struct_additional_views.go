// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PublishAppResourceStructView PublishAppResourceStruct
type PublishAppResourceStructView struct {
	AppUuid      string `json:"appUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
}
