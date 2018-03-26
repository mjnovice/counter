
FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/counter /opt/bin/counter
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/counter"]
