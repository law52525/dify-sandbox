FROM python:3.10-slim

# copy main binary to /main
COPY main /main
COPY conf/config.yaml /conf/config.yaml

# change source to TSINGHUA if environment['TSINGHUA'] is set
ARG TSINGHUA

RUN apt-get clean && \
    apt-get update && apt-get install -y pkg-config libseccomp-dev \
    && rm -rf /var/lib/apt/lists/* \
    && chmod +x /main \
    && pip3 install jinja2 \
    && wget -O /opt/node-v20.11.1-linux-x64.tar.xz https://nodejs.org/dist/v20.11.1/node-v20.11.1-linux-x64.tar.xz \
    && tar -xvf /opt/node-v20.11.1-linux-x64.tar.xz -C /opt \
    && ln -s /opt/node-v20.11.1-linux-x64/bin/node /usr/local/bin/node \
    && rm -f /opt/node-v20.11.1-linux-x64.tar.xz

ENTRYPOINT ["/main"]