// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ApplyRuleSetChanges 操作ApplyRuleSetChanges
func (cli *ZSClient) ApplyRuleSetChanges(uuid string, params param.ApplyRuleSetChangesParam) (*view.ApplyRuleSetChangesEventView, error) {
	resp := view.ApplyRuleSetChangesEventView{}
	if err := cli.Put("v1/vpcfirewalls/ruleSets/apply/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

