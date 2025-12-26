// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetPrimaryStorageLicenseInfo gets PrimaryStorageLicenseInfo by uuid
func (cli *ZSClient) GetPrimaryStorageLicenseInfo(uuid string) (*view.GetPrimaryStorageLicenseInfoView, error) {
	var resp view.GetPrimaryStorageLicenseInfoView
	if err := cli.Get("v1/primary-storage/{uuid}/license", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
