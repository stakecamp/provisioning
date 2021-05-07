
# Root directory && run arguments
ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)

TAG_TESTNET := v1.1.54
TAG_TESTNET_CONFIG := T1.1.54.0

git-update-repo:
	@git submodule update --init --recursive

docker-build-testnet: git-update-repo
	@cd elrond-config-testnet && git pull origin master --tags && git checkout ${TAG_TESTNET_CONFIG}
	@echo "building testnet version ${TAG_TESTNET}"
	@docker build --build-arg VERSION=${TAG_TESTNET} -t stakecamp/elrdnode:t${TAG_TESTNET} .

docker-push-testnet: docker-build-testnet
	@docker push stakecamp/elrdnode:t${TAG_TESTNET}

docker-run-testnet: docker-build-testnet
	@docker run -p '8080:8080' -v ${ROOT_DIR}/data:/data -it stakecamp/elrdnode:t${TAG_TESTNET} $(RUN_ARGS)


TAG_MAINNET := v1.1.54
TAG_MAINNET_CONFIG := v1.1.54.0

docker-build-mainnet: git-update-repo
	@cd elrond-config-mainnet && git pull origin master --tags && git checkout ${TAG_MAINNET_CONFIG}
	@echo "building mainnet version ${TAG_MAINNET}"
	@docker build --build-arg VERSION=${TAG_MAINNET} -t stakecamp/elrdnode . 

docker-push-mainnet: docker-build-mainnet
	@docker tag stakecamp/elrdnode:latest stakecamp/elrdnode:${TAG_MAINNET} 
	@docker push stakecamp/elrdnode:${TAG_MAINNET} 
	@docker push stakecamp/elrdnode:latest

docker-run-mainnet: docker-build-mainnet
	@docker run -p '8080:8080' -v ${ROOT_DIR}/data:/data -it stakecamp/elrdnode:t${TAG_TESTNET} $(RUN_ARGS)

run: docker-run-mainnet
build: docker-build-mainnet
push: docker-push-mainnet
all: build
