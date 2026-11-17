// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// CreateAliyunNasAccessGroupRule creates AliyunNasAccessGroupRule
func (cli *ZSClient) CreateAliyunNasAccessGroupRule(params param.CreateAliyunNasAccessGroupRuleParam) (*view.AliyunNasAccessRuleInventoryView, error) {
	resp := view.AliyunNasAccessRuleInventoryView{}
	if err := cli.Post("v1/nas/aliyun/rule", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAliyunNasAccessGroupRule deletes AliyunNasAccessGroupRule
func (cli *ZSClient) DeleteAliyunNasAccessGroupRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/nas/access/rule", uuid, string(deleteMode))
}
