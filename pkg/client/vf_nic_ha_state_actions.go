// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeVfNicHaState 操作VfNicHaState
func (cli *ZSClient) ChangeVfNicHaState(uuid string, params param.ChangeVfNicHaStateParam) (*view.ChangeVfNicHaStateEventView, error) {
	resp := view.ChangeVfNicHaStateEventView{}
	if err := cli.Put("v1/vm-instances/nics/{vfNicUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

