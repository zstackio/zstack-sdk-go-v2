// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateDataVolumeTemplateFromVolume 创建DataVolumeTemplateFromVolume
func (cli *ZSClient) CreateDataVolumeTemplateFromVolume(params param.CreateDataVolumeTemplateFromVolumeParam) (*view.CreateDataVolumeTemplateFromVolumeEventView, error) {
	resp := view.CreateDataVolumeTemplateFromVolumeEventView{}
	if err := cli.Post("v1/images/data-volume-templates/from/volumes/{volumeUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

