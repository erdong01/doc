#### server:  

config.json
```
{
	"datacenter":"rxt",//数据中心名称
	"node_name":"dev_192_168_31_243",
	"bootstrap_expect"   : 2,                     
	"client_addr"        : "192.168.31.243",  
	"bind_addr"          : "192.168.31.243",    
	"ui":true,
	"raft_protocol"      : 3,
    "enable_debug"       : false,
    "enable_syslog"      : false,
	"server"             : true,
	"log_level"          : "INFO",
	"data_dir"           : "/var/lib/consul",
	"retry_interval"     : "3s",
    "rejoin_after_leave" : true,
    "retry_join": ["192.168.31.241","192.168.31.243"],
    "start_join": ["192.168.31.241","192.168.31.243"],

}


```

acl.json

```

{
  "acl_datacenter": "rxt",
  "acl_master_token": "p2BE1AtpwPbrxZdC6k+eXA==",
  "acl_default_policy": "deny",
  "acl_down_policy": "extend-cache",
  "acl_agent_token": "4ac10325-551b-5465-05e0-e0a48a04971a"
}

```

### client:  

config.json
```
{  
    "datacenter":"rxt",   //数据中心名称
    "node_name":"dev_192_168_31_240", //节点                 
    "client_addr"        : "192.168.31.240",    
    "bind_addr"          : "192.168.31.240",
    "retry_join":["192.168.31.243"],
    "start_join":["192.168.31.243"],    
    "raft_protocol"      : 3,  
    "enable_debug"       : false,  
    "enable_syslog"      : false,  
    "log_level"          : "INFO",  
    "data_dir"           : "/var/lib/consul",  
    "retry_interval"     : "3s",  
    "server": false,  
    "ui": false  
}

```
acl.json

```

{
    "acl_datacenter": "rxt",
    "acl_down_policy": "extend-cache",
    "acl_token":"88fbcc3d-d529-026d-3edf-fb6825b876f4"
}

```

#### 密钥默认地址 data_dir 下面： /var/lib/consul/serf/local.keyring

datacenter 数据中心 server 和 server 和 client 之间必须是同一个

### 安装consul

#### linux 
consul下载地址：https://releases.hashicorp.com/consul/1.6.0/
1. wget https://releases.hashicorp.com/consul/1.6.0/consul_1.6.0_linux_amd64.zip

#解压下载的Consul，如果没有安装unzip的话可以先安装：yum install zip unzip
2. unzip consul_1.6.0_linux_amd64.zip -d /usr/local/bin

3. 编辑 /etc/profile 文件，添加环境变量，保存后退出 

```

vim /etc/profile
export CONSUL_HOME=/usr/local/bin/consul
export PATH=$PATH:CONSUL_HOME

```
使用环境变量配置生效
source /etc/profile

### 运行consul
 
1. 读取配置文件运行 

    consul agent -config-dir=/etc/consul.d
    
    /etc/consul.d目录下config.json 