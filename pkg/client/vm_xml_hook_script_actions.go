// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmXmlHookScript 获取VmXmlHookScript详情
func (cli *ZSClient) GetVmXmlHookScript(uuid string) (*view.GetVmXmlHookScriptView, error) {
	var resp view.GetVmXmlHookScriptView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/xml-hook-script", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

