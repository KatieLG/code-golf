FROM eclipse-temurin:22.0.1_8-jre-alpine AS builder

RUN apk add --no-cache build-base wget

RUN wget https://github.com/gaborsch/rocky/raw/bd2c1db/rocky.jar

COPY rockstar.c .

RUN gcc -Wall -Werror -Wextra -o rockstar -s rockstar.c

FROM codegolf/lang-base

COPY --from=0 /rocky.jar                 /
COPY --from=0 /lib/ld-musl-x86_64.so.1   /lib/
COPY --from=0 /opt/java/openjdk/bin/java /opt/java/openjdk/bin/
COPY --from=0 /opt/java/openjdk/lib      /opt/java/openjdk/lib
COPY --from=0 /rockstar                  /usr/bin/

ENTRYPOINT ["rockstar"]

CMD ["-h"]
