// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSNSUniversalSmsEndpoint 更新SNSUniversalSmsEndpoint
func (cli *ZSClient) UpdateSNSUniversalSmsEndpoint(uuid string, params param.UpdateSNSUniversalSmsEndpointParam) (*view.UpdateSNSApplicationEndpointEventView, error) {
	resp := view.UpdateSNSApplicationEndpointEventView{}
	if err := cli.Put("v1/sns/application-endpoints/universal-sms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

