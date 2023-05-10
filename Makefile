build:
	docker build -t ayrtonborges2/cluster-service:latest .

push-tags:
	docker push -a ayrtonborges2/cluster-service	