// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeDiskOfferingState changes DiskOfferingState
func (cli *ZSClient) ChangeDiskOfferingState(uuid string, params param.ChangeDiskOfferingStateParam) (*view.ChangeDiskOfferingStateEventView, error) {
	resp := view.ChangeDiskOfferingStateEventView{}
	if err := cli.Put("v1/disk-offerings/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
