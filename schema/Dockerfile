FROM amazoncorretto:11

RUN yum install -y add wget gzip tar && mkdir /schema

WORKDIR /schema

RUN wget -qO- https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/5.2.4/flyway-commandline-5.2.4-linux-x64.tar.gz | tar xz && ln -s $(pwd)/flyway-5.2.4/flyway /usr/local/bin

RUN flyway -v

COPY migrate.sh .

CMD [ "/schema/migrate.sh" ]
