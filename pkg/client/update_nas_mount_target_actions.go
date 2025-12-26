// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateNasMountTarget updates NasMountTarget
func (cli *ZSClient) UpdateNasMountTarget(uuid string, params param.UpdateNasMountTargetParam) (*view.UpdateNasMountTargetEventView, error) {
	resp := view.UpdateNasMountTargetEventView{}
	if err := cli.Put("v1/primary-storage/nas/mount/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
