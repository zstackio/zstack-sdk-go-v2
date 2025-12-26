// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateLogServer updates LogServer
func (cli *ZSClient) UpdateLogServer(uuid string, params param.UpdateLogServerParam) (*view.UpdateLogServerEventView, error) {
	resp := view.UpdateLogServerEventView{}
	if err := cli.Put("v1/log/servers", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
