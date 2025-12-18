// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateNasFileSystem 更新NasFileSystem
func (cli *ZSClient) UpdateNasFileSystem(uuid string, params param.UpdateNasFileSystemParam) (*view.UpdateNasFileSystemEventView, error) {
	resp := view.UpdateNasFileSystemEventView{}
	if err := cli.Put("v1/primary-storage/nas/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

