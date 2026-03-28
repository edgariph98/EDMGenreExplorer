# API Gateway

Single entry point for all client requests. Routes traffic to the appropriate
backend service and (optionally) aggregates responses.

## Responsibilities

- Reverse-proxy `/genres/*` → genre-service
- Reverse-proxy `/songs/*` → song-service
- Reverse-proxy `/artists/*` → artist-service
- (Optional) Aggregate genre + songs + artists into a single response

## API Routing (planned)

| Prefix | Target Service |
|---|---|
| `/api/genres` | genre-service :3001 |
| `/api/songs` | song-service :3002 |
| `/api/artists` | artist-service :3003 |
