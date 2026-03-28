# Frontend

React application for the EDM Genre Explorer.

## Tech Stack

- **React** – UI framework
- **React Flow** (or similar) – Interactive genre graph visualization
- **Zustand** (or Context API) – State management
- **TypeScript** – Type safety, consuming contracts from `../shared`

## Structure (planned)

```
frontend/
├── public/             # Static assets
├── src/
│   ├── components/     # Reusable UI components
│   ├── pages/          # Top-level page views
│   ├── services/       # API client calls (typed against ../shared)
│   ├── store/          # State management
│   └── main.tsx        # App entry point
├── package.json
└── tsconfig.json
```

## Getting Started

```bash
npm install
npm run dev
```
