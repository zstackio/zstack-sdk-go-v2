// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeVfNicHaState changes VfNicHaState
func (cli *ZSClient) ChangeVfNicHaState(uuid string, params param.ChangeVfNicHaStateParam) (*view.ChangeVfNicHaStateEventView, error) {
	resp := view.ChangeVfNicHaStateEventView{}
	if err := cli.Put("v1/vm-instances/nics/{vfNicUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
