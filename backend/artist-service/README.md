# Artist Service

Manages artists and their association with EDM genres.

## Responsibilities

- Return artists for a given genre
- Return a single artist's detail along with their top songs

## API Endpoints (planned)

| Method | Path | Description |
|---|---|---|
| `GET` | `/artists` | List all artists |
| `GET` | `/artists/:id` | Get artist detail |
| `GET` | `/artists?genreId=<id>` | Get artists for a genre |

All response shapes are defined in `../../shared/contracts/artist.contracts.ts`.
