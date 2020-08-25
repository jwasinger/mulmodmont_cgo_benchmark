#pragma once

#define BIGINT_BITS 384
#define LIMB_BITS 64
#define LIMB_BITS_OVERFLOW 128
#include "bigint.h"
const uint64_t mod[] = {0xb9feffffffffaaab, 0x1eabfffeb153ffff, 0x6730d2a0f6b0f624, 0x64774b84f38512bf, 0x4b1ba7b6434bacd7, 0x1a0111ea397fe69a};
const uint64_t modinv = 0x89f3fffcfffcfffd;

void mulmodmont(UINT *const out, UINT *const x, const UINT *const y) {
    FUNCNAME(mulmodmontCIOS)(out, x, y, mod, modinv);
}
