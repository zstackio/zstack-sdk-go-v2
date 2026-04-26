# zstack-sdk-go-v2 Bug Tracker

> 本目录登记 `zstack-sdk-go-v2` 中由 `terraform-provider-zstack` 在生产中遇到、并已在 provider 侧绕过的 SDK 缺陷。
>
> **目的**：让 SDK 团队按编号逐个修复；修好后下游 provider 可以根据每篇文档底部的"回收清单"删除 workaround。
>
> **生成日期**：2026-04-26
> **下游 provider 版本**：`terraform-provider-zstack` @ `fix/qa-20260422-p0-plus-story15`
> **SDK 版本**：`github.com/zstackio/zstack-sdk-go-v2 v0.0.4`（commit `898dccf9` 基线）

---

## Bug 索引

| 编号 | 标题 | 严重度 | 影响范围 | Provider workaround | 状态 |
|---|---|---|---|---|---|
| [SDK-BUG-001](./SDK-BUG-001-PutWithRespKey-Empty-Envelope.md) | `PutWithRespKey` / `PutWithSpec` 第 3 参数 `responseKey=""` 导致响应 envelope 解析失败 | **High** | 569 处调用，至少 23 个 provider 资源已受影响 | re-query (`Query*` / `Get*`) 替换 SDK 返回值 | ✅ 已修 |
| [SDK-BUG-002](./SDK-BUG-002-ZSClient-Post-URL-Template.md) | `ZSClient.Post()` 不解析 URL 路径占位符 `{xxx}` | **High** | 历史上影响 101+ 个 action 方法 | 直接走 `ZSHttpClient.Post()` 拼接 URL | ✅ 已修 |
| [SDK-BUG-003](./SDK-BUG-003-IAM2Project-Soft-Delete.md) | `DeleteIAM2Project` 是软删除，不释放 name | Medium | 1 个 SDK 方法 | provider Delete 后追加 `ExpungeIAM2Project` | ✅ 已修 |
| [SDK-BUG-004](./SDK-BUG-004-L3Network-Delete-URL.md) | `DeleteL3Network` URL 拼接异常 | High | 历史上影响 1 个 SDK 方法 | 暂未绕过；`zstack_l3network` Delete 报错 | ✅ 已修 |

---

## 优先级建议

| 优先级 | Bug |
|---|---|
| **P0** — 影响面最大 | SDK-BUG-001（修一处 SDK 底层 → 23 处 provider 自动受益） |
| **P0** — 已完成 | SDK-BUG-004（当前代码路径已确认会携带 UUID） |
| **P1** — 已完成 | SDK-BUG-002（当前代码路径不再依赖模板占位符解析） |
| **P2** — 单点功能受限 | SDK-BUG-003 |

---

## SDK 修复 → Provider 回收联动流程

每篇 bug 文档底部都有 **"SDK 修复后的回收清单"** 一节，列出 provider 侧可以删除的代码位置。建议工作流：

1. SDK 团队按本目录发版（`v0.0.5+`）
2. CHANGELOG 注明修复了哪些 `SDK-BUG-NNN`
3. 提示下游 provider：
   - 升级 SDK 依赖（`go get github.com/zstackio/zstack-sdk-go-v2@latest`）
   - 在 provider 仓库搜索关键字 `SDK-BUG-NNN`、`SDK-WA-NNN`，按文档清单删除 workaround
   - 跑 acceptance test 验证

---

## 相关原始文档（provider 仓库）

- `troubleshooting/SDK-BUG-UpdateAlarm-Empty-Response.md` — SDK-BUG-001 的发现报告
- `docs/SDK_URL_TEMPLATE_BUG.md` — SDK-BUG-002 的发现报告
- `_bmad-output/bug-tracker.md` § "SDK Workaround Registry" — 全量 workaround 索引
