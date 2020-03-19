## TgoShorten
基于mux和redis实现的短地址服务

#### 安装 & 运行

```bash
git clone git@github.com:nashlibby/TgoShorten.git
go build
./tgo-shorten --port=8088
```
#### 接口

> 生成短地址

地址
```bash
POST /api/v1/shorten
```

参数 json格式
- url `string` 长链接地址
- expiration `int` 过期分钟

>  获取短地址信息

地址
```bash
GET /api/v1/info
```

参数 
- url `string` 短链接地址

> 链接跳转

地址
```bash
GET /:url
```
参数 
- url `string` 短链接地址




