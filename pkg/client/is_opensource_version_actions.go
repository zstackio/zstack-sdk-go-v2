// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// IsOpensourceVersion operates on IsOpensourceVersion
func (cli *ZSClient) IsOpensourceVersion(params param.IsOpensourceVersionParam) (*view.IsOpensourceVersionView, error) {
	var resp view.IsOpensourceVersionView
	if err := cli.Get("v1/meta-data/opensource", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
