FROM golang:1.22-alpine3.20 as builder

RUN apk add \
    ca-certificates git ocaml gcc make automake autoconf pkgconfig m4 libtool wget \
    tar xz gettext-dev \
    patch \
    musl-dev \
    cmake \
    mingw-w64-gcc mingw-w64-binutils mingw-w64-headers mingw-w64-winpthreads

RUN mkdir -p /build/msgpack-c && \
    cd /build && \
    wget -O msgpack.tar.gz https://github.com/msgpack/msgpack-c/releases/download/c-6.1.0/msgpack-c-6.1.0.tar.gz && \
    cd /build/msgpack-c && \
    tar --strip-components=1 -xf ../msgpack.tar.gz && \
    mkdir build && \
    cd build && \
    CC=x86_64-w64-mingw32-gcc cmake .. -DCMAKE_SYSTEM_NAME=Windows -DCMAKE_BUILD_TYPE=RelWithDebInfo && \
    cmake --build . --config RelWithDebInfo && \
    cmake --install . --config RelWithDebInfo --prefix /build/sysroot

RUN mkdir -p /build && \
    cd /build && \
    git clone https://github.com/ebiggers/wimlib.git && \
    cd wimlib && \
    git checkout -f cd2a5e5d2e95c36e81d09077d06ad136f7d24950

RUN mkdir -p /build/sysroot && \
    cd /build/wimlib && \
    autoreconf -i && \
    ./configure --host=x86_64-w64-mingw32 --without-ntfs-3g --without-fuse && \
    make && \
    (make install DESTDIR=/build/sysroot || true)

RUN mkdir -p /build/src /build/dist
COPY . /build/src

WORKDIR /build/src
ENV CC=x86_64-w64-mingw32-gcc
ENV GOOS=windows
ENV CGO_ENABLED=1
ENV CFLAGS="-I/build/sysroot/include -I/build/wimlib/include"
RUN CGO_CFLAGS="-static -I/build/sysroot/include -I/build/wimlib/include" CGO_LDFLAGS="-L/build/sysroot/usr/local/lib -L/build/sysroot/lib /build/sysroot/lib/libmsgpack-c.a /build/sysroot/usr/local/lib/libwim.dll.a" go build -o /build/dist/go-wimlib-windows-dynamic.exe ./cmd/go-wimlib/
RUN CGO_CFLAGS="-static -I/build/sysroot/include -I/build/wimlib/include" CGO_LDFLAGS="-L/build/sysroot/usr/local/lib -L/build/sysroot/lib /build/sysroot/lib/libmsgpack-c.a /build/sysroot/usr/local/lib/libwim.a -lntdll" go build -o /build/dist/go-wimlib-windows-static.exe ./cmd/go-wimlib/

FROM scratch
COPY --from=builder /build/sysroot/usr/local/bin/libwim-15.dll /
COPY --from=builder /build/dist/ /
