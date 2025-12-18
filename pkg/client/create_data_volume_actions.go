// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateDataVolume creates DataVolume
func (cli *ZSClient) CreateDataVolume(params param.CreateDataVolumeParam) (*view.CreateDataVolumeEventView, error) {
	resp := view.CreateDataVolumeEventView{}
	if err := cli.Post("v1/volumes/data", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
