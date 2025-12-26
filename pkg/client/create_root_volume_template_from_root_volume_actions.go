// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateRootVolumeTemplateFromRootVolume creates RootVolumeTemplateFromRootVolume
func (cli *ZSClient) CreateRootVolumeTemplateFromRootVolume(params param.CreateRootVolumeTemplateFromRootVolumeParam) (*view.CreateRootVolumeTemplateFromRootVolumeEventView, error) {
	resp := view.CreateRootVolumeTemplateFromRootVolumeEventView{}
	if err := cli.Post("v1/images/root-volume-templates/from/volumes/{rootVolumeUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
