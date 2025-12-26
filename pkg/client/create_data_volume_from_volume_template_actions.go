// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateDataVolumeFromVolumeTemplate creates DataVolumeFromVolumeTemplate
func (cli *ZSClient) CreateDataVolumeFromVolumeTemplate(params param.CreateDataVolumeFromVolumeTemplateParam) (*view.CreateDataVolumeFromVolumeTemplateEventView, error) {
	resp := view.CreateDataVolumeFromVolumeTemplateEventView{}
	if err := cli.Post("v1/volumes/data/from/data-volume-templates/{imageUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
