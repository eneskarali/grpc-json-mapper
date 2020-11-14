FROM envoyproxy/envoy:v1.14-latest

COPY ./proto/helloworld.proto /helloworld.proto

ADD ./config /etc/envoy

RUN apt-get update
RUN apt-get install curl unzip git -y
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.12.1/protoc-3.12.1-linux-x86_64.zip
RUN unzip protoc-3.12.1-linux-x86_64.zip -d $HOME/protocomp/
RUN rm protoc-3.12.1-linux-x86_64.zip

RUN git clone --depth 1 https://github.com/googleapis/googleapis

RUN mkdir desc

RUN $HOME/protocomp/bin/protoc -I/googleapis -I. --include_imports --include_source_info --descriptor_set_out=desc/helloworld.pb helloworld.proto
