// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetImageCandidatesForVmToChange gets ImageCandidatesForVmToChange by uuid
func (cli *ZSClient) GetImageCandidatesForVmToChange(uuid string) (*view.GetImageCandidatesForVmToChangeView, error) {
	var resp view.GetImageCandidatesForVmToChangeView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/image-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
