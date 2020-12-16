PROJECT="corda-apiservice"
IMG="dockerxucheng/corda-apiservice:latest"

default:
	echo ${PROJECT}

docker-build:
	docker build . -t ${IMG}

docker-push:
	docker push ${IMG}

run: 
	docker run -it --rm \
		--name ${PROJECT} \
	 	-p 8080:8080 \
		${IMG}

ssh:
	docker exec -it ${PROJECT} bash

deploy: 
	kubectl apply -f ./deployment/k8s/

delete:
	kubectl delete -f ./deployment/k8s/