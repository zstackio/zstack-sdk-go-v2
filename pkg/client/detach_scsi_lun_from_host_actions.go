// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DetachScsiLunFromHost operates on ScsiLunFromHost
func (cli *ZSClient) DetachScsiLunFromHost(uuid string, params param.DetachScsiLunFromHostParam) (*view.DetachScsiLunFromHostEventView, error) {
	resp := view.DetachScsiLunFromHostEventView{}
	if err := cli.Put("v1/storage-devices/scsi-lun/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
