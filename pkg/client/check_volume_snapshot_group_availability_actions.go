// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CheckVolumeSnapshotGroupAvailability operates on CheckVolumeSnapshotGroupAvailability
func (cli *ZSClient) CheckVolumeSnapshotGroupAvailability(params param.CheckVolumeSnapshotGroupAvailabilityParam) (*view.CheckVolumeSnapshotGroupAvailabilityView, error) {
	var resp view.CheckVolumeSnapshotGroupAvailabilityView
	if err := cli.Get("v1/volume-snapshots/groups/availabilities", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
