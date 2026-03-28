/**
 * API contracts for the Artist Service.
 *
 * These types define the exact request/response shapes for every artist endpoint.
 * Both the frontend API client and the backend controllers must conform to these.
 */

import { Artist } from '../types';
import { ApiResponse } from './genre.contracts';

// ─── GET /artists ────────────────────────────────────────────────────────────

export interface GetArtistsQuery {
  /** Filter artists by genre */
  genreId?: string;
}

export type GetArtistsResponse = ApiResponse<Artist[]>;

// ─── GET /artists/:id ────────────────────────────────────────────────────────

export type GetArtistByIdResponse = ApiResponse<Artist>;
