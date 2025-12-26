// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateImageGroupFromImage creates ImageGroupFromImage
func (cli *ZSClient) CreateImageGroupFromImage(params param.CreateImageGroupFromImageParam) (*view.CreateImageGroupFromImageEventView, error) {
	resp := view.CreateImageGroupFromImageEventView{}
	if err := cli.Post("v1/imagegroup/from/image/{rootVolumeTemplateUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
