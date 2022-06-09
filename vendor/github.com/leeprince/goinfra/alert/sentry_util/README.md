# sentry

## sentry 自托管（本地安装）
参考：https://develop.sentry.dev/self-hosted

### 安装步骤
1. 下载 [最新版本的自托管存储库](https://github.com/getsentry/self-hosted/releases/latest) ，
2. 运行 `./install.sh` 在这个目录里面。 
该脚本将处理您开始所需的所有事情，包括基线配置，然后会告诉您运行 docker-compose up -d 启动sentry。 
3. 启动sentry:`docker-compose up -d`
4. 访问 sentry:`http://127.0.0.1:9000`
Sentry 绑定到端口 9000默认。 您应该能够访问 http://127.0.0.1:9000 . 

> 安装过程可能的问题
1. mac 没有 `realpath` 命令
```
1. brew install coreutiles
2. .bash_profile 中新建函数，如果没有此文件则新建
`
function realpath(){
    [[ $1 = /* ]] && echo "1 &quot; || echo &quot; 1&quot; || echo &quot; 1" || echo "PWD/${1#./}"
}
`
```
2. 默认访问 9000 端口，如果本地安装了 `php-fpm` 可能无法启动
可以更改宿主机映射到 docker sentry 环境的端口非 9000 即可。
该端口实际上是访问到 docker sentry 环境中的  nginx 代理服务器的80，由 nginx 进行反向代理
