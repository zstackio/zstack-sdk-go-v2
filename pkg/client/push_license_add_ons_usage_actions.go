// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PushLicenseAddOnsUsage operates on PushLicenseAddOnsUsage
func (cli *ZSClient) PushLicenseAddOnsUsage(uuid string, params param.PushLicenseAddOnsUsageParam) (*view.PushLicenseAddOnsUsageEventView, error) {
	resp := view.PushLicenseAddOnsUsageEventView{}
	if err := cli.Put("v1/licenses/addons/usage", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
