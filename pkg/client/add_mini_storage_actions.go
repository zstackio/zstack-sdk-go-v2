// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddMiniStorage adds MiniStorage
func (cli *ZSClient) AddMiniStorage(params param.AddMiniStorageParam) (*view.AddPrimaryStorageEventView, error) {
	resp := view.AddPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/mini", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
