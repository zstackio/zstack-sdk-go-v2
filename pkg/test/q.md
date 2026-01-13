
account_resource_ref_actions.go

```
好像api没有Get
func (cli *ZSClient) GetAccountResourceRef(uuid string) (*view.AccountResourceRefInventoryView, error) {
	var resp view.AccountResourceRefInventoryView
	if err := cli.Get("v1/accounts/resources/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

```

global_config_actions.go 
```
GET zstack/v1/resource-configurations/{resourceUuid}/{category}/{name}
func (cli *ZSClient) GetGlobalConfig(uuid string) (*view.GlobalConfigInventoryView, error) {
	var resp view.GlobalConfigInventoryView
	if err := cli.Get("v1/global-configurations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
```

在http_client.go中

1. 定义了Page相关函数，在资源对象中要增加PageXXX函数。例如镜像
```
func (cli *ZSHttpClient) Page(resource string, params *param.QueryParam, retVal interface{}) (int, error) {
	return cli.PageWithKey(resource, responseKeyInventories, params, retVal)
}
```
要增加PageImage函数
```
// PageImage Pagination
func (cli *ZSClient) PageImage(params param.QueryParam) ([]view.ImageView, int, error) {
	var images []view.ImageView
	total, err := cli.Page("v1/images", &params, &images)
	return images, total, err
}
```

2. http_client 在处理 Post, Put 请求时会自动解包 inventory 字段, 因此资源对象的Post, Put函数中不需要再处理inventory字段。例如AddImage

```
	responseKeyInventories = "inventories"
	responseKeyInventory   = "inventory"
```

```
func (cli *ZSHttpClient) Post(resource string, params interface{}, retVal interface{}) error {
	return cli.PostWithRespKey(resource, responseKeyInventory, params, retVal)
}
```

```
func (cli *ZSHttpClient) Put(resource, resourceId string, params interface{}, retVal interface{}) error {
	return cli.PutWithRespKey(resource, resourceId, responseKeyInventory, params, retVal)
}
```

例如AddImage, 无需返回&resp.Inventory， 而是直接返回ImageInventoryView
```
func (cli *ZSClient) AddImage(params param.AddImageParam) (*view.ImageInventoryView, error) {
	var resp view.ImageInventoryView
	if err := cli.Post("v1/images", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
```
