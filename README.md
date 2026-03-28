# 🎧 EDM Genre Explorer App
## 📌 Feature Requirements

### 1. Core Goal
Enable users to:
- Explore genres and subgenres visually
- Understand how genres are derived from others
- Discover songs and artists tied to each genre
- Navigate deeply nested subgenres

---

### 2. Core Features

#### A. Interactive Genre Graph
- Visual graph of genres and subgenres
- Nodes represent genres
- Edges represent relationships (derived from, influenced by)
- Zoom, pan, and drag functionality
- Expand/collapse subgenre branches
- Click node to view details

#### B. Genre Detail View
- Genre name, description, BPM range
- Parent and child genres
- Tags/characteristics
- Associated songs
- Associated artists

#### C. Song Discovery
- List songs for selected genre
- Preview playback (via external links)
- Filter songs by BPM or energy (optional later)

#### D. Artist View
- Artists associated with a genre
- Basic artist info
- List of top songs

#### E. Search
- Search for genres, songs, and artists
- Autocomplete suggestions
- Jump to node in graph

#### F. Navigation
- Breadcrumb navigation (parent → current → child)
- “Related genres” suggestions

---

### 3. Visualization Requirements
- Node size reflects popularity or importance
- Node color represents core genre family
- Edge styles:
  - Solid = derived from
  - Dashed = influenced by
- Layout:
  - Root genres at center
  - Subgenres expand outward

---

### 4. Data Requirements
- Genres with relationships (graph structure)
- Songs mapped to genres
- Artists mapped to genres
- Initial dataset:
  - 30–50 genres
  - 2–3 hierarchy levels
  - 5–10 songs per genre

---

## 🚀 MVP Requirements

### 🎯 MVP Goal
Deliver a working product that allows users to:
- Visualize genre relationships
- Click into a genre
- View associated songs and artists

---

### ✅ MVP Features (Must Have)

#### 1. Genre Graph (Basic)
- Render a simple graph (nodes + edges)
- Clickable nodes
- Basic zoom/pan

#### 2. Genre Service Integration
- Fetch:
  - Genre details
  - Parent/child relationships

#### 3. Genre Detail Panel
- Show:
  - Name
  - Description
  - Parent + child genres

#### 4. Song Integration
- Display songs for selected genre
- Show:
  - Title
  - Artist
  - External link (Spotify/SoundCloud)

#### 5. Artist Integration
- Display artists related to genre
- Basic info only

---

### ❌ Not Included in MVP
- AI-generated genres
- Advanced filtering (BPM sliders, mood)
- User accounts / authentication
- Playlists or favorites
- Complex recommendation systems

---

### 🛠️ MVP Tech Scope

#### Backend
- Genre Service (graph-based)
- Song Service
- Artist Service
- Simple API Gateway (optional for MVP)

#### Frontend
- React app
- Graph visualization (React Flow or similar)
- Basic state management (Zustand or Context)

---

### 📈 MVP Success Criteria
- User can explore at least 2–3 levels of genre hierarchy
- Clicking a node updates the detail panel
- Songs and artists load correctly for a genre
- Graph interactions are smooth and responsive

---

## 🔥 Key Differentiator
Unlike traditional music apps, this app:
> Visualizes music as a connected graph instead of a list

This enables deeper discovery and understanding of how genres evolve.
