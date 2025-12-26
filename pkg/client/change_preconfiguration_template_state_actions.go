// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangePreconfigurationTemplateState changes PreconfigurationTemplateState
func (cli *ZSClient) ChangePreconfigurationTemplateState(uuid string, params param.ChangePreconfigurationTemplateStateParam) (*view.ChangePreconfigurationTemplateStateEventView, error) {
	resp := view.ChangePreconfigurationTemplateStateEventView{}
	if err := cli.Put("v1/baremetal/preconfigurations/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
