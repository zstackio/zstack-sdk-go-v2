// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSSOClientAttribute 更新SSOClientAttribute
func (cli *ZSClient) UpdateSSOClientAttribute(uuid string, params param.UpdateSSOClientAttributeParam) (*view.UpdateSSOClientAttributeEventView, error) {
	resp := view.UpdateSSOClientAttributeEventView{}
	if err := cli.Put("v1/sso/client/attributes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

