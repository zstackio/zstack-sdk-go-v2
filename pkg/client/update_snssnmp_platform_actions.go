// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSNSSnmpPlatform updates SNSSnmpPlatform
func (cli *ZSClient) UpdateSNSSnmpPlatform(uuid string, params param.UpdateSNSSnmpPlatformParam) (*view.UpdateSNSApplicationPlatformEventView, error) {
	resp := view.UpdateSNSApplicationPlatformEventView{}
	if err := cli.Put("v1/sns/application-platforms/snmp/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
