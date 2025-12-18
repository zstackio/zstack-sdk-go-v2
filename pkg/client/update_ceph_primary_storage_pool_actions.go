// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateCephPrimaryStoragePool updates CephPrimaryStoragePool
func (cli *ZSClient) UpdateCephPrimaryStoragePool(uuid string, params param.UpdateCephPrimaryStoragePoolParam) (*view.UpdateCephPrimaryStoragePoolEventView, error) {
	resp := view.UpdateCephPrimaryStoragePoolEventView{}
	if err := cli.Put("v1/primary-storage/ceph/pools/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
