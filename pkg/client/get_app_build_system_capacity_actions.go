// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetAppBuildSystemCapacity gets AppBuildSystemCapacity by uuid
func (cli *ZSClient) GetAppBuildSystemCapacity(uuid string) (*view.GetAppBuildSystemCapacityView, error) {
	var resp view.GetAppBuildSystemCapacityView
	if err := cli.Get("v1/appcenter/buildsystem/capacities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
