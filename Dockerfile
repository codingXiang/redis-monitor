# Build 專案;
FROM    golang:alpine AS builder
LABEL   maintainer="NianXiang Lai"
USER    root

ENV     RUN_PATH=/monitor PROJ_PATH=${GOPATH}/src/RedisMonitor/

RUN     mkdir -p $RUN_PATH && apk add --no-cache ca-certificates bash git make build-base
ADD     . ${PROJ_PATH}
WORKDIR ${PROJ_PATH}
RUN     make all \
        && make dockerpack \
        && tar -zxf redis-monitor-v*.tar.gz -C ${RUN_PATH} \
        && rm -rf ${PROJ_PATH} 

# 打包 Image;
FROM    alpine
LABEL   maintainer="NianXiang Lai"
USER    root
ENV     RUN_PATH=/monitor
RUN     mkdir -p $RUN_PATH && apk add --no-cache ca-certificates bash git
COPY    --from=builder ${RUN_PATH} ${RUN_PATH}
WORKDIR ${RUN_PATH}
RUN     chmod +x *.sh
ENTRYPOINT ["bash", "-c", "entrypoint.sh"]