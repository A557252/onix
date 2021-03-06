#
#    Onix Config Manager - Artie Artefact Registry Image
#    Copyright (c) 2018-2020 by www.gatblau.org
#    Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
#    Contributors to this project, hereby assign copyright in this code to the project,
#    to be licensed under the same terms as the rest of the code.
#
FROM registry.access.redhat.com/ubi8/ubi-minimal

LABEL author="gatblau"
LABEL maintainer="onix@gatblau.org"

ARG UNAME=artie
ARG APP_NAME=artie

ENV UID=1000
ENV GID=1000
ENV APP=$APP_NAME

RUN microdnf update --disablerepo=* --enablerepo=ubi-8-appstream --enablerepo=ubi-8-baseos -y && \
    microdnf install shadow-utils.x86_64 && \
    groupadd -g $GID -o $UNAME && \
    # -M create the user with no /home
    useradd -M -u $UID -g $GID $UNAME && \
    rm -rf /var/cache/yum && \
    microdnf clean all && \
    # pessimistic locking for metadata changes within a virtual repository
    mkdir -p /.artie/locks && chown $UNAME /.artie/locks && chmod ug+wr /.artie/locks && \
    # temprorary file operations folder
    mkdir -p /.artie/tmp && chown $UNAME /.artie/tmp && chmod ug+wr /.artie/tmp && \
    # web hooks configuration
    mkdir -p /.artie/hooks && chown $UNAME /.artie/hooks && chmod ug+wr /.artie/hooks

WORKDIR /app

COPY ./images/bin/output/$APP_NAME /app/$APP_NAME

USER $UNAME

CMD ["sh", "-c", "/app/${APP} serve"]

EXPOSE 8082/tcp
