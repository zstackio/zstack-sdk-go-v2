// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckScsiLunClusterStatus 操作CheckScsiLunClusterStatus
func (cli *ZSClient) CheckScsiLunClusterStatus(uuid string, params param.CheckScsiLunClusterStatusParam) (*view.CheckScsiLunClusterStatusView, error) {
	resp := view.CheckScsiLunClusterStatusView{}
	if err := cli.Put("v1/storage-devices/scsi-lun/{uuid}/cluster/{clusterUuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

