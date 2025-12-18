// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckVolumeSnapshotGroupAvailability 操作CheckVolumeSnapshotGroupAvailability
func (cli *ZSClient) CheckVolumeSnapshotGroupAvailability(params param.CheckVolumeSnapshotGroupAvailabilityParam) (*view.CheckVolumeSnapshotGroupAvailabilityView, error) {
	var resp view.CheckVolumeSnapshotGroupAvailabilityView
	if err := cli.Get("v1/volume-snapshots/groups/availabilities", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

