// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeSNSApplicationEndpointState 操作SNSApplicationEndpointState
func (cli *ZSClient) ChangeSNSApplicationEndpointState(uuid string, params param.ChangeSNSApplicationEndpointStateParam) (*view.ChangeSNSApplicationEndpointStateEventView, error) {
	resp := view.ChangeSNSApplicationEndpointStateEventView{}
	if err := cli.Put("v1/sns/application-endpoints/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

