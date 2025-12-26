// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAliyunNasAccessGroupRule creates AliyunNasAccessGroupRule
func (cli *ZSClient) CreateAliyunNasAccessGroupRule(params param.CreateAliyunNasAccessGroupRuleParam) (*view.CreateAliyunNasAccessGroupRuleEventView, error) {
	resp := view.CreateAliyunNasAccessGroupRuleEventView{}
	if err := cli.Post("v1/nas/aliyun/rule", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
