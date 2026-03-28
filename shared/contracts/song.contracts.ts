/**
 * API contracts for the Song Service.
 *
 * These types define the exact request/response shapes for every song endpoint.
 * Both the frontend API client and the backend controllers must conform to these.
 */

import { Song } from '../types';
import { ApiResponse } from './genre.contracts';

// ─── GET /songs ──────────────────────────────────────────────────────────────

export interface GetSongsQuery {
  /** Filter songs by genre */
  genreId?: string;
  /** Filter songs by artist */
  artistId?: string;
}

export type GetSongsResponse = ApiResponse<Song[]>;

// ─── GET /songs/:id ──────────────────────────────────────────────────────────

export type GetSongByIdResponse = ApiResponse<Song>;
