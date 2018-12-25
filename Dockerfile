FROM scratch
LABEL maintainer="latschesar.atanassov@gmx.at"

ADD usersvc usersvc
EXPOSE 8080
ENTRYPOINT ["/usersvc"]