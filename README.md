# Stakecamp provisioning docker profiles

This repository we use to enhance deployment and provisioning of elrond nodes. We are using datadog for monitoring and therefore its part of docker compose.

Clone & Pull:

```bash
$ git clone https://github.com/stakecamp/provisioning.git
$ docker-compose up -d
```


Building:

```bash
$ git clone https://github.com/stakecamp/provisioning.git
$ git submodule update --init --recursive

$ docker-compose build
$ docker-compose up -d
```

## Environment

So in order to bootstrap in `/etc/enviroment` create two variables:

- `DD_API_KEY` - datadog api key
- `ELROND_MOUNT_VOLUME` - this will be path to data storage, recommended is backuped block storage volume

## Docker compose

Docker compose currently contains 4 images:

- `node` which is extended elrond docker image with config of
  given chain, you can supply arg CHAIN `--build-arg="mainnet|testnet"`.
- `datadog` is datadog agent, useful for log aggregation and 
- `autoheal` watching images for health
- `watchtower` for autoupdates


## ElrdKeep

Is small utility written by us handling healtchecks.
It's purpose is to find out public key of node and ask node for heartbeats. If last heartbeat is 10mins less then current time we mark container as unhealty.

This is useful for being sure node replies to network.


## Provisioning

Only thin needed is docker and docker-compose:

```bash
$ apt-get install -y docker.io docker-compose
$ systemctl enable docker # reboot persistency
```

## How to upgrade nodes

That is simply done by watchtower, only thing you need to handle is push latest image.
In our case that is done manually and we do revision before pushing new upgraded image.


## Redudndancy nodes

Remember that for every node there should be second one as a backup, that is somehwere else
ideailly in different datacenter.

