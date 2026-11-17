// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// CheckNetworkReachable operates on NetworkReachable
func (cli *ZSClient) CheckNetworkReachable(params param.CheckNetworkReachableParam) (*view.CheckNetworkReachableView, error) {
	if len(params.Params.TargetHostnames) == 0 {
		return nil, fmt.Errorf("targetHostnames is required")
	}

	query := struct {
		SourceHostnames []string `json:"sourceHostnames,omitempty"`
		TargetHostnames []string `json:"targetHostnames"`
		SystemTags      []string `json:"systemTags,omitempty"`
		UserTags        []string `json:"userTags,omitempty"`
		RequestIp       string   `json:"requestIp,omitempty"`
	}{
		SourceHostnames: params.Params.SourceHostnames,
		TargetHostnames: params.Params.TargetHostnames,
		SystemTags:      params.SystemTags,
		UserTags:        params.UserTags,
		RequestIp:       params.RequestIp,
	}

	var resp view.CheckNetworkReachableView
	if err := cli.GetWithRespKey("v1/zops/check/network", "", "", &query, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
