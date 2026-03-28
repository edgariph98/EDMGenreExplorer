# Song Service

Manages songs and their association with EDM genres.

## Responsibilities

- Return songs for a given genre
- Return a single song's detail

## API Endpoints (planned)

| Method | Path | Description |
|---|---|---|
| `GET` | `/songs` | List all songs |
| `GET` | `/songs/:id` | Get song detail |
| `GET` | `/songs?genreId=<id>` | Get songs for a genre |

All response shapes are defined in `../../shared/contracts/song.contracts.ts`.
