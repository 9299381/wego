# 基础镜像
FROM scratch
# 镜像作者
MAINTAINER 9299381@qq.com

# 复制文件或目录到容器指定路径中
COPY . /home

# 容器对外映射的本地端口，需要在 docker run 的时候使用-p或者-P选项生效
EXPOSE 8341


# 为后续的RUN、CMD、ENTRYPOINT指令配置工作目录
WORKDIR /home

# 在启动容器时提供一个命令执行选项，这里运行我们的应用
#CMD ["/home/main"]