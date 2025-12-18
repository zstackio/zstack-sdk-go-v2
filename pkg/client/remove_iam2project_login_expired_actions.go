// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveIAM2ProjectLoginExpired 操作RemoveIAM2ProjectLoginExpired
func (cli *ZSClient) RemoveIAM2ProjectLoginExpired(uuid string, params param.RemoveIAM2ProjectLoginExpiredParam) (*view.RemoveIAM2ProjectLoginExpiredEventView, error) {
	resp := view.RemoveIAM2ProjectLoginExpiredEventView{}
	if err := cli.Put("v1/iam2/projects/remove/login-expired/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

