// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateImagePackage updates ImagePackage
func (cli *ZSClient) UpdateImagePackage(uuid string, params param.UpdateImagePackageParam) (*view.UpdateImagePackageEventView, error) {
	resp := view.UpdateImagePackageEventView{}
	if err := cli.Put("v1/image-packages/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
