# testlikexin-12-15-demo

Frontend for testlikexin-12-15-demo

## 技术栈

- ⚛️ React 18
- ⚡ Vite
- 📘 TypeScript
- 🎨 CSS3

## 快速开始

### 安装依赖

```bash
npm install
# 或
pnpm install
```

### 启动开发服务器

```bash
npm run dev
# 或
pnpm dev
```

访问: http://localhost:8080

### 构建生产版本

```bash
npm run build
# 或
pnpm build
```

### 预览生产构建

```bash
npm run preview
# 或
pnpm preview
```

## 项目结构

```
testlikexin-12-15-demo/
├── public/          # 静态资源
├── src/
│   ├── components/  # 组件目录
│   ├── pages/       # 页面目录
│   ├── App.tsx      # 主应用组件
│   ├── main.tsx     # 应用入口
│   └── index.css    # 全局样式
├── index.html       # HTML 模板
├── package.json     # 依赖配置
├── tsconfig.json    # TypeScript 配置
└── vite.config.ts   # Vite 配置
```

## 开发指南

### 添加新页面

1. 在 `src/pages/` 目录创建新的页面组件
2. 在 `App.tsx` 中添加路由

### 添加新组件

在 `src/components/` 目录创建可复用的组件

### API 调用

项目已集成 axios，可以在 `src/api/` 目录创建 API 服务

## 部署

### 部署到 Vercel

```bash
npm install -g vercel
vercel
```

### 部署到 Netlify

```bash
npm install -g netlify-cli
netlify deploy --prod
```

## 许可证

MIT

---

由 [EZ-Dev Platform](https://ezdev.com) 生成

