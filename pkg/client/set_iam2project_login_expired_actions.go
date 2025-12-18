// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetIAM2ProjectLoginExpired 操作SetIAM2ProjectLoginExpired
func (cli *ZSClient) SetIAM2ProjectLoginExpired(uuid string, params param.SetIAM2ProjectLoginExpiredParam) (*view.SetIAM2ProjectLoginExpiredEventView, error) {
	resp := view.SetIAM2ProjectLoginExpiredEventView{}
	if err := cli.Put("v1/iam2/projects/add/login-expired/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

