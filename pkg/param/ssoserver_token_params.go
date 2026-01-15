// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// RefreshSSOServerTokenParamDetail RefreshSSOServerToken detail param
type RefreshSSOServerTokenParamDetail struct {
	Token string `json:"token" validate:"required"`
	Duration int64 `json:"duration,omitempty"`
}

// RefreshSSOServerTokenParam RefreshSSOServerToken request param
type RefreshSSOServerTokenParam struct {
	BaseParam
	Params RefreshSSOServerTokenParamDetail `json:"refreshSSOServerToken"`
}
