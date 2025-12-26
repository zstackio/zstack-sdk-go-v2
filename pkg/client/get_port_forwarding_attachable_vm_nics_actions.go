// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetPortForwardingAttachableVmNics gets PortForwardingAttachableVmNics by uuid
func (cli *ZSClient) GetPortForwardingAttachableVmNics(uuid string) (*view.GetPortForwardingAttachableVmNicsView, error) {
	var resp view.GetPortForwardingAttachableVmNicsView
	if err := cli.Get("v1/port-forwarding/{ruleUuid}/vm-instances/candidate-nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
