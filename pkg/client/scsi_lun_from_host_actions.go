// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachScsiLunFromHost 操作ScsiLunFromHost
func (cli *ZSClient) DetachScsiLunFromHost(uuid string, params param.DetachScsiLunFromHostParam) (*view.DetachScsiLunFromHostEventView, error) {
	resp := view.DetachScsiLunFromHostEventView{}
	if err := cli.Put("v1/storage-devices/scsi-lun/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

