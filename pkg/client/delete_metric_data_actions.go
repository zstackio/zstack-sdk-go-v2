// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteMetricData deletes MetricData
func (cli *ZSClient) DeleteMetricData(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/metrics", uuid, string(deleteMode))
}
