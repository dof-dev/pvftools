# Repository Guidelines

## Project Structure & Module Organization
`main.go` 是 Wails 应用入口。后端代码集中在 `backend/`：`api/` 暴露给前端的 App 方法，`internal/` 放领域逻辑，`dao/` 与 `model/` 处理数据访问和结构体，`common/` 存放设置、常量和通用工具。前端位于 `frontend/src/`：`views/` 是页面级视图，`components/` 是复用组件，`common/` 是共享工具。`frontend/wailsjs/` 为 Wails 生成代码，除非后端接口变更，否则不要手改。`build/` 保存各平台打包资源。

## Build, Test, and Development Commands
先在 `frontend/` 下执行 `npm install` 安装前端依赖；本仓库的 `wails.json` 也默认使用 `npm`。常用命令：

- `go test ./...`：运行 Go 包测试，并作为后端编译健康检查。
- `cd frontend && npm run dev`：仅启动 Vite 前端调试。
- `cd frontend && npm run build`：执行 `vue-tsc` 类型检查并构建前端。
- `wails dev`：启动桌面应用开发模式，联调 Go 和 Vue。
- `wails build`：生成桌面应用安装产物，Windows 打包配置见 `build/windows/`。

建议使用 Go `1.22.x` 和 Node `20.10.0`。

## Coding Style & Naming Conventions
Go 代码保持 `gofmt` 默认格式，包名全小写，文件名沿用现有下划线风格，如 `data_loader`、`user_storage_dao.go`。导出方法使用 PascalCase，局部变量使用简短 camelCase。Vue/TypeScript 延续现有 2 空格缩进，组件和视图文件使用 PascalCase，例如 `SkillView.vue`、`ItemSelector.vue`。改动遵循最小化原则，优先追加到现有相关模块后面，不做大范围重构。

## Testing Guidelines
当前仓库没有完整的自动化前端测试；提交前至少运行 `go test ./...` 和 `cd frontend && npm run build`。涉及界面改动时，再用 `wails dev` 手动验证对应页面。新增 Go 测试请放在目标包旁边，命名为 `*_test.go`。`backend/test/cli.go` 更像本地实验入口，不应替代正式测试。

## Commit & Pull Request Guidelines
现有提交历史以简短前缀为主，如 `feat: 任务生成器`、`fix: load job map`，新提交建议继续使用 `<type>: <summary>`。PR 应说明改动目的、影响范围、手动验证步骤；涉及 UI 时附截图。若后端接口未变更，避免混入 `frontend/wailsjs/` 生成文件；提交前也不要带入本机路径、测试数据或任何密钥。

## Security & Configuration Tips
仓库中存在本地实验代码和绝对路径示例，例如 `backend/test/cli.go`；提交前确认这些内容没有被扩散到正式逻辑。不要提交 `.env`、访问凭据、PVF 私有数据或带用户名的系统路径。发布流程目前由 `.github/workflows/build.yml` 在打 `tag` 时触发，若改动构建链路，请同时检查 `wails.json` 和 `build/` 下的平台配置是否一致。
