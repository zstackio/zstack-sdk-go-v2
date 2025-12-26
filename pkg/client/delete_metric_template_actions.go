// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteMetricTemplate deletes MetricTemplate
func (cli *ZSClient) DeleteMetricTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/metrics/httpreceivers/templates/{uuid}", uuid, string(deleteMode))
}
