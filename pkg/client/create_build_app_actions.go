// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateBuildApp creates BuildApp
func (cli *ZSClient) CreateBuildApp(params param.CreateBuildAppParam) (*view.CreateBuildAppEventView, error) {
	resp := view.CreateBuildAppEventView{}
	if err := cli.Post("v1/appcenter/buildapp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
