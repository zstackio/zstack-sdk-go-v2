// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncLicenseCapacity operates on SyncLicenseCapacity
func (cli *ZSClient) SyncLicenseCapacity(uuid string, params param.SyncLicenseCapacityParam) (*view.SyncLicenseCapacityEventView, error) {
	resp := view.SyncLicenseCapacityEventView{}
	if err := cli.Put("v1/license-server/authorized-capacity/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
