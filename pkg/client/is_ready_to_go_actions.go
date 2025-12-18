// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// IsReadyToGo 操作IsReadyToGo
func (cli *ZSClient) IsReadyToGo(params param.IsReadyToGoParam) (*view.IsReadyToGoView, error) {
	var resp view.IsReadyToGoView
	if err := cli.Get("v1/management-nodes/ready", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

