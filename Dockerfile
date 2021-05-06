FROM golang:1.16 as elrdkeep

RUN mkdir -p /go/stakecamp/elrdkeep
WORKDIR /go/stakecamp/elrdkeep

COPY ./stakecamp/elrdkeep .
RUN go build

FROM elrondnetwork/elrond-go-node:v1.1.54 as builder
FROM ubuntu:18.04

COPY --from=builder "/go/elrond-go/cmd/node/node" "/usr/bin/elrdnode"
COPY --from=builder "/go/elrond-go/cmd/node/arwen" "/usr/bin/arwen"
COPY --from=elrdkeep "/go/stakecamp/elrdkeep/elrdkeep" "/usr/bin/elrdkeep"
COPY --from=builder "/lib/libwasmer_linux_amd64.so" "/lib/libwasmer_linux_amd64.so"

ENV ARWEN_PATH /usr/bin/arwen

ARG CHAIN=mainnet

RUN apt-get -y update 
RUN apt-get install -y git

COPY ./elrond-config-${CHAIN} /config
RUN sed -i 's/\.\/config/\/config/' /config/genesisSmartContracts.json

RUN mkdir -p /data
VOLUME [ "/data" ]
WORKDIR /data

# "--use-log-view", \

CMD ["elrdnode", \
    "--validator-key-pem-file", "/data/validatorKey.pem", \
    "--config-preferences", "/data/prefs.toml", \
    "--disable-ansi-color", \
    "--use-health-service", \
    "--working-directory", "/data", \
    "--log-level", "*:INFO", \
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
    "--gas-costs-config", "/config/gasSchedules" \
]

HEALTHCHECK --start-period=30s --interval=2m --timeout=10s --retries=15 CMD elrdkeep --host="0.0.0.0:8080"
EXPOSE 8080