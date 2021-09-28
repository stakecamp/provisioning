ARG VERSION=v1.2.22


FROM golang:1.15.7 as builder

ARG VERSION

RUN apt-get update && apt-get install -y

WORKDIR /go/elrond-go
COPY ./elrond-go .
RUN GO111MODULE=on go mod vendor

WORKDIR /go/elrond-go/cmd/node
RUN go build -i -v -ldflags="-X main.appVersion=${VERSION}"
RUN cp /go/pkg/mod/github.com/!elrond!network/arwen-wasm-vm@$(cat /go/elrond-go/go.mod | grep arwen-wasm-vm | sed 's/.* //' | tail -n 1)/wasmer/libwasmer_linux_amd64.so /lib/libwasmer_linux_amd64.so


FROM golang:1.16 as elrdkeep

RUN mkdir -p /go/stakecamp/elrdkeep
WORKDIR /go/stakecamp/elrdkeep

COPY ./stakecamp/elrdkeep .
RUN go build


FROM ubuntu:18.04

ARG VERSION

RUN echo "Building container at version ${VERSION}"

COPY --from=builder "/go/elrond-go/cmd/node/node" "/usr/bin/elrdnode"
COPY --from=elrdkeep "/go/stakecamp/elrdkeep/elrdkeep" "/usr/bin/elrdkeep"
COPY --from=builder "/lib/libwasmer_linux_amd64.so" "/lib/libwasmer_linux_amd64.so"


RUN apt-get -y update 
RUN apt-get install -y git

COPY ./elrond-config-mainnet /config
RUN sed -i 's/\.\/config/\/config/' /config/genesisSmartContracts.json

RUN mkdir -p /data
VOLUME [ "/data" ]
WORKDIR /data

CMD ["elrdnode", \
    "--validator-key-pem-file", "/data/validatorKey.pem", \
    "--config-preferences", "/data/prefs.toml", \
    "--disable-ansi-color", \
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
    "--config-external", "/config/external.toml", \
    "--p2p-config", "/config/p2p.toml", \
    "--epoch-config", "/config/enableEpochs.toml", \
    "--gas-costs-config", "/config/gasSchedules" \
]

HEALTHCHECK --start-period=30s --interval=2m --timeout=10s --retries=15 CMD elrdkeep --host="0.0.0.0:8080"
EXPOSE 8080