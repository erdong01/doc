### 安装consul-template

#### linux 
consul-template：https://releases.hashicorp.com/consul-template/0.21.3/
1. wget https://releases.hashicorp.com/consul-template/0.21.3/consul-template_0.21.3_linux_amd64.zip

#解压下载的Consul，如果没有安装unzip的话可以先安装：yum install zip unzip
2. unzip consul-template_0.21.3_linux_amd64.zip -d /usr/local/bin

### 部署
1. consul-template 和 Nginx 必须在同一台服务器 

### 常见问题
注册服务 consul 服务名称service name 不要用点号 . 会发现不了

https://www.hi-linux.com/posts/36431.html

### 配置

nginx.hcl

```  

consul{
        address="192.168.31.243:8500"
}
template{
        source = "nginx.conf.ctmpl" //配置模版文件
        destination="/etc/nginx/consul/default.conf" //生成存放位置
        command="service nginx reload"
}


```

nginx.conf.ctmpl

```  

    upstream rxtai{
        {{range service "rxt_ai"}}    
            server {{.Address}}:{{.Port}} fail_timeout=0;    
        {{end}}    
    }
    server {
        listen       80;
        server_name  192.168.31.243;
        charset utf-8;
        location / {
            proxy_pass http://rxtai;
            proxy_connect_timeout 1;
                proxy_read_timeout 1;
                proxy_send_timeout 1;
        }
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }

```

### 运行命令 

1. consul-template -config "nginx.hcl" 启动



