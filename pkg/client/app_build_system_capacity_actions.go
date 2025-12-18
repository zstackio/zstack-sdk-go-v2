// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetAppBuildSystemCapacity 获取AppBuildSystemCapacity详情
func (cli *ZSClient) GetAppBuildSystemCapacity(uuid string) (*view.GetAppBuildSystemCapacityView, error) {
	var resp view.GetAppBuildSystemCapacityView
	if err := cli.Get("v1/appcenter/buildsystem/capacities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

