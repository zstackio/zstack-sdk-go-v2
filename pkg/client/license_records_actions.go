// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetLicenseRecords 获取LicenseRecords详情
func (cli *ZSClient) GetLicenseRecords(uuid string) (*view.GetLicenseRecordsView, error) {
	var resp view.GetLicenseRecordsView
	if err := cli.Get("v1/licenses/records", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

