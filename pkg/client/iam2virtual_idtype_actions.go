// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeIAM2VirtualIDType 操作IAM2VirtualIDType
func (cli *ZSClient) ChangeIAM2VirtualIDType(uuid string, params param.ChangeIAM2VirtualIDTypeParam) (*view.ChangeIAM2VirtualIDTypeEventView, error) {
	resp := view.ChangeIAM2VirtualIDTypeEventView{}
	if err := cli.Put("v1/iam2/virtual-ids/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

