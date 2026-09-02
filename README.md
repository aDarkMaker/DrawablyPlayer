# DrawablyPlayer

Drawably 风格的跨平台音乐 / 视频播放器。

基于 **Wails v3 + Go + React(TypeScript) + Vite(bun)**，同一套代码库支持 Windows / macOS / Linux 桌面端，以及实验性的 Android / iOS 移动端。

## 技术栈

- [Wails v3](https://wails.io) + Go 1.25
- React + TypeScript + Vite（`bun` 管理前端依赖）
- 构建编排：根 `Taskfile.yml`（Taskfile v3，被 `wails3` CLI 驱动）

## 目录结构

```
.
├── main.go               # 应用入口（embed frontend/dist + appicon）
├── Taskfile.yml          # 根构建编排（build/package/run/dev，按 GOOS 分发）
├── build/                # 各平台构建资产与 Taskfile
│   ├── config.yml        # 产品元信息（wails3 update build-assets 使用）
│   ├── darwin/ windows/ linux/ ios/ android/
├── frontend/             # React + TS + Vite 前端
│   └── bindings/         # 生成的 TS 绑定（不入库，构建时生成）
├── .github/workflows/    # CI 与自动发布
├── LICENSE / .gitignore / cspell.json / .prettierrc.cjs
```

## 开发

```bash
wails3 dev                          # 开发模式（前后端热更新）
wails3 task build                   # 生产构建（产物 bin/）
wails3 task run                     # 运行
wails3 task darwin:package:dmg      # macOS DMG
wails3 task package                 # Windows NSIS 安装包（windows 主机上）
wails3 task android:package:fat     # Android universal APK
```

## 发布

- push main 后由 [release-please](https://github.com/googleapis/release-please-action) 依据 Conventional Commits 自动定版、打 tag 并创建 release；
- 三个构建 job 并行产出 Windows exe + NSIS 安装包、macOS(arm64/amd64) DMG、Android universal APK，自动上传到该 release。

## License

[MIT](LICENSE)
