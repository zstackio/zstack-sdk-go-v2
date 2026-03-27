// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// NginxRedirectRuleView NginxRedirectRule
type NginxRedirectRuleView struct {
	DestUrl string `json:"destUrl,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	JupyterUrl string `json:"jupyterUrl,omitempty"`
	CurrentUrl string `json:"currentUrl,omitempty"`
	CurrentJupyterUrl string `json:"currentJupyterUrl,omitempty"`
	OverriddenUuid string `json:"overriddenUuid,omitempty"`
}

