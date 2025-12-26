// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddLogServer adds LogServer
func (cli *ZSClient) AddLogServer(params param.AddLogServerParam) (*view.AddLogServerEventView, error) {
	resp := view.AddLogServerEventView{}
	if err := cli.Post("v1/log/servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
