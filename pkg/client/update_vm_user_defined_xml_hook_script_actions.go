// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateVmUserDefinedXmlHookScript updates VmUserDefinedXmlHookScript
func (cli *ZSClient) UpdateVmUserDefinedXmlHookScript(uuid string, params param.UpdateVmUserDefinedXmlHookScriptParam) (*view.UpdateVmUserDefinedXmlHookScriptEventView, error) {
	resp := view.UpdateVmUserDefinedXmlHookScriptEventView{}
	if err := cli.Put("v1/vm-instances/xml-hook-script", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
