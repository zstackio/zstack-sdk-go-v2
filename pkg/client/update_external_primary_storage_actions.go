// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateExternalPrimaryStorage updates ExternalPrimaryStorage
func (cli *ZSClient) UpdateExternalPrimaryStorage(uuid string, params param.UpdateExternalPrimaryStorageParam) (*view.UpdateExternalPrimaryStorageEventView, error) {
	resp := view.UpdateExternalPrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/addon/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
