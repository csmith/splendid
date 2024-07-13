FROM gradle:jdk22 AS build

ENV LANG en_GB.UTF-8
ENV LANGUAGE en_GB:en
ENV LC_ALL en_GB.UTF-8
ENV JAVA_TOOL_OPTIONS -Dfile.encoding=UTF8

COPY . /home/gradle/splendid/
WORKDIR /home/gradle/splendid/

RUN gradle :composeApp:wasmJsBrowserDistribution :server:buildFatJar

FROM eclipse-temurin:22 AS run

COPY --from=build /home/gradle/splendid/server/build/libs/server-all.jar /server.jar
COPY --from=build /home/gradle/splendid/composeApp/build/dist/wasmJs/productionExecutable /static

WORKDIR /
CMD java -jar /server.jar