// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachIscsiServerToCluster 操作IscsiServerToCluster
func (cli *ZSClient) AttachIscsiServerToCluster(params param.AttachIscsiServerToClusterParam) (*view.AttachIscsiServerToClusterEventView, error) {
	resp := view.AttachIscsiServerToClusterEventView{}
	if err := cli.Post("v1/clusters/{clusterUuid}/storage-devices/iscsi/servers/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

