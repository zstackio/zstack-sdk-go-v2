// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CloneImage operates on Image
func (cli *ZSClient) CloneImage(params param.CloneImageParam) (*view.CloneImageEventView, error) {
	resp := view.CloneImageEventView{}
	if err := cli.Post("v1/image/clone/{imageUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
