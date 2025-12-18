// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PushLicenseAddOnsUsage 操作PushLicenseAddOnsUsage
func (cli *ZSClient) PushLicenseAddOnsUsage(uuid string, params param.PushLicenseAddOnsUsageParam) (*view.PushLicenseAddOnsUsageEventView, error) {
	resp := view.PushLicenseAddOnsUsageEventView{}
	if err := cli.Put("v1/licenses/addons/usage", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

