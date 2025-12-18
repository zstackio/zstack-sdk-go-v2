// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSSORedirectTemplate updates SSORedirectTemplate
func (cli *ZSClient) UpdateSSORedirectTemplate(uuid string, params param.UpdateSSORedirectTemplateParam) (*view.UpdateSSORedirectTemplateEventView, error) {
	resp := view.UpdateSSORedirectTemplateEventView{}
	if err := cli.Put("v1/update/sso/redirectTemplate", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
