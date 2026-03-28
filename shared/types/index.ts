/**
 * Core domain types shared between the frontend and all backend services.
 * Import these wherever you need plain data shapes (no HTTP concerns).
 */

// ─── Genre ───────────────────────────────────────────────────────────────────

export interface Genre {
  id: string;
  name: string;
  description: string;
  bpmMin?: number;
  bpmMax?: number;
  tags: string[];
  /** ID of the parent genre, or null for root genres */
  parentId: string | null;
  /** IDs of direct child genres */
  childIds: string[];
}

// ─── Graph ───────────────────────────────────────────────────────────────────

export type EdgeType = 'derived_from' | 'influenced_by';

export interface GenreNode {
  id: string;
  label: string;
  /** Core genre family used for colour-coding */
  family: string;
  /** Relative popularity (0–1), used to scale node size */
  popularity: number;
}

export interface GenreEdge {
  id: string;
  source: string;
  target: string;
  type: EdgeType;
}

export interface GenreGraph {
  nodes: GenreNode[];
  edges: GenreEdge[];
}

// ─── Song ────────────────────────────────────────────────────────────────────

export interface Song {
  id: string;
  title: string;
  artistId: string;
  genreId: string;
  bpm?: number;
  energy?: number;
  /** External preview URL (Spotify, SoundCloud, etc.) */
  externalUrl?: string;
}

// ─── Artist ──────────────────────────────────────────────────────────────────

export interface Artist {
  id: string;
  name: string;
  genreIds: string[];
  bio?: string;
  /** IDs of the artist's top songs */
  topSongIds: string[];
}
