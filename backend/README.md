# Backend

Collection of microservices powering the EDM Genre Explorer API.

## Services

| Service | Port | Responsibility |
|---|---|---|
| `genre-service` | 3001 | Genre graph data, parent/child relationships |
| `song-service` | 3002 | Songs mapped to genres |
| `artist-service` | 3003 | Artists mapped to genres |
| `api-gateway` | 3000 | Single entry point, routes to services |

Each service is an independent Node.js application. All request/response shapes
are defined in `../shared/contracts` so the frontend and services always stay in sync.

## Getting Started

```bash
# From each service directory
npm install
npm run dev
```
