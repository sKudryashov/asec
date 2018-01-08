build:
	docker-compose -f docker/docker-compose.yml build
up:
	docker-compose -f docker/docker-compose.yml up --force-recreate
restart:
	docker-compose restart
stop:
	docker-compose -f docker/docker-compose.yml stop
test:
	docker exec -i docker_fileserver_1 /bin/sh test.sh
cmdserver:
	docker exec -ti docker_fileserver_1 /bin/sh
cmdminer:
	docker exec -ti docker_fileminer_1 /bin/sh
cmdreader:
	docker exec -ti docker_fileinforeader_1 /bin/sh
initdb:
	cat migrations/dump.sql | docker exec -i docker_db_1 psql -U postgres --dbname=asc