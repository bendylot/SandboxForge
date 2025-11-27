PROJECT=sandboxforge

up:
	docker compose -p $(PROJECT) -f apps/backend/compose.yml -f apps/frontend/compose.yml --profile db --profile api --profile web up -d


down:
	docker compose -p $(PROJECT) down

api:
	cd C:/Users/bendylot/source/vscode/SandboxForge/apps/casino/backend && \
	swag init && \
	cd .. && \
	docker compose -f compose.yml up -d --build casino-api

web:
	docker compose -p sandboxforge -f apps/frontend/compose.yml up -d

logs:
	docker compose -p $(PROJECT) logs -f
dev-backend:
	cd apps\casino\web\src \
	npm run dev
