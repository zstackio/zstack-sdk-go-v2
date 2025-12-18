// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetObservabilityServerServiceData 获取ObservabilityServerServiceData详情
func (cli *ZSClient) GetObservabilityServerServiceData(uuid string) (*view.GetObservabilityServerServiceDataView, error) {
	var resp view.GetObservabilityServerServiceDataView
	if err := cli.Get("v1/observability-server/{observabilityServerUuid}/service-data", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

