# API Gateway

Single entry point for all client requests. Routes traffic to the appropriate
backend service and (optionally) aggregates responses.

## Responsibilities

- Reverse-proxy `/genres/*` → genre-service
- Reverse-proxy `/songs/*` → song-service
- Reverse-proxy `/artists/*` → artist-service
- (Optional) Aggregate genre + songs + artists into a single response

## API Routing

| Prefix | Target Service |
|---|---|
| `/api/genres` | genre-service :3001 |
| `/api/songs` | song-service :3002 |
| `/api/artists` | artist-service :3003 |

## Configuration

Routes are defined in [`config.yaml`](./config.yaml). Each service entry has:

| Field | Required | Description |
|---|---|---|
| `name` | yes | Service identifier (used in logs/errors) |
| `prefix` | yes | Gateway path prefix; stripped before forwarding |
| `url` | yes | Upstream base URL |
| `methods` | no | Allowed HTTP methods (defaults to `default_methods`) |

Values support `${VAR}` and `${VAR:-default}` environment expansion, so
`PORT`, `GENRE_SERVICE_URL`, `SONG_SERVICE_URL` and `ARTIST_SERVICE_URL`
still override the defaults. Use a different file with `-config <path>` or
`GATEWAY_CONFIG=<path>`.

```bash
go run . -config config.yaml
```
