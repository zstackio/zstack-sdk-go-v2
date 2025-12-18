// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteMetricData deletes MetricData
func (cli *ZSClient) DeleteMetricData(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/metrics", uuid, string(deleteMode))
}
