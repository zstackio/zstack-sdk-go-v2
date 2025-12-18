// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddCephPrimaryStorage 操作AddCephPrimaryStorage
func (cli *ZSClient) AddCephPrimaryStorage(params param.AddCephPrimaryStorageParam) (*view.AddPrimaryStorageEventView, error) {
	resp := view.AddPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/ceph", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

