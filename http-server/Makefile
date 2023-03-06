IMAGE_NAME=medvedewem/pow_http_server
CONTAINER_NAME=pow_http_server
BRIDGE_NAME=bridge_net

.PHONY: build run

build:
	docker build -t ${IMAGE_NAME} .

run:
	docker network inspect bridge_net >/dev/null 2>&1 || docker network create bridge_net
	docker run --rm --name ${CONTAINER_NAME} --net ${BRIDGE_NAME} ${IMAGE_NAME}