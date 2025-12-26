// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSSOClientAttribute updates SSOClientAttribute
func (cli *ZSClient) UpdateSSOClientAttribute(uuid string, params param.UpdateSSOClientAttributeParam) (*view.UpdateSSOClientAttributeEventView, error) {
	resp := view.UpdateSSOClientAttributeEventView{}
	if err := cli.Put("v1/sso/client/attributes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
