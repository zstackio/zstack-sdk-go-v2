// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmXml gets VmXml by uuid
func (cli *ZSClient) GetVmXml(uuid string) (*view.GetVmXmlView, error) {
	var resp view.GetVmXmlView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/xml", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
