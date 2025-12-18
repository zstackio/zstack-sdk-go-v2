// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidatePrimaryStoragesForCreatingVm 获取CandidatePrimaryStoragesForCreatingVm详情
func (cli *ZSClient) GetCandidatePrimaryStoragesForCreatingVm(uuid string) (*view.GetCandidatePrimaryStoragesForCreatingVmView, error) {
	var resp view.GetCandidatePrimaryStoragesForCreatingVmView
	if err := cli.Get("v1/vm-instances/candidate-storages", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

