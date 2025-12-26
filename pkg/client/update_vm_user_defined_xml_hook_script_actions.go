// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVmUserDefinedXmlHookScript updates VmUserDefinedXmlHookScript
func (cli *ZSClient) UpdateVmUserDefinedXmlHookScript(uuid string, params param.UpdateVmUserDefinedXmlHookScriptParam) (*view.UpdateVmUserDefinedXmlHookScriptEventView, error) {
	resp := view.UpdateVmUserDefinedXmlHookScriptEventView{}
	if err := cli.Put("v1/vm-instances/xml-hook-script", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
