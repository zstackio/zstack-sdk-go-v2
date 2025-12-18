// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ValidateSNSAliyunSmsEndpoint 操作ValidateSNSAliyunSmsEndpoint
func (cli *ZSClient) ValidateSNSAliyunSmsEndpoint(uuid string, params param.ValidateSNSAliyunSmsEndpointParam) (*view.ValidateSNSAliyunSmsEndpointEventView, error) {
	resp := view.ValidateSNSAliyunSmsEndpointEventView{}
	if err := cli.Put("v1/sns/sms-endpoints/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

