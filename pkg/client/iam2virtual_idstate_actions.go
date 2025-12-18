// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeIAM2VirtualIDState 操作IAM2VirtualIDState
func (cli *ZSClient) ChangeIAM2VirtualIDState(uuid string, params param.ChangeIAM2VirtualIDStateParam) (*view.ChangeIAM2VirtualIDStateEventView, error) {
	resp := view.ChangeIAM2VirtualIDStateEventView{}
	if err := cli.Put("v1/iam2/virtual-ids/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

