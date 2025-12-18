// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeAffinityGroupState 操作AffinityGroupState
func (cli *ZSClient) ChangeAffinityGroupState(uuid string, params param.ChangeAffinityGroupStateParam) (*view.ChangeAffinityGroupStateEventView, error) {
	resp := view.ChangeAffinityGroupStateEventView{}
	if err := cli.Put("v1/affinity-groups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

