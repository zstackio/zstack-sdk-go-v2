// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CheckScsiLunClusterStatus operates on CheckScsiLunClusterStatus
func (cli *ZSClient) CheckScsiLunClusterStatus(uuid string, params param.CheckScsiLunClusterStatusParam) (*view.CheckScsiLunClusterStatusView, error) {
	resp := view.CheckScsiLunClusterStatusView{}
	if err := cli.Put("v1/storage-devices/scsi-lun/{uuid}/cluster/{clusterUuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
