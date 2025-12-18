// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateNasMountTarget updates NasMountTarget
func (cli *ZSClient) UpdateNasMountTarget(uuid string, params param.UpdateNasMountTargetParam) (*view.UpdateNasMountTargetEventView, error) {
	resp := view.UpdateNasMountTargetEventView{}
	if err := cli.Put("v1/primary-storage/nas/mount/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
