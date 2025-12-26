// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmXmlHookScript gets VmXmlHookScript by uuid
func (cli *ZSClient) GetVmXmlHookScript(uuid string) (*view.GetVmXmlHookScriptView, error) {
	var resp view.GetVmXmlHookScriptView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/xml-hook-script", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
