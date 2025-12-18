// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAliyunNasAccessGroupRule 创建AliyunNasAccessGroupRule
func (cli *ZSClient) CreateAliyunNasAccessGroupRule(params param.CreateAliyunNasAccessGroupRuleParam) (*view.CreateAliyunNasAccessGroupRuleEventView, error) {
	resp := view.CreateAliyunNasAccessGroupRuleEventView{}
	if err := cli.Post("v1/nas/aliyun/rule", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

