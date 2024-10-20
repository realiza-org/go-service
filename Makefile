# *********************
# 			DOCKER
# *********************

up:
	docker-compose up -d --build 

down:
	docker-compose down 

dockerhub-publish:
	docker build -t joaowillamy/go-service:latest .
	docker push joaowillamy/go-service:latest

# *********************
# 			K8S
# *********************

k8s-local:
	minikube start

k8s-apply:
	kubectl apply -f k8s/

k8s-open-tunnel:
	minikube tunnel

k8s-get-ip-local:
	kubectl get services go-app-service

k8s-close-tunnel:
	minikube pause
	minikube stop

# ******************************* others, jsut learning ******************
k8s-get-contexts:
	kubectl config get-contexts

k8s-get-users:
	kubectl config get-users

k8s-create-context:
	kubectl config set-context context-realiza-org \
		--cluster=cluster-realiza-org \
		--user=willamy.sousa \
		--namespace=namespace-realiza-org

k8s-use-context:
	kubectl config use-context context-realiza-org