// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSNSWeComEndpoint updates SNSWeComEndpoint
func (cli *ZSClient) UpdateSNSWeComEndpoint(uuid string, params param.UpdateSNSWeComEndpointParam) (*view.UpdateSNSApplicationEndpointEventView, error) {
	resp := view.UpdateSNSApplicationEndpointEventView{}
	if err := cli.Put("v1/sns/application-endpoints/we-com/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
