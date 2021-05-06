FROM golang:1.15.7 as builder

RUN apt-get update

WORKDIR /go/elrond-go
COPY ./elrond-go .
RUN GO111MODULE=on go mod vendor

# Elrond node
WORKDIR /go/elrond-go/cmd/node
COPY ./stakecamp/logger.go /go/elrond-go/cmd/node/logger.go
RUN go build -i -v -ldflags="-X main.appVersion=$(git describe --tags --long --dirty)"
RUN cp /go/pkg/mod/github.com/!elrond!network/arwen-wasm-vm@$(cat /go/elrond-go/go.mod | grep arwen-wasm-vm | sed 's/.* //')/wasmer/libwasmer_linux_amd64.so /lib/libwasmer_linux_amd64.so

# Arwen node
WORKDIR /go/elrond-go
RUN go get github.com/ElrondNetwork/arwen-wasm-vm/cmd/arwen@$(cat /go/elrond-go/go.mod | grep arwen-wasm-vm | sed 's/.* //')
RUN go build -o ./arwen github.com/ElrondNetwork/arwen-wasm-vm/cmd/arwen
RUN cp /go/elrond-go/arwen /go/elrond-go/cmd/node/

# Keygen
WORKDIR /go/elrond-go/cmd/keygenerator
RUN go build

FROM ubuntu:18.04

COPY --from=builder "/go/elrond-go/cmd/node/node" "/usr/bin/elrdnode"
COPY --from=builder "/go/elrond-go/cmd/node/arwen" "/usr/bin/arwen"
COPY --from=builder "/go/elrond-go/cmd/keygenerator/keygenerator" "/usr/bin/elrdkeygen"
COPY --from=builder "/lib/libwasmer_linux_amd64.so" "/lib/libwasmer_linux_amd64.so"

ENV ARWEN_PATH /usr/bin/arwen

ARG CHAIN=mainnet

RUN apt-get -y update 
RUN apt-get install -y git 

COPY ./elrond-config-${CHAIN} /config
COPY ./stakecamp/prefs.toml /config/prefs.toml

RUN sed -i 's/\.\/config/\/config/' /config/genesisSmartContracts.json

RUN mkdir -p /data
VOLUME [ "/data" ]
WORKDIR /data

CMD ["elrdnode", \
    "--validator-key-pem-file", "/data/validatorKey.pem", \
    "--use-log-view", \
    "--use-health-service", \
    "--working-directory", "/data", \
    "--log-level", "*:DEBUG", \
    "--rest-api-interface", "0.0.0.0:8080", \
    "--genesis-file", "/config/genesis.json", \
    "--smart-contracts-file", "/config/genesisSmartContracts.json", \
    "--nodes-setup-file", "/config/nodesSetup.json", \
    "--config", "/config/config.toml", \
    "--config-api", "/config/api.toml", \
    "--config-economics", "/config/economics.toml", \
    "--config-systemSmartContracts", "/config/systemSmartContractsConfig.toml", \
    "--config-ratings", "/config/ratings.toml", \
    "--config-preferences", "/config/prefs.toml", \
    "--config-external", "/config/external.toml", \
    "--p2p-config", "/config/p2p.toml", \
    "--gas-costs-config", "/config/gasSchedules" \
]

EXPOSE 8080
HEALTHCHECK --start-period=60m --interval=5m --timeout=3s CMD curl -f http://0.0.0.0:8080/node/status || exit 1