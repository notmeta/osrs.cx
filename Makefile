TOKEN=`cat .local_token`
DOCKER_CONTAINER_NAME="cx.store"


docker_run:
	docker run -p 6379:6379 -v /docker/redis-data:/data \
      --name $(DOCKER_CONTAINER_NAME) -d redis \
      redis-server --appendonly yes

docker_kill:
	docker kill $(DOCKER_CONTAINER_NAME)

docker_rm:
	docker rm $(DOCKER_CONTAINER_NAME)

pull:
	git pull

kill:
	-kill `pgrep osrs.cx`

fresh_deploy:
	make kill
	make pull
	make build
	-make docker_run
	make run

build:
	go build

run:
	nohup ./osrs.cx -t "$(TOKEN)" &
