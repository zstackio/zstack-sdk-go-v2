package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// GetEffectiveResourceConfig returns the value selected by Cloud resource-config
// inheritance for a resource, including its global fallback.
func (cli *ZSClient) GetEffectiveResourceConfig(resourceUuid, category, name string) (*view.GetResourceConfigView, error) {
	var resp view.GetResourceConfigView
	if err := cli.GetWithSpec("v1/resource-configurations", resourceUuid, fmt.Sprintf("%s/%s", category, name), "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceBindableConfig gets ResourceBindableConfig by uuid
func (cli *ZSClient) GetResourceBindableConfig() (*view.GetResourceBindableConfigView, error) {
	var resp view.GetResourceBindableConfigView
	if err := cli.GetWithRespKey("v1/resource-configurations/bindable", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
