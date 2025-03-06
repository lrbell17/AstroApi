COMPOSE_FILE=docker/docker-compose.yml
COMPOSE_CMD=docker-compose -f $(COMPOSE_FILE)

start:
	$(COMPOSE_CMD) up -d

stop:
	$(COMPOSE_CMD) down

rebuild:
	$(COMPOSE_CMD) up -d --build

clean:
	$(COMPOSE_CMD) down -v

logs:
	$(COMPOSE_CMD) logs -f
