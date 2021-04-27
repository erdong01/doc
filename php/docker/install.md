
* 启用阿里云 Docker源

https://developer.aliyun.com/mirror/

* 安装Docker
```
sudo apt-get install docker-ce
```

*  启动并加入开机启动
```
systemctl start docker
systemctl enable docker
```

修改docker的cgroupdriver为systemd（默认为cgroups），因kubernetes默认使用的cgroupdriver为systemd（kubernetes自带的推荐使用）
```
vi /etc/docker/daemon.json

{
"exec-opts":["native.cgroupdriver=systemd"]
}

systemctl restart docker
```


docker build -t erdong01/marstm-app:v0.0.1 .
docker run -p 5003:5003 -v /var/code/go/src/rxt/log:/bin/log  -d -it  erdong01/rxt:v0.0.3