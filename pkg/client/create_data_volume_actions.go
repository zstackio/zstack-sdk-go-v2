// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateDataVolume creates DataVolume
func (cli *ZSClient) CreateDataVolume(params param.CreateDataVolumeParam) (*view.CreateDataVolumeEventView, error) {
	resp := view.CreateDataVolumeEventView{}
	if err := cli.Post("v1/volumes/data", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
