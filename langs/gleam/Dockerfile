FROM rust:1.86.0-slim-bookworm AS builder

RUN apt-get update                   \
 && DEBIAN_FRONTEND='noninteractive' \
    apt-get install --yes build-essential curl ncurses-dev

ENV ERLVER=27.3.3 VER=1.10.0

RUN curl -#L https://github.com/erlang/otp/releases/download/OTP-$ERLVER/otp_src_$ERLVER.tar.gz | tar xz

WORKDIR /otp_src_$ERLVER

RUN ./configure            \
    --prefix=/usr          \
 && make -j`nproc` install \
 && strip /usr/bin/escript

RUN curl -#L https://github.com/gleam-lang/gleam/archive/refs/tags/v$VER.tar.gz \
  | tar xz --strip-components 1

RUN make -j`nproc` install

WORKDIR /gleam

RUN gleam new --name main --skip-git .

RUN gleam add argv \
 && gleam remove gleeunit

RUN chmod -R 777 .

COPY gleam.c /

RUN gcc -Wall -Werror -Wextra -o /usr/bin/gleam -s /gleam.c

FROM codegolf/lang-base

COPY --from=0 /bin/dash                                       /bin/sh
COPY --from=0 /gleam/gleam.toml /gleam/manifest.toml          /gleam/
COPY --from=0 /gleam/build/packages                           /gleam/build/packages
COPY --from=0 /gleam/src                                      /gleam/src
COPY --from=0 /lib/x86_64-linux-gnu/libc.so.6                 \
              /lib/x86_64-linux-gnu/libgcc_s.so.1             \
              /lib/x86_64-linux-gnu/libm.so.6                 \
              /lib/x86_64-linux-gnu/libstdc++.so.6            \
              /lib/x86_64-linux-gnu/libtinfo.so.6             /lib/
COPY --from=0 /lib64/ld-linux-x86-64.so.2                     /lib64/
COPY --from=0 /usr/bin/basename                               \
              /usr/bin/dirname                                \
              /usr/bin/erl                                    \
              /usr/bin/escript                                \
              /usr/bin/gleam                                  /usr/bin/
COPY --from=0 /usr/lib/erlang/bin/no_dot_erlang.boot          \
              /usr/lib/erlang/bin/start.boot                  /usr/lib/erlang/bin/
COPY --from=0 /usr/lib/erlang/erts-15.2.6/bin/beam.smp        \
              /usr/lib/erlang/erts-15.2.6/bin/erl_child_setup \
              /usr/lib/erlang/erts-15.2.6/bin/erlexec         \
              /usr/lib/erlang/erts-15.2.6/bin/inet_gethost    /usr/lib/erlang/erts-15.2.6/bin/
COPY --from=0 /usr/lib/erlang/lib/compiler-8.6.1/ebin         /usr/lib/erlang/lib/compiler-8.6.1/ebin
COPY --from=0 /usr/lib/erlang/lib/kernel-10.2.6/ebin          /usr/lib/erlang/lib/kernel-10.2.6/ebin
COPY --from=0 /usr/lib/erlang/lib/stdlib-6.2.2/ebin           /usr/lib/erlang/lib/stdlib-6.2.2/ebin
COPY --from=0 /usr/lib/locale                                 /usr/lib/locale
COPY --from=0 /usr/local/cargo/bin/gleam                      /usr/local/bin/

ENTRYPOINT ["gleam"]

CMD ["--version"]
