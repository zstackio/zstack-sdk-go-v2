// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateScsiLun updates ScsiLun
func (cli *ZSClient) UpdateScsiLun(uuid string, params param.UpdateScsiLunParam) (*view.UpdateScsiLunEventView, error) {
	resp := view.UpdateScsiLunEventView{}
	if err := cli.Put("v1/storage-devices/scsi-lun/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
