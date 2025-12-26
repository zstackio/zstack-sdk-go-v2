// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeSNSApplicationEndpointState changes SNSApplicationEndpointState
func (cli *ZSClient) ChangeSNSApplicationEndpointState(uuid string, params param.ChangeSNSApplicationEndpointStateParam) (*view.ChangeSNSApplicationEndpointStateEventView, error) {
	resp := view.ChangeSNSApplicationEndpointStateEventView{}
	if err := cli.Put("v1/sns/application-endpoints/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
