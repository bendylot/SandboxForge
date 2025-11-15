PROJECT=sandboxforge

up:
	docker compose -p $(PROJECT) -f apps/backend/compose.yml -f apps/frontend/compose.yml --profile db --profile api --profile web up -d


down:
	docker compose -p $(PROJECT) down

api:
	docker compose -p $(PROJECT) -f apps/backend/compose.yml --profile db --profile api up -d

web:
	docker compose -p sandboxforge -f apps/frontend/compose.yml up -d

logs:
	docker compose -p $(PROJECT) logs -f
