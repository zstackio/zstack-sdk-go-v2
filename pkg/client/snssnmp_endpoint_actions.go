// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSNSSnmpEndpoint 创建SNSSnmpEndpoint
func (cli *ZSClient) CreateSNSSnmpEndpoint(params param.CreateSNSSnmpEndpointParam) (*view.CreateSNSApplicationEndpointEventView, error) {
	resp := view.CreateSNSApplicationEndpointEventView{}
	if err := cli.Post("v1/sns/application-endpoints/snmp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

