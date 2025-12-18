// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachScsiLunToVmInstance operates on ScsiLunToVmInstance
func (cli *ZSClient) AttachScsiLunToVmInstance(params param.AttachScsiLunToVmInstanceParam) (*view.AttachScsiLunToVmInstanceEventView, error) {
	resp := view.AttachScsiLunToVmInstanceEventView{}
	if err := cli.Post("v1/vm-instances/{vmInstanceUuid}/scsi-lun/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
