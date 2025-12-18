// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddCephPrimaryStoragePool 操作AddCephPrimaryStoragePool
func (cli *ZSClient) AddCephPrimaryStoragePool(params param.AddCephPrimaryStoragePoolParam) (*view.AddCephPrimaryStoragePoolEventView, error) {
	resp := view.AddCephPrimaryStoragePoolEventView{}
	if err := cli.Post("v1/primary-storage/ceph/{primaryStorageUuid}/pools", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

