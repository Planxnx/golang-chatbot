APP?=golang-chatbot
RELEASE?=1
PORT?=8080
GOOS?=linux
GOARCH?=amd64
CONTAINER_IMAGE?=docker.io/webdeva/${APP}

# clean:
# 	rm -f ${APP}

# build: clean
# 	go build -o ${APP}

# run: build
# 	PORT=${PORT} ./${APP}

container: 
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

run: container
	docker stop $(APP):$(RELEASE) || true && docker rm $(APP):$(RELEASE) || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		$(APP):$(RELEASE)

push: container
	docker push $(CONTAINER_IMAGE):$(RELEASE)
