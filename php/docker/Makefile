ROOT=$(shell pwd )

#镜像版本号  备注：我们会跟tower任务编号一样
VERSION=2190


#暴露端口号
PORT=5200
#容器命名空间  对应的是docker仓库的命名空间
NAMESPACE=renxuetang_est
#容器名词  对应的是docker仓库的名称
NAME=api
#生产环境项目地址
PRODUCTION_FILE=/usr/src/production
#本地环境项目地址
LOCAL_FILE=/usr/src/est
#启动项目
START=php ./bin/laravels start
#更新composer命令
COMPOSER_UPDATE="composer update --no-plugins"
#docker仓库供应商
DOCKER_PROVIDER=registry.cn-shanghai.aliyuncs.com

DOCKER_NAME=$(DOCKER_PROVIDER)/$(NAMESPACE)/$(NAME):$(VERSION)


#本地环境启动docker
local_run:
	docker run -d -p $(PORT):$(PORT) --name=$(NAMESPACE)_$(NAME)_$(VERSION) -v $(ROOT):$(LOCAL_FILE) -it $(DOCKER_NAME) $(START)

#本地环境构建docker
local_build:
	docker build --build-arg FILE=$(LOCAL_FILE) --build-arg COMPOSER_UPDATE="" --build-arg PORT=$(PORT) -t $(DOCKER_NAME) .

#生产环境启动docker
production_run:
	docker run -d -p $(PORT):$(PORT) --restart="on-failure"  --name=$(NAMESPACE)_$(NAME)_$(VERSION) -it $(DOCKER_NAME) $(START)

#生产环境构建docker
production_build:
	docker build  --build-arg FILE=$(PRODUCTION_FILE)  --build-arg COMPOSER_UPDATE=$(COMPOSER_UPDATE)  --build-arg PORT=$(PORT) -t $(DOCKER_NAME) .

#更新本地composer包
composer_install:
	docker run  --name=$(NAMESPACE)_$(NAME)_composer_$(VERSION)   -v $(ROOT):$(LOCAL_FILE) -it $(DOCKER_NAME) composer update --no-plugins

#composer 命令
composer:
	docker exec -it $(NAMESPACE)_$(NAME)_$(VERSION) composer $(C)

#laravel artisan  因为和make artisan冲突换command
command:
	docker exec -it $(NAMESPACE)_$(NAME)_$(VERSION) php artisan $(C)

#重启docker
restart:
	docker restart $(NAMESPACE)_$(NAME)_$(VERSION)

#启动docker
start:
	docker start $(NAMESPACE)_$(NAME)_$(VERSION)

#停止docker
stop:
	docker stop $(NAMESPACE)_$(NAME)_$(VERSION)

#拉取docker镜像
pull:
	docker pull $(DOCKER_NAME)

#推送docker镜像
push:
	docker push $(DOCKER_NAME)


docker_login:
	docker login --username=$(USERNAME) $(DOCKER_PROVIDER)
