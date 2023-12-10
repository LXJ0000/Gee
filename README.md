# Gee

## 上下文
```bash
# 测试
curl -i localhost

curl -i localhost/hello?name=lxj

curl -i "localhost/login" -X POST -d "username=LXJ&password=root"
```
## 前缀树路由
```bash
curl "http://localhost/hello/jannan"
curl "http://localhost/assets/css/1.css"
```
## Group
```bash
curl -i "127.0.0.1/index" 

curl -i "127.0.0.1/v1"
curl -i "127.0.0.1/v1/hello?name=Jannan"

curl -i "127.0.0.1/v2" 
curl -i "127.0.0.1/v2/hello/Jannan"
```