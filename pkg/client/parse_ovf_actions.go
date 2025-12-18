// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ParseOvf 操作ParseOvf
func (cli *ZSClient) ParseOvf(params param.ParseOvfParam) (*view.ParseOvfView, error) {
	resp := view.ParseOvfView{}
	if err := cli.Post("v1/ovf/parse", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

