// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryModel queries Model list
func (cli *ZSClient) QueryModel(params *param.QueryParam) ([]view.ModelInventoryView, error) {
	var resp []view.ModelInventoryView
	return resp, cli.List("v1/ai/models", params, &resp)
}

func (cli *ZSClient) GetModel(uuid string) (*view.ModelInventoryView, error) {
	var resp view.ModelInventoryView
	if err := cli.Get("v1/ai/models", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageModel Pagination
func (cli *ZSClient) PageModel(params *param.QueryParam) ([]view.ModelInventoryView, int, error) {
	var models []view.ModelInventoryView
	total, err := cli.Page("v1/ai/models", params, &models)
	return models, total, err
}
// DeleteModel deletes Model
func (cli *ZSClient) DeleteModel(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ai/models", uuid, string(deleteMode))
}
// UpdateModel updates Model
func (cli *ZSClient) UpdateModel(uuid string, params param.UpdateModelParam) (*view.ModelInventoryView, error) {
	resp := view.ModelInventoryView{}
	if err := cli.PutWithRespKey("v1/ai/models", uuid, "", map[string]interface{}{
		"updateModel": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddModel adds Model
func (cli *ZSClient) AddModel(params param.AddModelParam) (*view.ModelInventoryView, error) {
	resp := view.ModelInventoryView{}
	if err := cli.Post("v1/ai/models", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddModelAsync Async
func (cli *ZSClient) AddModelAsync(params param.AddModelParam) (string, error) {

	resource := "v1/ai/models"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// GetSupportedIdentityModels gets SupportedIdentityModels by uuid
func (cli *ZSClient) GetSupportedIdentityModels() (*view.GetSupportedIdentityModelsView, error) {
	var resp view.GetSupportedIdentityModelsView
	if err := cli.GetWithRespKey("v1/identity-models", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteModels sends one UUID to the collection endpoint.
func (cli *ZSClient) DeleteModels(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteModelsWithParams(deleteMode, param.DeleteModelsParamDetail{Uuids: []string{uuid}})
}

// DeleteModelsWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DeleteModelsWithParams(deleteMode param.DeleteMode, params param.DeleteModelsParamDetail) error {
	query := url.Values{
		"uuids": params.Uuids,
	}
	return cli.deleteActionWithQuery("v1/ai/models", "", "", deleteModeOrDefault(deleteMode, params.DeleteMode), query, "uuids")
}

// GenerateModelMetadata operates on ModelMetadata
func (cli *ZSClient) GenerateModelMetadata(params param.GenerateModelMetadataParam) (*view.GenerateModelMetadataEventView, error) {
	resp := view.GenerateModelMetadataEventView{}
	if err := cli.PutWithRespKey("v1/ai/model/metadata/generate", "", "", map[string]interface{}{
		"generateModelMetadata": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BindModelToService operates on BindModelToService
func (cli *ZSClient) BindModelToService(modelUuid string, modelServiceUuid string, params param.BindModelToServiceParam) (*view.ModelServiceInventoryView, error) {
	resp := view.ModelServiceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/ai/models/%s/model-services/%s", modelUuid, modelServiceUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnbindModelFromService operates on UnbindModelFromService
func (cli *ZSClient) UnbindModelFromService(modelUuid string, modelServiceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/ai/models", modelUuid, fmt.Sprintf("model-services/%s", modelServiceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
