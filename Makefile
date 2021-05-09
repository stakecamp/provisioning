
# Root directory && run arguments
ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)

TAG_MAINNET := v1.1.54
TAG_MAINNET_CONFIG := v1.1.54.0

git-update-repo:
	@git submodule update --init --recursive

docker-build-mainnet: git-update-repo
	@-cd elrond-go && git pull origin master --tags && git checkout -q ${TAG_MAINNET} && git describe --tags --long --dirty
	@-cd elrond-config-mainnet && git pull origin master --tags && git checkout -q ${TAG_MAINNET_CONFIG} && git describe --tags --long --dirty
	@echo "building mainnet version ${TAG_MAINNET}"
	@docker build --progress plain --build-arg VERSION=${TAG_MAINNET} -t stakecamp/elrdnode:${TAG_MAINNET} . 
	@cd elrond-config-mainnet && git checkout master
	@cd elrond-go && git checkout master
	@echo "\n\nContainer stakecamp/elrdnode:${TAG_MAINNET}"
	@echo "Container stakecamp/elrdnode:latest"

docker-push-mainnet: docker-build-mainnet
	@docker tag stakecamp/elrdnode:${TAG_MAINNET} stakecamp/elrdnode:latest 
	@docker push stakecamp/elrdnode:${TAG_MAINNET} 
	@docker push stakecamp/elrdnode:latest

docker-run-mainnet: docker-build-mainnet
	@docker run -p '8080:8080' -v ${ROOT_DIR}/data:/data -it stakecamp/elrdnode:${TAG_MAINNET} $(RUN_ARGS)

run: docker-run-mainnet
build: docker-build-mainnet
push: docker-push-mainnet
all: build
