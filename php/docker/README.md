## makefile:

##### 配置:


    VERSION=2190                                          镜像版本号  备注：我们会跟tower任务编号一样

    DOCKER_PROVIDER=registry.cn-shanghai.aliyuncs.com     docker仓库供应商



##### 操作命令:

    local_build              本地环境构建docker
    
    production_build         生产环境构建docker
    
    composer_install         安装本地composer包
    
    composer C=操作命令       composer 命令
    
    command C=操作命令        laravel artisan  因为和make artisan冲突换command
    
    restart                  重启docker
    
    start                    启动docker
    
    start                    停止docker
    
    pull                     拉取docker镜像 主要生产环境用
    
    push                     推送docker镜像 需要事先配置好账号
    
    docker_login             登录docker仓库
    
## 阿里云私有仓库登录
    https://cr.console.aliyun.com/cn-shanghai/instances/credentials
    
## 本地开发环境

##### 1. make git 拉取代码并进入项目目录
##### 2. make make local_build  构建docker
##### 3. make composer_install 本地composer依赖包更新
##### 4. make local_run  启动运行 

#### 其它命令：
    
    示例1. make  command C=config:cache  等于 php artisan config:cache 缓存配置
    
## 生产环境构建

##### 1. git 拉取代码并进入项目目录
##### 2. make production_build 生产环境构建docker
##### 3. make push  提交私有仓库

备注：私有仓库需要先登录设置用户权限

## 生产环境机器部署

##### 1：拷贝makefile 文件夹
##### 2: make production_run

备注：私有仓库需要先登录设置用户权限
