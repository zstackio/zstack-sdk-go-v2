// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PolicyView Policy
type PolicyView struct {
	MetaClass interface{} `json:"metaClass,omitempty"`
	Statements []interface{} `json:"statements,omitempty"`
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
}

// PolicyInventoryView Policy
type PolicyInventoryView struct {
	Statements []interface{} `json:"statements,omitempty"`
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
}

