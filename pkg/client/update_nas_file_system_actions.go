// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateNasFileSystem updates NasFileSystem
func (cli *ZSClient) UpdateNasFileSystem(uuid string, params param.UpdateNasFileSystemParam) (*view.UpdateNasFileSystemEventView, error) {
	resp := view.UpdateNasFileSystemEventView{}
	if err := cli.Put("v1/primary-storage/nas/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
