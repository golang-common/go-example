# GRPC基础操作

## 安装proto buffer
1. 下载相应平台的命令安装包https://github.com/protocolbuffers/protobuf/releases
2. 将执行文件加入$PATH
3. 执行protoc即可

## 安装GO的proto buffer插件
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

go get -u google.golang.org/grpc

## 使用protoc编译.proto文件

protoc --proto_path=IMPORT_PATH 
--cpp_out=DST_DIR --java_out=DST_DIR 
--python_out=DST_DIR --go_out=DST_DIR --ruby_out=DST_DIR 
--javanano_out=DST_DIR --objc_out=DST_DIR --csharp_out=DST_DIR path/to/file.proto

## GO编译参数

- -I 参数：指定import路径，可以指定多个-I参数，编译时按顺序查找，不指定时默认查找当前目录,就是如果多个proto文件之间有互相依赖，生成某个proto文件时，需要import其他几个proto文件，这时候就要用-I来指定搜索目录
- --go_out ：golang编译支持，支持以下参数
    - plugins=plugin1+plugin2 - 指定插件，目前只支持grpc，即：plugins=grpc
    - M 参数 - 指定导入的.proto文件路径编译后对应的golang包名(不指定本参数默认就是.proto文件中import语句的路径)
    - import_prefix=xxx - 为生成的go文件的所有import路径添加前缀
    - import_path=foo/bar - 用于指定proto文件中未声明package或go_package的文件的包名，最右面的斜线前的字符会被忽略
    - 末尾 :编译文件路径 .proto文件路径(支持通配符)

## 完整示例
protoc -I . --go_out=plugins=grpc,Mfoo/bar.proto=bar,import_prefix=foo/,import_path=foo/bar:. ./*.proto

## 证书制作
### 制作私钥(.key)
>Key considerations for algorithm "RSA" ≥ 2048-bit

openssl genrsa -out server.key 2048
    
>Key considerations for algorithm "ECDSA" ≥ secp384r1
List ECDSA the supported curves (openssl ecparam -list_curves)

openssl ecparam -genkey -name secp384r1 -out server.key

### 自签名公钥(x509)(PEM-encodings .pem|.crt)
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650

### 自定义信息
Country Name (2 letter code) [AU]:CN
State or Province Name (full name) [Some-State]:XxXx
Locality Name (eg, city) []:XxXx
Organization Name (eg, company) [Internet Widgits Pty Ltd]:XX Co. Ltd
Organizational Unit Name (eg, section) []:Dev
Common Name (e.g. server FQDN or YOUR name) []:server name
Email Address []:xxx@xxx.com

