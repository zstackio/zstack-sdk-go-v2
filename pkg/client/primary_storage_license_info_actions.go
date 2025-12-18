// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetPrimaryStorageLicenseInfo 获取PrimaryStorageLicenseInfo详情
func (cli *ZSClient) GetPrimaryStorageLicenseInfo(uuid string) (*view.GetPrimaryStorageLicenseInfoView, error) {
	var resp view.GetPrimaryStorageLicenseInfoView
	if err := cli.Get("v1/primary-storage/{uuid}/license", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

