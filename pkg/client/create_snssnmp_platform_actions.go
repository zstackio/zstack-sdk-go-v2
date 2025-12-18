// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSNSSnmpPlatform creates SNSSnmpPlatform
func (cli *ZSClient) CreateSNSSnmpPlatform(params param.CreateSNSSnmpPlatformParam) (*view.CreateSNSApplicationPlatformEventView, error) {
	resp := view.CreateSNSApplicationPlatformEventView{}
	if err := cli.Post("v1/sns/application-platforms/snmp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
