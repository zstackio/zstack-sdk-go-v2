// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PolicyStatementView PolicyStatement
type PolicyStatementView struct {
	Name string `json:"name,omitempty"`
	Effect string `json:"effect,omitempty"`
	Principals []string `json:"principals,omitempty"`
	Actions []string `json:"actions,omitempty"`
	Resources []string `json:"resources,omitempty"`
}

