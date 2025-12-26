// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CheckBuildAppParameters operates on CheckBuildAppParameters
func (cli *ZSClient) CheckBuildAppParameters(params param.CheckBuildAppParametersParam) (*view.CheckBuildAppParametersView, error) {
	resp := view.CheckBuildAppParametersView{}
	if err := cli.Post("v1/appcenter/buildapp/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
