# 项目架构与开发注意事项

## 项目概览
- 技术栈：Vue 3（Composition API）+ TypeScript + Vite + Pinia + Vue Router + Element Plus
- 统一响应：`BaseResp<T> = { code, message, data }`
- 环境配置集中化：`src/config/env.ts` 提供 `API_BASE_URL`、`UPLOAD_BASE_URL`、`REQUEST_TIMEOUT`

## 目录结构
- `src/config/` 环境与常量
  - `env.ts` 暴露 `API_BASE_URL`、`UPLOAD_BASE_URL`、`REQUEST_TIMEOUT`
- `src/api/` 网络层
  - `axios.ts` 统一 Axios 实例与拦截器（读取环境配置）
  - `login.ts`、`order.ts`、`summary.ts` 等按资源域拆分
- `src/stores/` 全局状态（Pinia）
  - `user.ts` 登录态、头像、角色等
- `src/types/` DTO 与类型转换
  - 每个接口一个文件，定义 `Resp/Model` 与 `toModel/toResp`
- `src/components/` 通用组件
  - `components/common/` 全局复用组件（如 `GlobalBack`、`GlobalLogout`、`NoticeCircle`、`BaseButton` 等）
- `src/layouts/` 页面骨架（如 `AuthLayout.vue`）
- `src/views/` 路由页面（`Home`、`OrderList`、`OrderDetail`、`CreateSummary`、`SummaryList` 等）
- `src/styles/` 全局样式与设计令牌（`base.css`、`tokens.css`、`utilities.css`）
- `src/assets/` 静态资源（图片、图标等）

## 放置规范
- 环境与常量：所有服务地址/超时等集中在 `src/config/env.ts`，业务代码不得硬编码
- 网络层：统一使用 `src/api/axios.ts` 的实例；按资源域拆分 API 文件
- 类型：所有 ID（雪花 `int64`）统一用 `string`；时间戳字段按整数处理；避免 `any`
- 组件：可复用组件放 `components/common/`；页面专用组件内聚到对应 `view`
- 路由与布局：布局放 `layouts/`；路由页面放 `views/`
- 样式：全局样式与设计令牌放 `styles/`；组件样式使用 `scoped`

## 环境与配置
- 变量说明（在 `.env.*` 中设置）：
  - `VITE_API_BASE_URL`：后端 API 基地址（如 `http://127.0.0.1:9000`）
  - `VITE_UPLOAD_BASE_URL`：静态资源前缀（如 `http://127.0.0.1:9000/uploads/`）
  - `VITE_REQUEST_TIMEOUT`：请求超时（默认 `10000`）
- 代码读取：由 `src/config/env.ts` 统一读取 `import.meta.env` 并提供默认值

## 网络层与错误处理
- 请求拦截器：统一注入 `Authorization: Bearer <token>`（从 `Pinia` 拿）
- 响应拦截器：默认返回 `response.data`
- 建议统一错误处理：
  - 401：清理登录态、跳转登录页并保留 `redirect`
  - 业务错误：读取 `message` 提示
  - 网络错误：提示“网络异常或服务器无响应”

## 类型约定
- ID 类型：统一 `string`（避免 JS 精度问题）
- 时间戳：`ctime/utime` 使用整数时间戳（常见为秒），展示时本地格式化
- DTO 转换：`Resp → Model`、`Model → Resp` 由 `toModel/toResp` 负责

## 路由与权限
- 登录守卫：通过 `meta.requiresAuth` 控制受限页面访问
- 公共页重定向：已登录访问 `Login/Register/ResetPassword` 时重定向到主页
- 建议扩展：在 `meta.roles` 指定角色数组，守卫中统一判定（如 `COMMON_USER/INTERNAL_USER/ADMIN`）

## 组件与样式
- 通用组件：放在 `components/common/`，供跨页面复用
- 页面骨架：`layouts/` 中定义（如登录/注册用 `AuthLayout`）
- 样式：
  - 全局样式与设计令牌放在 `styles/*`
  - 组件内部样式使用 `scoped`，避免污染

## 开发注意事项
- 禁止在业务代码硬编码服务地址或上传前缀，统一使用 `src/config/env.ts`
- 接口层与 DTO 不使用 `any`，请求/响应类型明确且有转换函数
- 路由页面专注展示与交互；数据请求在 `api/*`，全局状态在 `stores/*`
- 静态资源展示统一用 `UPLOAD_BASE_URL + fileName`
- 引入外部库前先检查项目依赖并遵循现有用法与风格

## 快速上手
- 启动开发：`npm run dev`
- 类型检查：`npm run type-check`
- 格式化：`npm run format`
- 构建产物：`npm run build`

## 最佳实践建议（后续可迭代）
- 统一 401 处理与错误提示策略（响应拦截器增强）
- 在路由 `meta.roles` 上进行角色守卫，减少页面内权限逻辑分散
- 抽取更多通用组件（如 `EmptyState/Loading`）与样式变量，提升复用与一致性
- 引入 ESLint 规则与最小单测（如 `vitest` 覆盖路由守卫与网络层错误分支）