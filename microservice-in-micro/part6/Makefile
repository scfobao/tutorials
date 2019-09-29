.PHONY: network
network:
	docker network create --driver overlay my-network

.PHONY: clear
clear:
	docker service rm $(docker service ls -q)
	docker network rm my-network


