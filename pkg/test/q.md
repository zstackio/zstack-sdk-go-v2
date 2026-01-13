
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

3. ExpungeXXX 函数统一如下, 原接口时cli.Delete修改为Put
```
func (cli *ZSClient) ExpungeImage(uuid string) error {
	params := map[string]interface{}{
		"expungeImage": jsonutils.NewDict(),
	}
	return cli.Put("v1/images", uuid, jsonutils.Marshal(params), nil)
}
```
NewDict定义如下
```
func NewDict(objs ...JSONPair) *JSONDict {
	dict := JSONDict{data: sortedmap.NewSortedMapWithCapa(len(objs))}
	for _, o := range objs {
		dict.Set(o.key, o.val)
	}
	return &dict
}
```

VmInstance
```
// ExpungeVmInstance Permanently delete a VM instance
func (cli *ZSClient) ExpungeVmInstance(uuid string) error {
	params := map[string]struct{}{
		"expungeVmInstance": {},
	}
	return cli.Put("v1/vm-instances", uuid, params, nil)
}
```
DataVolume
```
func (cli *ZSClient) ExpungeDataVolume(uuid string) error {
	return cli.Put("v1/volumes", uuid, map[string]struct{}{"expungeDataVolume": {}}, nil)
}
```

4. 所有的View 无需使用指针
例如
```
// ImageInventoryView Image
type ImageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description *string `json:"description,omitempty"`
	State *string `json:"state,omitempty"`
	Status *string `json:"status,omitempty"`
	Size *int64 `json:"size,omitempty"`
	ActualSize *int64 `json:"actualSize,omitempty"`
	Md5Sum *string `json:"md5Sum,omitempty"`
	Url *string `json:"url,omitempty"`
	MediaType *string `json:"mediaType,omitempty"`
	GuestOsType *string `json:"guestOsType,omitempty"`
	Type *string `json:"type,omitempty"`
	Platform *string `json:"platform,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	Format *string `json:"format,omitempty"`
	System *bool `json:"system,omitempty"`
	Virtio *bool `json:"virtio,omitempty"`
	BackupStorageRefs []ImageBackupStorageRefInventoryView `json:"backupStorageRefs,omitempty"`
	SystemTags []SystemTagInventoryView `json:"systemTags,omitempty"`
}
```
修改为
```
type ImageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Size int64 `json:"size,omitempty"`
	ActualSize int64 `json:"actualSize,omitempty"`
	Md5Sum string `json:"md5Sum,omitempty"`
	Url string `json:"url,omitempty"`
	MediaType string `json:"mediaType,omitempty"`
	GuestOsType string `json:"guestOsType,omitempty"`
	Type string `json:"type,omitempty"`
	Platform string `json:"platform,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	Format string `json:"format,omitempty"`
	System bool `json:"system,omitempty"`
	Virtio bool `json:"virtio,omitempty"`
	BackupStorageRefs []ImageBackupStorageRefInventoryView `json:"backupStorageRefs,omitempty"`
	SystemTags []SystemTagInventoryView `json:"systemTags,omitempty"`
}
```

5. param包下的UpdateXXXParam的UpdateXXXDetailParam 的json错误，不是params，而是updateXXX
例如
```
type UpdateImageParam struct {
	BaseParam
	Params UpdateImageParamDetail `json:"updateImage"`
}
```
```
type UpdateVmInstanceParam struct {
	BaseParam
	UpdateVmInstance UpdateVmInstanceDetailParam `json:"updateVmInstance"`
}
```

6. ChangeXXXStateParam的ChangeXXXStateDetailParam 的json错误，不是params，而是changeXXXState, 从image查知，可能其他资源类似，需检查
例如
```
type ChangeImageStateParam struct {
	BaseParam
	Params ChangeImageStateParamDetail `json:"changeImageState"`
}
```

7. ChangeXXXState函数的参数uuid不需要，和ChangeXXXStateDetailParam中的Uuid一致。
例如
```
	state, err := accountLoginCli.ChangeImageState("304ff48eeed54f59802152f41a600bfb", param.ChangeImageStateParam{
		BaseParam: param.BaseParam{},
		Params: param.ChangeImageStateParamDetail{
			Uuid:       "304ff48eeed54f59802152f41a600bfb",
			StateEvent: "enable",
		},
	})
```

8. 参数问题可能还有
CloneXXX, 不是params，是cloneXXX
例如
```
type CloneVmInstanceParam struct {
	BaseParam
	CloneVmInstance CloneVmInstanceDetailParam `json:"cloneVmInstance"`
}
```
StartXXX，不是params，是startXXX，StopXXX
例如
```
type StartVmInstanceParam struct {
	BaseParam
	StartVmInstance StartVmInstanceDetailParam `json:"startVmInstance"`
}
```

```
type StopVmInstanceParam struct {
	BaseParam
	StopVmInstance StopVmInstanceDetailParam `json:"stopVmInstance"` // Requires uuid and type
}
```

SetXXX, 不是params, 是setXXX
例如
```
type SetVmBootModeParam struct {
	BaseParam
	SetVmBootMode SetVmBootModeDetailParam `json:"setVmBootMode"`
}
```
```
type SetVmSshKeyParam struct {
	BaseParam
	Params SetVmSshKeyParamDetail `json:"setVmSshKey"`
}
```

ChangeXXX
```
type ChangeVmPasswordParam struct {
	BaseParam
	Params ChangeVmPasswordParamDetail `json:"changeVmPassword"`
}
```