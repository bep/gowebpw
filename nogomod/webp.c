/*
 * Copyright 2025 Bjørn Erik Pedersen
 * SPDX-License-Identifier: MIT
 */

#include <stdlib.h>
#include <string.h>
#include <webp/encode.h>

static uint8_t *encodeNRGBA(WebPConfig *config, const uint8_t *rgba, int width, int height, int stride, size_t *output_size)
{
    WebPPicture pic;
    WebPMemoryWriter wrt;
    int ok;
    if (!WebPPictureInit(&pic))
    {
        return NULL;
    }
    pic.use_argb = 1;
    pic.width = width;
    pic.height = height;
    pic.writer = WebPMemoryWrite;
    pic.custom_ptr = &wrt;
    WebPMemoryWriterInit(&wrt);
    ok = WebPPictureImportRGBA(&pic, rgba, stride) && WebPEncode(config, &pic);
    WebPPictureFree(&pic);
    if (!ok)
    {
        WebPMemoryWriterClear(&wrt);
        return NULL;
    }
    *output_size = wrt.size;
    return wrt.mem;
}

static uint8_t *encodeGray(WebPConfig *config, uint8_t *y, int width, int height, int stride, size_t *output_size)
{
    WebPPicture pic;
    WebPMemoryWriter wrt;

    int ok;
    if (!WebPPictureInit(&pic))
    {
        return NULL;
    }

    pic.use_argb = 0;
    pic.width = width;
    pic.height = height;
    pic.y_stride = stride;
    pic.writer = WebPMemoryWrite;
    pic.custom_ptr = &wrt;
    WebPMemoryWriterInit(&wrt);

    const int uvWidth = (int)(((int64_t)width + 1) >> 1);
    const int uvHeight = (int)(((int64_t)height + 1) >> 1);
    const int uvStride = uvWidth;
    const int uvSize = uvStride * uvHeight;
    const int gray = 128;
    uint8_t *chroma;

    chroma = malloc(uvSize);
    if (!chroma)
    {
        return 0;
    }
    memset(chroma, gray, uvSize);

    pic.y = y;
    pic.u = chroma;
    pic.v = chroma;
    pic.uv_stride = uvStride;

    ok = WebPEncode(config, &pic);

    free(chroma);

    WebPPictureFree(&pic);
    if (!ok)
    {
        WebPMemoryWriterClear(&wrt);
        return NULL;
    }
    *output_size = wrt.size;
    return wrt.mem;
}