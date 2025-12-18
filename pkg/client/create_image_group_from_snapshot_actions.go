// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateImageGroupFromSnapshot creates ImageGroupFromSnapshot
func (cli *ZSClient) CreateImageGroupFromSnapshot(params param.CreateImageGroupFromSnapshotParam) (*view.CreateImageGroupFromSnapshotEventView, error) {
	resp := view.CreateImageGroupFromSnapshotEventView{}
	if err := cli.Post("v1/imagegroup/from/snapshot/{rootVolumeSnapshotUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
