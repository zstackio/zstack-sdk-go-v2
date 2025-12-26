// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetResourceFromPublishApp gets ResourceFromPublishApp by uuid
func (cli *ZSClient) GetResourceFromPublishApp(uuid string) (*view.GetResourceFromPublishAppView, error) {
	var resp view.GetResourceFromPublishAppView
	if err := cli.Get("v1/appcenter/app/resources", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
