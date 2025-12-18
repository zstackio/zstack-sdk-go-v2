// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ExpungeDataVolume 操作DataVolume
func (cli *ZSClient) ExpungeDataVolume(uuid string, params param.ExpungeDataVolumeParam) (*view.ExpungeDataVolumeEventView, error) {
	resp := view.ExpungeDataVolumeEventView{}
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

