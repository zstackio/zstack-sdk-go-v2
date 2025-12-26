// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSSORedirectTemplate creates SSORedirectTemplate
func (cli *ZSClient) CreateSSORedirectTemplate(params param.CreateSSORedirectTemplateParam) (*view.CreateSSORedirectTemplateEventView, error) {
	resp := view.CreateSSORedirectTemplateEventView{}
	if err := cli.Post("v1/create/sso/redirect/template/", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
