# go-project

## 環境構築

### Docker

```bash
docker compose up -d
```

### Go

```bash
# go-rest-api
GO_ENV=dev go run migrate/migrate.go
GO_ENV=dev go run main.go
```

### Next.js

```bash
# frontend
npm run dev
```
