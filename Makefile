init:
	mkdir -p ~/go/bin ~/go/pkg ~/go/src/github.com/arshukla98/sample-controller && \
	cp -r . ~/go/src/github.com/arshukla98/sample-controller && \
	echo "Run -> cd ~/go/src/github.com/arshukla98/sample-controller"

env:
	export GOPATH=~/go && \
	export GO111MODULE=off

deps:
	go get -v ./...

create_crd:
	kubectl create -f artifacts/crd.yaml

build:
	go build -o main .

run:
	./main --kubeconfig=${HOME}/.kube/config

get_cr_yaml:
	cat artifacts/example.yaml

get_all:
	echo 'Getting DeployServices' && \
	kubectl get deployservices && \
	echo 'Getting Deployments, Services' && \
	kubectl get deploy,svc

create_cr:
	kubectl create -f artifacts/example.yaml

desc:
	watch -n 0.1 'kubectl describe deployservices example-kube1'

cleanup:
	kubectl delete deploy nginx && kubectl delete svc nginx-svc && kubectl delete deployservices example-kube1
