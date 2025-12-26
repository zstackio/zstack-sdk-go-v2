// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetPrimaryStorageTypes gets PrimaryStorageTypes by uuid
func (cli *ZSClient) GetPrimaryStorageTypes(uuid string) (*view.GetPrimaryStorageTypesView, error) {
	var resp view.GetPrimaryStorageTypesView
	if err := cli.Get("v1/primary-storage/types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
