// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachScsiLunToVmInstance operates on ScsiLunToVmInstance
func (cli *ZSClient) AttachScsiLunToVmInstance(params param.AttachScsiLunToVmInstanceParam) (*view.AttachScsiLunToVmInstanceEventView, error) {
	resp := view.AttachScsiLunToVmInstanceEventView{}
	if err := cli.Post("v1/vm-instances/{vmInstanceUuid}/scsi-lun/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
