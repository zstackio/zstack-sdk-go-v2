// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVolumeQos operates on SetVolumeQos
func (cli *ZSClient) SetVolumeQos(uuid string, params param.SetVolumeQosParam) (*view.SetVolumeQosEventView, error) {
	resp := view.SetVolumeQosEventView{}
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
