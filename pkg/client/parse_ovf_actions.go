// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ParseOvf operates on ParseOvf
func (cli *ZSClient) ParseOvf(params param.ParseOvfParam) (*view.ParseOvfView, error) {
	resp := view.ParseOvfView{}
	if err := cli.Post("v1/ovf/parse", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
