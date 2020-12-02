# jqe

jqe 是一个简单的 json 文件编辑命令工具，可以修改 json 文件中的任意字段。该工具启发于 jq 命令，由于 jq 命令无法进行 json 文件的编辑，所以开发了此工具。

# 使用方式

* 安装

```
# Linux
wget https://github.com/barnettZQG/jqe/releases/download/0.1/linux-jqe -O /usr/bin/jqe
Or
wget https://buildpack.oss-cn-shanghai.aliyuncs.com/common/utils/jqe -O /usr/bin/jqe

```

* 新增或更新某个字段
```
./jqe update -f test.json a1="a3"
./jqe update -f test.json a2.a33=43 -t int 
./jqe update -f test.json a3.a33=true -t bool
```

* 获取某字段的值
```
./jqe get a2.a33 -f test.json
```