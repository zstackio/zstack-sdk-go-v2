// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddImage 操作AddImage
func (cli *ZSClient) AddImage(params param.AddImageParam) (*view.AddImageEventView, error) {
	resp := view.AddImageEventView{}
	if err := cli.Post("v1/images", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

