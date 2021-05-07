# Stakecamp provisioning docker profiles

This repository we use to enhance deployment and provisioning of elrond nodes. We are using datadog for monitoring and therefore its part of docker compose.


```bash
$ git clone https://github.com/stakecamp/provisioning.git

# mainnet
$ make build # will build mainnet docker image
$ make run # will run mainnet in docker

# for testing
$ make docker-run-testnet 
```

## Environment

So in order to bootstrap in `/etc/enviroment` create two variables:

- `DD_API_KEY` - datadog api key
- `ELROND_MOUNT_VOLUME` - this will be path to data storage, recommended is backuped block storage volume

## Docker compose

Docker compose currently contains 4 images:

- `node` which is extended elrond docker image containing some imporvmenets, such as elrdkeep healtchecks and default configs.
- `datadog` is datadog agent, useful for log aggregation and monitoring
- `autoheal` watching images for health
- `watchtower` for autoupdates, will update based on docker hub and :latest tag.


## ElrdKeep

Is small utility written by us handling healtchecks.
It's purpose is to find out public key of node and ask node for heartbeats. If last heartbeat is 10mins less then current time we mark container as unhealty.

This is useful for being sure node replies to network.


## Provisioning

Only thing needed is docker and docker-compose:

```bash
$ apt-get install -y docker.io docker-compose
$ systemctl enable docker # reboot persistency
```

## How to upgrade nodes

That is simply done by watchtower, only thing you need to handle is push latest image.
In our case that is done manually and we do revision before pushing new upgraded image.


## Redudndancy nodes

Remember that for every node there should be second one as a backup, that is somehwere else
ideally different datacenter.

That being said you want to be carful with Redundancy level to not run same node twice with same level.
This could cause double singing.
