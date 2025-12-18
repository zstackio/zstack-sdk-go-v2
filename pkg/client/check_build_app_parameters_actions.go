// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckBuildAppParameters 操作CheckBuildAppParameters
func (cli *ZSClient) CheckBuildAppParameters(params param.CheckBuildAppParametersParam) (*view.CheckBuildAppParametersView, error) {
	resp := view.CheckBuildAppParametersView{}
	if err := cli.Post("v1/appcenter/buildapp/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

