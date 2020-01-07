# deploy
Deploy service in kubernetes

[![Build Status](https://travis-ci.com/tio-serverless/deploy.svg?branch=master)](https://travis-ci.com/tio-serverless/deploy) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=tio-serverless_deploy&metric=alert_status)](https://sonarcloud.io/dashboard?id=tio-serverless_deploy)

### Environment

+ TIO_DEPLOY_CONFIG 配置文件路径



### 配置说明

+ 默认配置

```toml
log="debug"
port=80
[inject]
  grpc=""
  http=""
[k8s]
  config=""
  consul=""
  //部署使用的Sidecar, 默认使用tioserverless/consul-agent
  sidecar="tioserverless/consul-agent:v0.1.0-develop"
```