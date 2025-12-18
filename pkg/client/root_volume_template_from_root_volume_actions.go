// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateRootVolumeTemplateFromRootVolume 创建RootVolumeTemplateFromRootVolume
func (cli *ZSClient) CreateRootVolumeTemplateFromRootVolume(params param.CreateRootVolumeTemplateFromRootVolumeParam) (*view.CreateRootVolumeTemplateFromRootVolumeEventView, error) {
	resp := view.CreateRootVolumeTemplateFromRootVolumeEventView{}
	if err := cli.Post("v1/images/root-volume-templates/from/volumes/{rootVolumeUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

