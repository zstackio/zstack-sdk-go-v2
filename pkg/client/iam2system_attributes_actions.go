// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetIAM2SystemAttributes 获取IAM2SystemAttributes详情
func (cli *ZSClient) GetIAM2SystemAttributes(uuid string) (*view.GetIAM2SystemAttributesView, error) {
	var resp view.GetIAM2SystemAttributesView
	if err := cli.Get("v1/iam2/attributes/system", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

