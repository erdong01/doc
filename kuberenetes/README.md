

### 启用阿里云 Docker源

https://developer.aliyun.com/mirror/

###  启动并加入开机启动
```
systemctl start kubectl
systemctl enable kubectl

systemctl start kubelet

journalctl -xefu kubelet
```
### 安装 cri-dockerd
https://github.com/Mirantis/cri-dockerd

解决方法： 重新安装CNI网络插件(实践时采用了虚拟机，可能是因为当时使用的快照没包含网络插件)，然后重新清理结点，最后再重新加入结点
```
CNI_VERSION="v1.2.0"
mkdir -p /opt/cni/bin
curl -L "https://github.com/containernetworking/plugins/releases/download/${CNI_VERSION}/cni-plugins-linux-amd64-${CNI_VERSION}.tgz" | sudo tar -C /opt/cni/bin -xz
```

#### 配置转发相关参数，否则可能会出错
vim  /etc/sysctl.conf
```
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_nonlocal_bind = 1
net.ipv4.ip_forward = 1
vm.swappiness=0

```
sysctl --system


### 添加路由规则
https://www.iteye.com/blog/codyjava-2435846
```
iptables -P INPUT ACCEPT
iptables -P FORWARD ACCEPT
iptables -F
iptables -L -n
iptables -t nat -I POSTROUTING -s 10.244.0.0/24 -j MASQUERADE
iptables-save
```

#### kubectl: 配置操作

echo "export KUBECONFIG=/etc/kubernetes/admin.conf" >> ~/.bash_profile

source ~/.bash_profile

#### kubeadm：命令

* 生成TOKEN kubeadm token generate
* 初始化：sudo kubeadm init --config ./init.yaml --max-mutating-requests-inflight 3000  --max-requests-inflight 1000
* 加入：sudo kubeadm join --config /opt/kubeadm/join.yaml
* 重置：sudo kubeadm reset
* 创建token:
````
    kubectl create serviceaccount dashboard-admin -n kube-system
    kubectl create clusterrolebinding dashbord-admin --clusterrole=cluster-admin --serviceaccount=kube-system:dashboard-admin
````
* 获取token
````
    kubectl describe secret dashboard-admin-token-cdssv -n kube-system
````

```
kubeadm reset --cri-socket unix:///var/run/cri-dockerd.sock

sudo kubelet --container-runtime-endpoint=unix:///var/run/cri-dockerd.sock
```
#### 允许master污点：
```
kubectl taint nodes --all node-role.kubernetes.io/master-

kubectl taint node master node-role.kubernetes.io/master-
```
允许master污点 https://zhuanlan.zhihu.com/p/446621713

### 日志:
#### 记录架构
https://kubernetes.io/zh/docs/concepts/cluster-administration/logging/

https://github.com/ielepro/kubernetes-efk/blob/master/fluentd_daemonset.yaml

https://github.com/kubernetes/kubernetes/tree/master/cluster/addons/fluentd-elasticsearch


### 问题：

1.
    Get http://localhost:10248/healthz: dial tcp [::1]:10248: connect: connectio
    问题原因：kubelet之前的配置没有清理干净，Github 上说systemctl stop kubepods-burstable.slice即可解决问题，但是测试并没有作用。kubeadm reset之后重启系统解决问题：
    
    journalctl -xeu kubelet

    ifconfig cni0 down
    
    ifconfig flannel.1 down
    
    ifconfig del flannel.1
    
    ifconfig del cni0
    
    ip link del flannel.1
    
    ip link del cni0
    
    brctl delbr  flannel.1
    
    brctl delbr cni0
    
    rm -rf /var/lib/cni/flannel/* && rm -rf /var/lib/cni/networks/cbr0/* && ip link delete cni0 &&  rm -rf /var/lib/cni/network/cni0/*
    
    kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep dasboard-admin | awk '{print $1}')



### 使用技巧：

#### 1. Liveness和Readiness探针
默认情况下Liveness和Readiness探针是没有被设置的。并且有时候一直保持这种情况。

但是当出现不可恢复的错误，你的服务该怎样重启？负载平衡器如何知道特定的Pod可以开始处理流量？或者处理更多的流量？

人们通常不知道这两者的区别。

Liveness探针在失败的时候会重启Pod

Readiness探针在失败后会将失败的Pod和Kubernetes服务（可以通过kubectl get endpoints检测）断开连接，并且流量不再分发给这个Pod，除非探测再次成功

