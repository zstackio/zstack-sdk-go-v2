// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetLicenseUKeyStatus 获取LicenseUKeyStatus详情
func (cli *ZSClient) GetLicenseUKeyStatus(uuid string) (*view.GetLicenseUKeyStatusEventView, error) {
	var resp view.GetLicenseUKeyStatusEventView
	if err := cli.Get("v1/licenses/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

