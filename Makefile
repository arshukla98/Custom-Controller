
deps:
	export GO111MODULE=off && go get -v ./...

crd:
	kubectl create -f artifacts/crd.yaml

test:
	kubectl create -f artifacts/example.yaml

build:
	go build -o main .

run:
	./main --kubeconfig=${HOME}/.kube/config

get:
	echo 'Getting UpgradeKubes' && \
	kubectl get upgradekubes && \
	echo 'Getting Deployments, Services' && \
	kubectl get deploy,svc

desc:
	watch -n 0.1 'kubectl describe upgradekubes example-kube1'

cleanup:
	kubectl delete deploy nginx && kubectl delete svc nginx-svc && kubectl delete upgradekubes example-kube1
