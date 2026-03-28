# Shared

TypeScript types and API contracts shared between the **frontend** and every **backend service**.

## Why a shared folder?

A single source of truth for all data shapes means:
- Changing an API response in the backend automatically surfaces type errors in the frontend.
- No duplication of interface definitions across packages.

## Structure

```
shared/
├── types/
│   └── index.ts          # Core domain types (Genre, Song, Artist, GenreGraph…)
├── contracts/
│   ├── genre.contracts.ts    # Request/response shapes for Genre Service
│   ├── song.contracts.ts     # Request/response shapes for Song Service
│   └── artist.contracts.ts   # Request/response shapes for Artist Service
└── index.ts              # Barrel export – import from here
```

## Usage

```ts
// In the frontend or any backend service:
import { Genre, GetGenreGraphResponse } from '../shared';
```

> When the monorepo is set up with workspaces you can alias this as `@edm/shared`.
