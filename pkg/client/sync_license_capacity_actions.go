// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncLicenseCapacity operates on SyncLicenseCapacity
func (cli *ZSClient) SyncLicenseCapacity(uuid string, params param.SyncLicenseCapacityParam) (*view.SyncLicenseCapacityEventView, error) {
	resp := view.SyncLicenseCapacityEventView{}
	if err := cli.Put("v1/license-server/authorized-capacity/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
