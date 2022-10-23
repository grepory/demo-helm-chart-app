app:
	go build -v -o ./app

docker:
	docker build -t grepory/demo-helm-chart-app .
