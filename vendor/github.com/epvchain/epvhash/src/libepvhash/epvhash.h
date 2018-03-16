/*
  This file is part of epvhash.

  epvhash is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  epvhash is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with epvhash.  If not, see <http://www.gnu.org/licenses/>.
*/

/** @file epvhash.h
* @date 2015
*/
#pragma once

#include <stdint.h>
#include <stdbool.h>
#include <string.h>
#include <stddef.h>
#include "compiler.h"

#define EPVHASH_REVISION 23
#define EPVHASH_DATASET_BYTES_INIT 1073741824U // 2**30
#define EPVHASH_DATASET_BYTES_GROWTH 8388608U  // 2**23
#define EPVHASH_CACHE_BYTES_INIT 1073741824U // 2**24
#define EPVHASH_CACHE_BYTES_GROWTH 131072U  // 2**17
#define EPVHASH_EPOCH_LENGTH 30000U
#define EPVHASH_MIX_BYTES 128
#define EPVHASH_HASH_BYTES 64
#define EPVHASH_DATASET_PARENTS 256
#define EPVHASH_CACHE_ROUNDS 3
#define EPVHASH_ACCESSES 64
#define EPVHASH_DAG_MAGIC_NUM_SIZE 8
#define EPVHASH_DAG_MAGIC_NUM 0xFEE1DEADBADDCAFE

#ifdef __cplusplus
extern "C" {
#endif

/// Type of a seedhash/blockhash e.t.c.
typedef struct epvhash_h256 { uint8_t b[32]; } epvhash_h256_t;

// convenience macro to statically initialize an h256_t
// usage:
// epvhash_h256_t a = epvhash_h256_static_init(1, 2, 3, ... )
// have to provide all 32 values. If you don't provide all the rest
// will simply be unitialized (not guranteed to be 0)
#define epvhash_h256_static_init(...)			\
	{ {__VA_ARGS__} }

struct epvhash_light;
typedef struct epvhash_light* epvhash_light_t;
struct epvhash_full;
typedef struct epvhash_full* epvhash_full_t;
typedef int(*epvhash_callback_t)(unsigned);

typedef struct epvhash_return_value {
	epvhash_h256_t result;
	epvhash_h256_t mix_hash;
	bool success;
} epvhash_return_value_t;

/**
 * Allocate and initialize a new epvhash_light handler
 *
 * @param block_number   The block number for which to create the handler
 * @return               Newly allocated epvhash_light handler or NULL in case of
 *                       ERRNOMEM or invalid parameters used for @ref epvhash_compute_cache_nodes()
 */
epvhash_light_t epvhash_light_new(uint64_t block_number);
/**
 * Frees a previously allocated epvhash_light handler
 * @param light        The light handler to free
 */
void epvhash_light_delete(epvhash_light_t light);
/**
 * Calculate the light client data
 *
 * @param light          The light client handler
 * @param header_hash    The header hash to pack into the mix
 * @param nonce          The nonce to pack into the mix
 * @return               an object of epvhash_return_value_t holding the return values
 */
epvhash_return_value_t epvhash_light_compute(
	epvhash_light_t light,
	epvhash_h256_t const header_hash,
	uint64_t nonce
);

/**
 * Allocate and initialize a new epvhash_full handler
 *
 * @param light         The light handler containing the cache.
 * @param callback      A callback function with signature of @ref epvhash_callback_t
 *                      It accepts an unsigned with which a progress of DAG calculation
 *                      can be displayed. If all goes well the callback should return 0.
 *                      If a non-zero value is returned then DAG generation will stop.
 *                      Be advised. A progress value of 100 means that DAG creation is
 *                      almost complete and that this function will soon return succesfully.
 *                      It does not mean that the function has already had a succesfull return.
 * @return              Newly allocated epvhash_full handler or NULL in case of
 *                      ERRNOMEM or invalid parameters used for @ref epvhash_compute_full_data()
 */
epvhash_full_t epvhash_full_new(epvhash_light_t light, epvhash_callback_t callback);

/**
 * Frees a previously allocated epvhash_full handler
 * @param full    The light handler to free
 */
void epvhash_full_delete(epvhash_full_t full);
/**
 * Calculate the full client data
 *
 * @param full           The full client handler
 * @param header_hash    The header hash to pack into the mix
 * @param nonce          The nonce to pack into the mix
 * @return               An object of epvhash_return_value to hold the return value
 */
epvhash_return_value_t epvhash_full_compute(
	epvhash_full_t full,
	epvhash_h256_t const header_hash,
	uint64_t nonce
);
/**
 * Get a pointer to the full DAG data
 */
void const* epvhash_full_dag(epvhash_full_t full);
/**
 * Get the size of the DAG data
 */
uint64_t epvhash_full_dag_size(epvhash_full_t full);

/**
 * Calculate the seedhash for a given block number
 */
epvhash_h256_t epvhash_get_seedhash(uint64_t block_number);

#ifdef __cplusplus
}
#endif
