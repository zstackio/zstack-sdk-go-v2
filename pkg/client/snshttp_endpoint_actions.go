// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSNSHttpEndpoint 更新SNSHttpEndpoint
func (cli *ZSClient) UpdateSNSHttpEndpoint(uuid string, params param.UpdateSNSHttpEndpointParam) (*view.UpdateSNSApplicationEndpointEventView, error) {
	resp := view.UpdateSNSApplicationEndpointEventView{}
	if err := cli.Put("v1/sns/application-endpoints/http/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

