// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RecoverImage operates on Image
func (cli *ZSClient) RecoverImage(uuid string, params param.RecoverImageParam) (*view.RecoverImageEventView, error) {
	resp := view.RecoverImageEventView{}
	if err := cli.Put("v1/images/{imageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
