PROJECT="corda-apiservice"
IMG="corda-apiservice:latest"

default:
	echo ${PROJECT}

docker-build:
	docker build . -t ${IMG}

run: 
	docker run -it --rm \
		--name ${PROJECT} \
	 	-p 8080:8080 \
		${IMG}

ssh:
	docker exec -it ${PROJECT} bash