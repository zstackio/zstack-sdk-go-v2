// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSNSSnmpPlatform updates SNSSnmpPlatform
func (cli *ZSClient) UpdateSNSSnmpPlatform(uuid string, params param.UpdateSNSSnmpPlatformParam) (*view.UpdateSNSApplicationPlatformEventView, error) {
	resp := view.UpdateSNSApplicationPlatformEventView{}
	if err := cli.Put("v1/sns/application-platforms/snmp/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
