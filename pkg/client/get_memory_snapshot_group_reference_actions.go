// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetMemorySnapshotGroupReference gets MemorySnapshotGroupReference by uuid
func (cli *ZSClient) GetMemorySnapshotGroupReference(uuid string) (*view.GetMemorySnapshotGroupReferenceView, error) {
	var resp view.GetMemorySnapshotGroupReferenceView
	if err := cli.Get("v1/memory-snapshots/group/reference", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
