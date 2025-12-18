// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVmUserDefinedXmlHookScript 操作SetVmUserDefinedXmlHookScript
func (cli *ZSClient) SetVmUserDefinedXmlHookScript(uuid string, params param.SetVmUserDefinedXmlHookScriptParam) (*view.SetVmUserDefinedXmlHookScriptEventView, error) {
	resp := view.SetVmUserDefinedXmlHookScriptEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

