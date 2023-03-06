IMAGE_NAME=medvedewem/pow_example
CONTAINER_NAME=pow_example
BRIDGE_NAME=bridge_net

.PHONY: build run

build:
	docker build -t ${IMAGE_NAME} .

run:
	docker run --rm --name ${CONTAINER_NAME} --net ${BRIDGE_NAME} ${IMAGE_NAME}