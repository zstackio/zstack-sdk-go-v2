// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateEmailMedia 创建EmailMedia
func (cli *ZSClient) CreateEmailMedia(params param.CreateEmailMediaParam) (*view.CreateMediaEventView, error) {
	resp := view.CreateMediaEventView{}
	if err := cli.Post("v1/media/emails", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

