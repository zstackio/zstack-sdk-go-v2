// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddLogServer adds LogServer
func (cli *ZSClient) AddLogServer(params param.AddLogServerParam) (*view.AddLogServerEventView, error) {
	resp := view.AddLogServerEventView{}
	if err := cli.Post("v1/log/servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
