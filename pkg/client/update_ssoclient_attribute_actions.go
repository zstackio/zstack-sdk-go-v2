// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSSOClientAttribute updates SSOClientAttribute
func (cli *ZSClient) UpdateSSOClientAttribute(uuid string, params param.UpdateSSOClientAttributeParam) (*view.UpdateSSOClientAttributeEventView, error) {
	resp := view.UpdateSSOClientAttributeEventView{}
	if err := cli.Put("v1/sso/client/attributes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
