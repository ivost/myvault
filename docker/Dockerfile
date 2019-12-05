FROM alpine:3.10 as runtime

RUN apk --no-cache --update add dumb-init
EXPOSE 8080 52052

WORKDIR /
COPY ./docker-entrypoint.sh .
COPY ./lbuild/server .
COPY ./lbuild/client .
COPY ./config.yaml .
# COPY ssl/ ./ssl/
#RUN ls -altrh *.sh
ENTRYPOINT ["/server"]
#ENTRYPOINT ["/usr/bin/dumb-init","--"]
#CMD ["/docker-entrypoint.sh","server"]
#CMD ["/bin/sh", "-c", "while true; do sleep 10; echo $(date)  -  sleeping ...; done"]
