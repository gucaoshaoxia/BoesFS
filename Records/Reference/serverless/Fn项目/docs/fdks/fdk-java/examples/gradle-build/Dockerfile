FROM gradle:5.6.2-jdk11 as build-stage

WORKDIR /function
# needed for gradle?
USER root
ENV GRADLE_USER_HOME /function/.gradle

# Code build
# Copy any build files into the build container and build
COPY *.gradle /function/
RUN ["gradle", "-s", "--no-daemon","--console","plain","cacheDeps"]

# Copies build source into build container
COPY src /function/src

RUN ["gradle", "-s", "--no-daemon","--console","plain","build"]
# Container build
FROM fnproject/fn-java-fdk:jre11-1.0.98
WORKDIR /function
COPY --from=build-stage /function/build/libs/*.jar /function/build/deps/*.jar /function/app/
CMD ["com.example.fn.HelloFunction::handleRequest"]
