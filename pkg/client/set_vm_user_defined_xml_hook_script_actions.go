// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmUserDefinedXmlHookScript operates on SetVmUserDefinedXmlHookScript
func (cli *ZSClient) SetVmUserDefinedXmlHookScript(uuid string, params param.SetVmUserDefinedXmlHookScriptParam) (*view.SetVmUserDefinedXmlHookScriptEventView, error) {
	resp := view.SetVmUserDefinedXmlHookScriptEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
