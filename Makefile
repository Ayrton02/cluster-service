build:
	docker build -t ayrtonborges2/cluster-service:$(shell date +'%F.%6N') .

push-tags:
	docker push -a ayrtonborges2/cluster-service	