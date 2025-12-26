// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetEipAttachableVmNics gets EipAttachableVmNics by uuid
func (cli *ZSClient) GetEipAttachableVmNics(uuid string) (*view.GetEipAttachableVmNicsView, error) {
	var resp view.GetEipAttachableVmNicsView
	if err := cli.Get("v1/eips/{eipUuid}/vm-instances/candidate-nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
