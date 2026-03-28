# Genre Service

Manages the EDM genre graph: nodes (genres) and edges (relationships between genres).

## Responsibilities

- Return all genres with their parent/child relationships
- Return a single genre's full detail
- Provide graph-ready data (nodes + edges)

## API Endpoints (planned)

| Method | Path | Description |
|---|---|---|
| `GET` | `/genres` | List all genres |
| `GET` | `/genres/:id` | Get genre detail |
| `GET` | `/genres/:id/related` | Get parent + child genres |
| `GET` | `/genres/graph` | Get full graph (nodes + edges) |

All response shapes are defined in `../../shared/contracts/genre.contracts.ts`.
