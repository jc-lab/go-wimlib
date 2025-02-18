FROM golang:1.22-alpine3.20 as builder

RUN apk add \
    ca-certificates git ocaml gcc make automake autoconf pkgconfig m4 libtool \
    tar xz gettext-dev \
    patch \
    musl-dev \
    linux-headers \
    fuse3-dev \
    msgpack-c-dev

RUN mkdir -p /build && \
    cd /build && \
    git clone https://github.com/ebiggers/wimlib.git && \
    cd wimlib && \
    git checkout -f cd2a5e5d2e95c36e81d09077d06ad136f7d24950

COPY wimlib-patches/ /build/wimlib-patches/

RUN cd /build/wimlib && \
    find /build/wimlib-patches/ -type f -name "*.patch" | sort | while read name; do patch -p1 < $name; done && \
    autoreconf -i

ARG USE_NTFS_3G=true

RUN [ "x${USE_NTFS_3G:-}" != "xtrue" ] && echo "skip ntfs-3g" || apk add ntfs-3g-dev

RUN cd /build/wimlib && \
    CONFIGURE_ARGS="" && \
    [ "x${USE_NTFS_3G:-}" = "xtrue" ] || CONFIGURE_ARGS="${CONFIGURE_ARGS} --without-ntfs-3g" && \
    echo "CONFIGURE_ARGS: ${CONFIGURE_ARGS}" && \
    ./configure ${CONFIGURE_ARGS} && \
    make && \
    make install

RUN mkdir -p /build/src /build/dist
COPY . /build/src

RUN apk add fuse3-static ntfs-3g-static

WORKDIR /build/src
ENV CGO_ENABLED=1
RUN go build -o /build/dist/go-wimlib-linux-dynamic.exe ./cmd/go-wimlib/
RUN CGO_LDFLAGS="-static -lfuse3" && \
    [ "x${USE_NTFS_3G:-}" = "xtrue" ] && CGO_LDFLAGS="${CGO_LDFLAGS} -lntfs-3g" || true && \
    GO_LD_FLAGS="-linkmode external -extldflags '${CGO_LDFLAGS}'" && \
    go build --ldflags "${GO_LD_FLAGS}" -o /build/dist/go-wimlib-linux-static ./cmd/go-wimlib/

FROM scratch
COPY --from=builder /build/dist/ /
