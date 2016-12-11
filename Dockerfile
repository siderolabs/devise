FROM alpine:latest

ADD devise /devise
RUN chmod +x /devise

ONBUILD ADD plan.yml /plan.yml
ONBUILD ADD ./templates /templates

ENTRYPOINT ["/devise"]
