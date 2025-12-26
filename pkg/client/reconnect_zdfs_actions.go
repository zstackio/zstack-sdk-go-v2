// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReconnectZdfs operates on ReconnectZdfs
func (cli *ZSClient) ReconnectZdfs(uuid string, params param.ReconnectZdfsParam) (*view.ReconnectZdfsEventView, error) {
	resp := view.ReconnectZdfsEventView{}
	if err := cli.Put("v1/zdfs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
