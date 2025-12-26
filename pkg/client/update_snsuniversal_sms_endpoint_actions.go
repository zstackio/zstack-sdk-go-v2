// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSNSUniversalSmsEndpoint updates SNSUniversalSmsEndpoint
func (cli *ZSClient) UpdateSNSUniversalSmsEndpoint(uuid string, params param.UpdateSNSUniversalSmsEndpointParam) (*view.UpdateSNSApplicationEndpointEventView, error) {
	resp := view.UpdateSNSApplicationEndpointEventView{}
	if err := cli.Put("v1/sns/application-endpoints/universal-sms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
