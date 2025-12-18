// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateBuildApp updates BuildApp
func (cli *ZSClient) UpdateBuildApp(uuid string, params param.UpdateBuildAppParam) (*view.UpdateBuildAppEventView, error) {
	resp := view.UpdateBuildAppEventView{}
	if err := cli.Put("v1/appcenter/buildapp/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
