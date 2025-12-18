// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddDataCenterFromRemote 操作AddDataCenterFromRemote
func (cli *ZSClient) AddDataCenterFromRemote(params param.AddDataCenterFromRemoteParam) (*view.AddDataCenterFromRemoteEventView, error) {
	resp := view.AddDataCenterFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/data-center", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

