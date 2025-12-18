// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddMonToCephPrimaryStorage 操作AddMonToCephPrimaryStorage
func (cli *ZSClient) AddMonToCephPrimaryStorage(params param.AddMonToCephPrimaryStorageParam) (*view.AddMonToCephPrimaryStorageEventView, error) {
	resp := view.AddMonToCephPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/ceph/{uuid}/mons", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

