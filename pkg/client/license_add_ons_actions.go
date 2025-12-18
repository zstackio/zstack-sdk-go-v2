// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetLicenseAddOns 获取LicenseAddOns详情
func (cli *ZSClient) GetLicenseAddOns(uuid string) (*view.GetLicenseAddOnsView, error) {
	var resp view.GetLicenseAddOnsView
	if err := cli.Get("v1/licenses/addons", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

