# 部署目录说明（deploy）

本目录用于存放项目部署相关的所有配置、脚本和文档，建议结构如下：

```
deploy/
├── README.md                # 部署目录说明
├── docker/
│   ├── Dockerfile           # 生产环境Dockerfile（如需多环境可分开）
│   └── entrypoint.sh        # 容器启动脚本（可选）
├── k8s/
│   ├── deployment.yaml      # Kubernetes Deployment配置
│   ├── service.yaml         # Kubernetes Service配置
│   └── ingress.yaml         # Kubernetes Ingress配置（如有）
├── systemd/
│   └── gin-init.service     # systemd服务配置示例
├── nginx/
│   └── gin-init.conf        # Nginx反向代理配置示例
├── env/
│   ├── .env.example         # 环境变量示例文件
│   └── .env.prod            # 生产环境变量（建议不入库）
└── scripts/
    ├── build.sh             # 构建脚本
    ├── start.sh             # 启动脚本
    └── stop.sh              # 停止脚本
```

## 推荐部署项说明

- **docker/**：容器化相关文件，支持本地和云原生部署。
- **k8s/**：Kubernetes部署清单，适合云原生环境。
- **systemd/**：Linux下以服务方式运行的配置。
- **nginx/**：常用反向代理配置。
- **env/**：环境变量模板，便于多环境切换。
- **scripts/**：常用自动化脚本，提升运维效率。

> 建议所有敏感信息（如数据库密码）通过环境变量或K8s Secret注入，不要直接写入配置文件。

如需具体某项部署文件模板，请告知！ 