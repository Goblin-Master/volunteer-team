# Refactor 提案（基于原始版本）

## 项目概览
- 技术栈：Vue 3（Composition API）+ TypeScript + Vite
- 依赖：vue-router、pinia、axios、element-plus、@element-plus/icons-vue
- 结构：`src/router` 路由与守卫，`src/stores/user` 用户状态，`src/api/*` 网络层，`src/views/*` 页面，`src/types/*` 类型
- 响应结构：统一 `BaseResp<T> = { code, message, data }`

## 优点
- 基础框架清晰：路由守卫 + Pinia + Axios 拦截器，能完成登录态与权限的最小闭环
- 类型有雏形：接口层返回 `BaseResp<T>`，常见页面表单具备基本的 TypeScript 注解
- 代码组织上手快：API 与页面按功能分包，易于协作与定位
- UI 使用 Element Plus：表单与交互一致性较好，自动引入降低模板样板代码

## 主要问题与风险
- 鉴权与 JWT
  - `parse_token` 直接 `atob` 解析，无签名校验与健壮过期处理；异常 token 易导致登录流程报错
  - `is_user_login` 仅以本地 `token` 是否存在判断，不校验有效性/过期/401 处理
  - 角色与权限：角色枚举与后端 `role_id` 映射缺少统一约定，权限判定分散在路由与页面
- 网络层与错误处理
  - 接口路径与 `baseURL` 硬编码；缺少环境变量与集中常量管理
  - 响应错误仅 `console.error`，无统一用户提示策略、401 自动登出/跳转登录
  - 未兼容文档说明的 401“纯字符串响应体”场景
- 类型与可维护性
  - 请求参数常用 `any`，DTO 未统一；校验器 `rule: any` 等影响类型安全
  - 缺少公共 `types` 复用与 API 字段规范（如 `int64` 用字符串序列化、`ctime/utime` 时间戳）
- 代码质量与测试
  - 无 ESLint 检查与规则，`prettier` 仅格式化
  - 无单元/集成测试（路由、鉴权、网络层均不可验证），修改风险高
- 样式与复用
  - 页面样式重复，缺少通用布局/卡片/按钮等基础组件抽取
  - 主题与变量未统一（仅少量 `--el-color-primary` 引用）

## 改进目标
- 提升稳定性：JWT 解析与错误容错、401 统一处理
- 强化可维护性：接口规范与类型统一、环境变量与常量集中化
- 优化用户体验：一致的错误提示与鉴权跳转
- 降低风险：引入代码质量工具与最小测试覆盖

## 优先级改进清单（按迭代顺序）
1) 基础治理（P0）
- 配置集中化：新增 `src/config/env.ts` 读取 `import.meta.env`，集中 `BASE_URL`、`UPLOAD_URL`、超时等常量
- 接口常量：新增 `src/config/api.ts` 维护路由前缀与资源路径，消除魔法字符串
- ESLint：加入 `eslint + @typescript-eslint`，规则覆盖空值、any、未使用变量、import 顺序等；脚本 `npm run lint`

2) 鉴权与状态（P0）
- JWT 容错：解析失败不崩溃，缺失 `exp` 视为未过期；从后端 `data` 兜底 `username/avatar/role`
- 401 统一处理：响应拦截器识别纯字符串体，统一清理状态并跳转登录（保留 `redirect`）
- 权限中间件：基于路由 `meta.roles` 添加角色检查（`COMMON_USER/INTERNAL_USER/ADMIN`）

3) 网络层（P0）
- `ApiService` 统一封装：请求/响应拦截器、错误分类（网络/业务/401）、重试选项与取消支持
- 文档对齐：严格遵循 API.md
  - 所有 `int64` ID 用字符串序列化（显示/提交均按字符串）
  - 时间字段 `ctime/utime` 按整型时间戳处理
- 环境切换：`baseURL` 来自 `env`，开发/测试/生产可切换

4) 类型与接口（P1）
- 请求/响应 DTO：在 `src/types/*` 统一定义并复用，消除 `any`
- 校验器类型：抽出常用 `FormRules` 与校验函数类型，减少 `any`

5) 组件与样式（P1）
- 抽取通用组件：`BaseCard / BaseButton / EmptyState / Loading` 等，减少页面重复
- 主题与变量：建立 `src/styles/variables.css`/`scss`，统一主题色与间距

6) 路由结构（P1）
- `meta` 细化：`requiresAuth` + `roles`，在守卫中统一判定
- 登录后跳转：保留 `redirect`，避免多处硬编码跳转

7) 错误与提示（P1）
- 统一错误提示：业务错误读 `message`，网络错误给出“网络异常/请重试”，401 显示“登录过期”并跳转
- 日志：在拦截器中记录请求耗时与关键错误（限制敏感信息）

8) 测试与质量（P2）
- 最小单测：`vitest` 覆盖 `parseToken`、`hasPermission`、路由守卫与 ApiService 错误分支
- 集成测试（可选）：`cypress` 覆盖登录流程与受限路由访问

## 与 API.md 的对齐要点（必须项）
- 所有 ID 字段（`user_id/order_id/summary_id`）以字符串传递与展示（雪花 `int64` → string）
- 时间戳字段 `ctime/utime` 按整数处理（秒级）
- 401 可能返回纯字符串体，前端需要兼容（拦截器单独分支）
- 静态资源展示：拼接 `{scheme}://{host}:{port}/uploads/{fileName}`

## 后端 JWT 与 ID 规范（新增）
- 角色枚举（后端）：`COMMON_USER=1`、`INTERNAL_USER=2`、`ADMIN=3`
- JWT 载荷（后端生成）：
  - `user_id: int64`（雪花 ID）
  - `role_id: Role`
  - `exp/iAT` 由 `jwt.RegisteredClaims` 提供
- 前端处理：
  - `Authorization: Bearer <token>` 传递
  - 解析 token 读取 `role_id` 与 `user_id`（注意键名与大小写）
  - 兼容 401 响应体为纯字符串（如 `"token已过期"`）的场景

### ID 一致性与类型（强制）
- 雪花 ID（`int64`）在返回中一律以字符串序列化；前端所有 ID 类型统一为 `string`
  - 字段示例：`user_id`、`order_id`、`summary_id`
  - 前端提交参数也必须使用字符串（避免 JS 精度问题）
- 时间字段：`ctime/utime` 为整型时间戳（常见为秒）；前端按整数处理并本地格式化展示

### 统一响应与错误
- 统一包：`{ code: 0|-1, message: string, data: object|null }`
- 权限不足：HTTP 200 + `code=-1`，`message="权限不足"`
- token 为空/过期/非法：HTTP 401 + 纯字符串响应体（非统一 JSON 包）

## 前端落地规范（颗粒度收敛）
- 类型与 DTO：
  - `type ID = string`；接口 DTO 中所有 ID 字段使用 `string`
  - 将 `order_id/summary_id/user_id`、`ctime/utime` 明确为必选/可选并复用公共类型
- 网络层与拦截器：
  - 识别 401 纯字符串：清理登录态、跳转登录、保留 `redirect`
  - 统一错误提示：业务错误读 `message`；网络错误提示“网络异常或服务器无响应”
- 权限与路由：
  - 路由 `meta.roles` 使用后端枚举值（1/2/3）进行判定
  - 登录态判断不以 token 是否存在为准；加入过期与 401 处理
- 静态资源：
  - 头像等文件展示统一拼接 `/uploads/<fileName>`，基于环境变量构造完整 URL

## 接口字段快览（按 API.md 与后端约定）
- 登录 `POST /api/user/login`
  - 请求：`account/password` 或 `email/code`；`login_type: "account"|"email"`
  - 响应 `data`：`token: string`、`username: string`、`avatar: string(fileName)`
- 订单列表 `GET /api/order/list`
  - 响应 `data.orders[]`：`order_id: ID`、`ctime: int`、`problem_description: string`
- 订单详情 `GET /api/order/detail?order_id: ID`
  - 响应 `data`：`username, student_id, campus_location, department, phone_number, wechat_id, address, device_model, os_version, problem_description, notes, ctime`
- 修机总结创建 `POST /api/summary/create`
  - 请求：`order_id: ID`、`problem_type`、`problem_description`、`repair_summary`、`receiver_name`
- 修机总结列表 `GET /api/summary/list`
  - 响应 `data.summaries[]`：`summary_id: ID`、`order_id: ID`、`problem_type`、`problem_description`、`repair_summary`、`receiver_name`、`utime: int`
- 修机总结详情 `GET /api/summary/detail?summary_id: ID`
  - 响应同上（单条）
- 订单完成 `PUT /api/order/finish?order_id: ID`
  - 响应 `data`: `"订单已完成"`

## 里程碑与验收标准
- 里程碑 1：完成 P0（配置集中化、鉴权容错、ApiService 统一、API.md 全面对齐）；登录稳定、401 正常、接口切环境不需要改代码
- 里程碑 2：完成 P1（类型统一、通用组件与样式、路由角色守卫、错误提示统一）；代码可读性与一致性明显提升
- 里程碑 3：完成 P2（最小单测与构建告警消除）；修改风险降低

## 风险与规避
- JWT 信息不完全：采取容错与兜底策略，避免因缺字段导致崩溃
- 不同环境差异：通过 `env` 与常量集中化规避；把 `baseURL` 与 `/uploads` 前缀外置
- 变更影响面：以迭代方式推进，优先不破坏现有路由/页面行为，确保最小回归风险

---
如确认以上方向，我将按“P0 → P1 → P2”迭代执行，并在每个里程碑交付可运行版本与变更清单，确保你随时可回滚与对齐后端文档。
