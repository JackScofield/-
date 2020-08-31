## Dockerfile定制镜像

之前有一个对`image`的 **黑箱操作**，叫做`docker commit` 但是一般开发过程中不会使用，因为除了操作者本身，别人看不到修改了什么。

此外我们知道 `image`的定制实际上就是定制每一层所添加的配置，文件。我们可以写一个脚本，把每一层的修改，安装，构建和操作的命令都写进去。这个脚本就叫 `Dockerfile` 



### FROM指定一个基础镜像

如果我们在`dockfile`上面建立的镜像是要从一个已有的镜像基础上建立的, FROM 必须是第一个指令。基础镜像包括：

* 服务类镜像：nginx, redis, mongo, mysql, php, tomcat
* 语言应用的镜像：node, openjdk, python, golang
* 操作系统镜像：ubuntu，centos，debian，fedora

如果我们在Dockerfile不想在任何基础上，就可以

```
FROM scratch
```



### RUN 执行命令

* *shell* 格式：`RUN <命令>`，就像直接在命令行中输入的命令一样。刚才写的 Dockerfile 中的 `RUN` 指令就是这种格式。
* *exec* 格式：`RUN ["可执行文件", "参数1", "参数2"]`，这更像是函数调用中的格式。



docker run -dp 9090:9090 -v /tmp/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml -v /tmp/prometheus/alert.rules:/etc/prometheus/alert.rules --name prometheus prom/prometheus