// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PolicyInventoryView Policy
type PolicyInventoryView struct {
	BaseInfoView
	BaseTimeView
	Statements []PolicyStatementView `json:"statements,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
}

// DeletePolicyEventView DeletePolicyEvent
type DeletePolicyEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryPolicyView QueryPolicy
type QueryPolicyView struct {
	Inventories []PolicyInventoryView `json:"inventories,omitempty"`
}

// CreatePolicyEventView CreatePolicyEvent
type CreatePolicyEventView struct {
	Inventory PolicyInventoryView `json:"inventory,omitempty"`
}

// PolicyView Policy
type PolicyView struct {
	MetaClass interface{} `json:"metaClass,omitempty"`
	Statements []PolicyStatementView `json:"statements,omitempty"`
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
}

