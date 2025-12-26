// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateDataVolumeTemplateFromVolume creates DataVolumeTemplateFromVolume
func (cli *ZSClient) CreateDataVolumeTemplateFromVolume(params param.CreateDataVolumeTemplateFromVolumeParam) (*view.CreateDataVolumeTemplateFromVolumeEventView, error) {
	resp := view.CreateDataVolumeTemplateFromVolumeEventView{}
	if err := cli.Post("v1/images/data-volume-templates/from/volumes/{volumeUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
