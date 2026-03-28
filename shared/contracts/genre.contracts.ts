/**
 * API contracts for the Genre Service.
 *
 * These types define the exact request/response shapes for every genre endpoint.
 * Both the frontend API client and the backend controllers must conform to these.
 */

import { Genre, GenreGraph } from '../types';

// ─── Common wrapper ───────────────────────────────────────────────────────────

export interface ApiResponse<T> {
  data: T;
  /** Human-readable message, e.g. "ok" */
  message: string;
}

export interface ApiErrorResponse {
  error: string;
  statusCode: number;
}

// ─── GET /genres ─────────────────────────────────────────────────────────────

export type GetGenresResponse = ApiResponse<Genre[]>;

// ─── GET /genres/:id ─────────────────────────────────────────────────────────

export type GetGenreByIdResponse = ApiResponse<Genre>;

// ─── GET /genres/:id/related ─────────────────────────────────────────────────

export interface RelatedGenres {
  parent: Genre | null;
  children: Genre[];
}

export type GetRelatedGenresResponse = ApiResponse<RelatedGenres>;

// ─── GET /genres/graph ───────────────────────────────────────────────────────

export type GetGenreGraphResponse = ApiResponse<GenreGraph>;
