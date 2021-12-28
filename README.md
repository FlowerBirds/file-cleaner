# file-cleaner
文件删除工具，根据delete.conf文件中的内容删除

## 用法
```
Usage: file-cleaner [options...]
Version: 0.0.3
删除delete.conf文件中列出来的文件，需要设置CLOUD_HOME环境变量或者通过-c进行指定根目录

Options:
        -c set CLOUD_HOME directory for find delete.conf and delete files.
        -d set the delete conf file name at CLOUD_HOME dir, default is delete.conf.
        -h help info.

```

例如：
```
set CLOUD_HOME=D:\tempo
file-cleaner.exe

```

## 编译
```shell
make

output binary file: file-cleaner.exe version 0.0.3
buil success
Usage: file-cleaner [options...]
Version: 0.0.3
删除delete.conf文件中列出来的文件，需要设置CLOUD_HOME环境变量或者通过-c进行指定根目录

Options:
        -c set CLOUD_HOME directory for find delete.conf and delete files.
        -d set the delete conf file name at CLOUD_HOME dir, default is delete.conf.
        -h help info.

```
在Makefile中最后执行了exe，导致退出码不为0，所以make后有`make: *** [Makefile:10: build] Error 1`

编译时修改Makefile中的版本号VERSION变量即可
```makefile
.PHONY: build

VERSION=0.0.3
BANIRY=file-cleaner.exe

build:
	@go build -ldflags "-X 'main.Version=${VERSION}'" -o ${BANIRY} main.go
	@echo output binary file: ${BANIRY} version ${VERSION}
	@echo buil success
	@${BANIRY}
```

## 编译所有架构
参考文章：[Go 使用Makefile编译所有架构](https://www.nhooo.com/note/qa5geu.html)