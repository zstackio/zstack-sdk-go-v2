// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSNSSnmpPlatform creates SNSSnmpPlatform
func (cli *ZSClient) CreateSNSSnmpPlatform(params param.CreateSNSSnmpPlatformParam) (*view.CreateSNSApplicationPlatformEventView, error) {
	resp := view.CreateSNSApplicationPlatformEventView{}
	if err := cli.Post("v1/sns/application-platforms/snmp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
