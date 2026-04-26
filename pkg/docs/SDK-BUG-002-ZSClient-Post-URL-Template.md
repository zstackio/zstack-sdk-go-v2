# SDK-BUG-002 — `ZSClient.Post()` 不解析 URL 路径占位符 `{xxx}`

| 字段 | 值 |
|---|---|
| 严重度 | **High** |
| 状态 | ✅ 已修（2026-04-26，按当前代码实现复核） |
| 发现日期 | 2026-03-26 |
| 影响 SDK 文件 | `pkg/client/client.go`（当前 `ZSClient.Post` shadow） + 历史 action 调用路径 |
| 影响 provider 资源数 | 7 已绕过 + 至少 1 待绕过（`zstack_volume`） |
| 跟踪关键字 | `SDK-BUG-002`、`SDK-WA-002`、`ZSHttpClient.Post(` |

---

## 1. 现象

调用 `*_actions.go` 中的 attach/operate 类方法（URL 含 `{xxx}` 占位符）时，请求 URL 会**直接带字面 `{xxx}` 字符串**，例如：

```
PUT http://zstack:8080/v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}
```

ZStack API 必然返回 404 或 400，资源 attach 失败。

---

## 2. 复核结论（当前代码）

当前代码里，`ZSClient.Post` 已直接委托到底层 `ZSHttpClient.Post`：

```go
func (cli *ZSClient) Post(path string, params interface{}, result interface{}) error {
    return cli.ZSHttpClient.Post(context.Background(), path, params, result)
}
```

底层 `ZSHttpClient.Post` 通过 `getPostURL()` / `getURL()` 走统一 URL 构造路径。与此同时，当前 `pkg/client` 中已没有仍然依赖 `{xxx}` 模板占位符的 action 实现；action 方法使用的是具体资源路径，资源标识通过 `resourceId` 参数进入 URL 拼接逻辑。

因此，**文档里记录的历史问题在当前代码中已不再成立**。

---

## 3. 历史根因（保留归档）

### 2.1 SDK 内有两套 HTTP 客户端

| Client | `Post` 方法位置 | URL 构建方式 |
|---|---|---|
| `ZSHttpClient`（底层） | `pkg/client/http_client.go` `Post(ctx, resource, params, retVal)` | 走 `getPostURL()` → `getURL()`，**正确替换占位符** |
| `ZSClient`（顶层 wrapper） | `pkg/client/client.go:473` 老版 `Post(path, params, result)` + `pkg/client/client.go:31` 新增 shadow `Post(path, params, result)` | 直接 `fmt.Sprintf("%s/%s", baseURL, path)`，**不替换占位符** |

`ZSClient` 嵌入了 `*ZSHttpClient`，但又**自己定义了 `Post()`**（签名不同：`(path, params, result)` vs `(ctx, resource, params, retVal)`）。这导致 Go 方法解析时优先匹配自身的 `Post()`，遮蔽了 `ZSHttpClient.Post()`。

### 2.2 大量 action 方法走错路径

```bash
$ grep -c 'cli\.Post("v1/.*{' pkg/client/other_actions.go
# 历史：321  目前样本：101
```

`other_actions.go` 里的方法全部定义在 `*ZSClient` 上，调用 `cli.Post(...)` 走的是上面的"错误路径"，URL 模板永远不被解析。

### 2.3 client.go 还有重复定义

```bash
$ grep -c "^func (cli \*ZSClient) Post\b" pkg/client/client.go
2   ← 顶部一个 shadow（正确，走 ZSHttpClient），底部一个老版（错误）
```

底部第 473 行的老版 `Post` 是历史遗留；目前 commit `898dccf9` 已经在 `pkg/client/client.go:31` 加了 shadow 版本来 bridge ZSHttpClient，但**老版本仍然存在**且方法签名相同，靠源码顺序是不可靠的（Go 同一个 receiver type 上不能有两个同名方法 — 应该会编译失败；如果编译通过说明只有一个版本生效，需 SDK 团队核实当前状态）。

---

## 4. 历史复现方式

```go
// SDK 测试代码
cli := client.NewZSClient(cfg)
err := cli.AttachL2NetworkToCluster(param.AttachL2NetworkToClusterParam{
    L2NetworkUuid: "abc",
    ClusterUuid:   "def",
})
// 实际请求 URL: /v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}
// 期望请求 URL: /v1/l2-networks/abc/clusters/def
```

---

## 5. 修复结果

当前仓库已不再需要依赖 "删除旧版 `ZSClient.Post`" 或 "对 `{xxx}` 模板做自动替换" 这类补救方案。现状是：

1. `ZSClient.Post` 已桥接到 `ZSHttpClient.Post`
2. 当前 `pkg/client` action 代码不再保留 `{placeholder}` 模板路径调用

建议将此问题视为**历史 bug，当前已修复**。

---

## 6. 历史修复方案（保留归档）

### 方案 A — 删除 `ZSClient.Post()` 老版本（推荐）

只保留 `client.go:31` 的 shadow 版本：

```go
// 删除 pkg/client/client.go:473 起的老版本：
// func (cli *ZSClient) Post(path string, params interface{}, result interface{}) error {
//     url := fmt.Sprintf("%s/%s", cli.baseURL(), path)
//     return cli.doRequest("POST", url, params, result)
// }
```

shadow 版本会直接走 `ZSHttpClient.Post(ctx, ...)`，后者已经正确解析 URL。

同样需要核查并删除老版的 `Get` / `Put` / `Delete`（client.go:404, 479, 485）。

### 方案 B — 给 `ZSClient.Post()` 接管 URL 模板替换

如果保留老版本是有原因的（不带 ctx 的 API），让它先做模板替换再发出：

```go
func (cli *ZSClient) Post(path string, params interface{}, result interface{}) error {
    resolved := resolveURLTemplate(path, params) // 新加：从 params 提取 *Uuid 字段填充 {xxx}
    return cli.ZSHttpClient.Post(context.Background(), resolved, params, result)
}
```

复杂度高，不如方案 A 干净。

### 方案 C — 使所有 `_actions.go` 方法显式调用 `ZSHttpClient`

```diff
-cli.Post("v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}", params, &resp)
+cli.ZSHttpClient.Post(ctx, "v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}", params, &resp)
```

需要修 101 处。建议在代码生成器里改一次性扫修。

---

## 7. 已确认 provider 已绕过的资源

> Provider 直接走 `r.client.ZSHttpClient.Post(ctx, fmt.Sprintf("v1/...", uuid), params, &resp)` 手动拼接 URL。

| 资源 | 文件 | Provider 函数 | 绕过的 SDK 方法 | URL 模板 |
|---|---|---|---|---|
| `zstack_l2vlan_network` | `resource_zstack_l2vlan_network.go` | `attachCluster()` | `AttachL2NetworkToCluster` | `v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}` |
| `zstack_port_forwarding_rule` | `resource_zstack_port_forwarding_rule.go` | `attachToVmNic()` | `AttachPortForwardingRule` | `v1/port-forwarding/{ruleUuid}/vm-instances/nics/{vmNicUuid}` |
| `zstack_load_balancer_listener` | `resource_zstack_load_balancer_listener.go` | `Create()` | `CreateLoadBalancerListener` | `v1/load-balancers/{loadBalancerUuid}/listeners` |
| `zstack_primary_storage` (Local) | `resource_zstack_primary_storage.go` | `Create()` | `AddLocalPrimaryStorage` | `v1/primary-storage/local-storage` |
| `zstack_primary_storage` (NFS) | `resource_zstack_primary_storage.go` | `Create()` | `AddNfsPrimaryStorage` | `v1/primary-storage/nfs` |
| `zstack_primary_storage` | `resource_zstack_primary_storage.go` | `attachCluster()` | `AttachPrimaryStorageToCluster` | `v1/clusters/{clusterUuid}/primary-storage/{primaryStorageUuid}` |
| `zstack_backup_storage` | `resource_zstack_backup_storage.go` | `attachZone()` | `AttachBackupStorageToZone` | `v1/zones/{zoneUuid}/backup-storage/{backupStorageUuid}` |

### 待绕过（同症状）

| 资源 | 文件 | 受影响方法 |
|---|---|---|
| `zstack_volume` | `resource_zstack_volume.go` 第 208/356 行 | `AttachDataVolumeToVm` (`v1/volumes/{volumeUuid}/vm-instances/{vmInstanceUuid}`) |

### 后续新资源会撞上的方法

- `AddBackendServerToServerGroup`（load_balancer 增强）
- `AddSchedulerJobToTrigger`（scheduler 增强）

---

## 8. SDK 修复后的 Provider 回收清单

### 当前状态

- 当前 SDK 代码路径已确认不再保留该 bug。
- 下游 provider 可优先验证是否仍需保留历史 workaround。
- 若 provider 侧升级到当前 SDK 后验证通过，可逐步回收 `ZSHttpClient.Post(...)` 的手工绕过代码。

```bash
cd terraform-provider-zstack

# 1. 升级 SDK
go get github.com/zstackio/zstack-sdk-go-v2@latest

# 2. 找回收点
grep -rn "ZSHttpClient.Post(" zstack/provider/

# 3. 改回标准 SDK 方法调用
# 把:
#   err := r.client.ZSHttpClient.Post(
#       ctx,
#       fmt.Sprintf("v1/l2-networks/%s/clusters/%s", l2NetUuid, clusterUuid),
#       params,
#       &resp,
#   )
# 改回:
#   resp, err := r.client.AttachL2NetworkToCluster(param.AttachL2NetworkToClusterParam{
#       L2NetworkUuid: l2NetUuid,
#       ClusterUuid:   clusterUuid,
#   })

# 4. 删 docs/SDK_URL_TEMPLATE_BUG.md 中"按资源 Workaround 记录"的对应行
# 5. 删除每个绕过点旁边的 // SDK bug: 注释
```

如果 SDK 同时把 `Add*PrimaryStorage` 类方法的入参签名补全（目前 SDK 没参数），还可以把 primary_storage 的 Local / NFS Create 也回收。

---

## 9. 原始报告

- `terraform-provider-zstack/docs/SDK_URL_TEMPLATE_BUG.md` — 中文完整调查
- `terraform-provider-zstack/docs/sdk-url-template-bug.md` — 同根（异名文件）
- `terraform-provider-zstack/_bmad-output/bug-tracker.md` § "SDK-WA-002"
