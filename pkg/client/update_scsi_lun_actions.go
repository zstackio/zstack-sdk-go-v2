// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateScsiLun updates ScsiLun
func (cli *ZSClient) UpdateScsiLun(uuid string, params param.UpdateScsiLunParam) (*view.UpdateScsiLunEventView, error) {
	resp := view.UpdateScsiLunEventView{}
	if err := cli.Put("v1/storage-devices/scsi-lun/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
