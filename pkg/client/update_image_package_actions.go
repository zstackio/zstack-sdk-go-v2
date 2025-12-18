// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateImagePackage updates ImagePackage
func (cli *ZSClient) UpdateImagePackage(uuid string, params param.UpdateImagePackageParam) (*view.UpdateImagePackageEventView, error) {
	resp := view.UpdateImagePackageEventView{}
	if err := cli.Put("v1/image-packages/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
