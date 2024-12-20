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
# 	K8S locally - minikuke
# *********************

k8s-local:
	minikube start

k8s-open-tunnel:
	minikube tunnel

k8s-close-tunnel:
	minikube pause
	minikube stop

# *********************
# 			K8S
# *********************

k8s-apply:
	kubectl apply -f k8s/

k8s-get-ip-local:
	kubectl get services go-app-service


# *********************
# 	ECS
# *********************

ecs-register-task:
    aws ecs register-task-definition --cli-input-json file://ecs-task-definition.json

ecs-create-service:
    aws ecs create-service --cluster your-cluster-name --service-name go-service --task-definition go-service --desired-count 1 --launch-type FARGATE --network-configuration "awsvpcConfiguration={subnets=[subnet-xxxxxx],securityGroups=[sg-xxxxxx],assignPublicIp=ENABLED}"

ecs-update-service:
    aws ecs update-service --cluster your-cluster-name --service go-service --force-new-deployment

ecs-delete-service:
    aws ecs delete-service --cluster your-cluster-name --service go-service --force


# ******************************* others ******************
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