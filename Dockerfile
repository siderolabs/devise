FROM alpine:latest

RUN apk --update add curl tar && rm -f /var/cache/apk/*
RUN curl -L https://github.com/autonomy/devise/releases/download/v0.0.3/devise-linux-amd64.tar.gz | tar -xz -C /bin

ONBUILD ADD plan.yml /plan.yml
ONBUILD ADD ./templates /templates

ENTRYPOINT ["devise"]
