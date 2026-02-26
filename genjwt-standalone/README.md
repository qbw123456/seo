# JWT accessToken 生成工具（独立版）

从 serp-server-seo 项目抽出的独立程序，用于生成与后端校验一致的 JWT，供 Apifox 等调试使用。

## 使用

1. 进入本目录，安装依赖并编译：
   ```bash
   cd D:\seo\genjwt-standalone
   go mod tidy
   go build -o genjwt.exe .
   ```

2. 运行（配置文件指向 SEO 项目的 config）：
   ```bash
   .\genjwt.exe -c D:\seo\serp-server-seo-master\resources\config.dev.yaml -u 37
   ```
   - `-c`：配置文件路径（必须，使用 SEO 项目里的 `resources/config.dev.yaml`）
   - `-u`：用户 ID（可选，默认 37，即 qiaobowen）

3. 将输出的 Token 填入 Apifox：**Authorization** 值为 `Bearer <输出的整段 Token>`。

## 不编译直接运行

```bash
go run . -c D:\seo\serp-server-seo-master\resources\config.dev.yaml -u 37
```
