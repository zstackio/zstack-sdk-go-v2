// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SdnVlanRangeView SdnVlanRange
type SdnVlanRangeView struct {
	StartVlan int `json:"startVlan,omitempty"`
	EndVlan   int `json:"endVlan,omitempty"`
}
