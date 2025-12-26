// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddDataCenterFromRemote adds DataCenterFromRemote
func (cli *ZSClient) AddDataCenterFromRemote(params param.AddDataCenterFromRemoteParam) (*view.AddDataCenterFromRemoteEventView, error) {
	resp := view.AddDataCenterFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/data-center", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
