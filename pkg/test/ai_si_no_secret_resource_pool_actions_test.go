// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"
)

func TestCreateAiSiNoSecretResourcePool(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateAiSiNoSecretResourcePool is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateAiSiNoSecretResourcePoolParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateAiSiNoSecretResourcePoolParamDetail{
	// 		Name: "test-aisinosecretresourcepool",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateAiSiNoSecretResourcePool(context.Background(), createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateAiSiNoSecretResourcePool error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateAiSiNoSecretResourcePool result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteAiSiNoSecretResourcePool(context.Background(), result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteAiSiNoSecretResourcePool error: %v", err)
	// }
}
