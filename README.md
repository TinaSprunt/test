# kubesphere架构

![kubesphere架构图](https://pek3b.qingstor.com/kubesphere-docs/png/20190627223641.png)


# console本地调试流程

> 参考官方console github文档
> https://github.com/kubesphere/console
> https://github.com/kubesphere/console/blob/master/docs/access-backend.md

1. 公开ks-apiserver服务，开启调试端口30881。然后就可以通过<node_ip>：<30881>端口访问ks-apiserver服务，命令如下
```shell
kubectl -n kubesphere-system patch svc ks-apiserver -p '{"spec":{"type":"NodePort","ports":[{"name":"ks-apiserver","port":80,"protocal":"TCP","targetPort":9090,"nodePort":30881}]}}'
```
2. 下载console代码
3. 切换到 console/server/下建立 local_config.yaml 文件，用于指定当前本地调式的这份console代码调用那个后端，文件内容如下
```yaml
server:
  apiServer:
    url: http://[后端ip]:30881
    wsUrl: ws://[后端ip]:30881
```
4. 安装node.js及yarn
```shell
# 安装node.js官方要求是node版本大于10.16的都可以,但是实际上装15.x以上的就有问题，10.16.x到15.x之间比较安全
$ curl --silent --location https://rpm.nodesource.com/setup_10.x | sudo bash -
yum install -y nodejs

#安装yarn
$ npm install -g yarn
```
> 如果出现包问题：参考
> https://blog.csdn.net/qq_33909098/article/details/109455087


5. 执行如下yarn命令
```shell
# 命令一
$ yarn
# 命令二
$ yarn lego
# 命令三
$ yarn start
# 如果yarn start执行失败的原因是 node-sass没有下载成功
$ yarn add node-sass@4.14.1  #默认会装5.0.0这个装上也用不了，需要指定版本号
# 在vscode本地运行不用yarn start
$ yarn dev:client
$ yarn dev:server
```
> ps:没有yarn的先安装yarn，国内用户可以先切换淘宝镜像会快一点
```shell
# 安装yarn
npm install -g yarn
# 切换淘宝镜像，在执行 yarn lego之前
yarn config set registry https://registry.npm.taobao.org
```
5.执行成功就可以在页面使用 **http://[后端ip]:8000** 进行访问了,每次修改之后重新执行第三步的yarn命令
> ps: 本地调试console访问后端用30881端口，界面访问console用8000端口，并不是界面直接访问30881


6. 源码编译打包
> 参考官网最新更新的3.0的编译文档：
> https://github.com/kubesphere/community/pull/172/files
> ps : 编译部分目前还没有实践，本地调试部分已实践可以   


7. 功能调试需要API文档
> v3.0.0不再集成swagger了，所以接口的调用方式也和目前官网文档提供的2.1版本API访问方式有所不同，参考官方github文档
> https://github.com/kubesphere/website/blob/3698258e1fea8f50fa1c6eb495e4124dc0d3f9cc/content/en/docs/api-documentation/kubesphere-api.md
