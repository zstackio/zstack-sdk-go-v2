// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangePreconfigurationTemplateState 操作PreconfigurationTemplateState
func (cli *ZSClient) ChangePreconfigurationTemplateState(uuid string, params param.ChangePreconfigurationTemplateStateParam) (*view.ChangePreconfigurationTemplateStateEventView, error) {
	resp := view.ChangePreconfigurationTemplateStateEventView{}
	if err := cli.Put("v1/baremetal/preconfigurations/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

