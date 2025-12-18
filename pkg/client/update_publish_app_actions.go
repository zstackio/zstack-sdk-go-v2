// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdatePublishApp updates PublishApp
func (cli *ZSClient) UpdatePublishApp(uuid string, params param.UpdatePublishAppParam) (*view.UpdatePublishAppEventView, error) {
	resp := view.UpdatePublishAppEventView{}
	if err := cli.Put("v1/appcenter/app/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
