// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSSORedirectTemplate creates SSORedirectTemplate
func (cli *ZSClient) CreateSSORedirectTemplate(params param.CreateSSORedirectTemplateParam) (*view.CreateSSORedirectTemplateEventView, error) {
	resp := view.CreateSSORedirectTemplateEventView{}
	if err := cli.Post("v1/create/sso/redirect/template/", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
