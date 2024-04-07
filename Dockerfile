FROM registry.cn-hangzhou.aliyuncs.com/eryajf/chrome-go-rod:base

LABEL maintainer eryajf

ADD care-screenshot /usr/bin

RUN chmod +x /usr/bin/care-screenshot