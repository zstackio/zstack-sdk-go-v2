// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ApplyRuleSetChanges operates on ApplyRuleSetChanges
func (cli *ZSClient) ApplyRuleSetChanges(uuid string, params param.ApplyRuleSetChangesParam) (*view.ApplyRuleSetChangesEventView, error) {
	resp := view.ApplyRuleSetChangesEventView{}
	if err := cli.Put("v1/vpcfirewalls/ruleSets/apply/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
