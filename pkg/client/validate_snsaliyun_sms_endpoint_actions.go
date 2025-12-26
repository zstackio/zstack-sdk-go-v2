// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ValidateSNSAliyunSmsEndpoint operates on ValidateSNSAliyunSmsEndpoint
func (cli *ZSClient) ValidateSNSAliyunSmsEndpoint(uuid string, params param.ValidateSNSAliyunSmsEndpointParam) (*view.ValidateSNSAliyunSmsEndpointEventView, error) {
	resp := view.ValidateSNSAliyunSmsEndpointEventView{}
	if err := cli.Put("v1/sns/sms-endpoints/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
