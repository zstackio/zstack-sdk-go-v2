// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetLicenseCapabilities 获取LicenseCapabilities详情
func (cli *ZSClient) GetLicenseCapabilities(uuid string) (*view.GetLicenseCapabilitiesView, error) {
	var resp view.GetLicenseCapabilitiesView
	if err := cli.Get("v1/licenses/capabilities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

