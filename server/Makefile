GENERIC := $(wordlist 2,$(words $(MAKECMDGOALS)), $(MAKECMDGOALS))
$(eval $(GENERIC):;@:)

file = docker/docker-compose.dev.yaml
ifeq (${CONTEXT}, production)
		file = docker/docker-compose.yaml
endif

project = discord-clone
cc = docker compose -p $(project) -f $(file)
ex = docker exec -it discord-server

.PHONY: *
.DEFAULT_GOAL := detach

build:
		$(cc) build $(GENERIC)
up:
		$(cc) up $(GENERIC)
detach:
		$(cc) up -d $(GENERIC)
down:
		$(cc) down $(GENERIC)
run:
		make build && make up
debug:
		export BUILD_TARGET=debug && make run
dev:
		export BUILD_TARGET=dev && make run
stop:
		$(cc) stop $(GENERIC)
console:
		$(ex) /bin/sh
