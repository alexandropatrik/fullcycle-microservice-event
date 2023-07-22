FROM maven:3.6.3-jdk-11
WORKDIR /app
COPY . .
RUN mvn clean install -DskipTests
#CMD mvn spring-boot:run

#FROM openjdk:11

#COPY ./target/wallet-balance.jar /app/
#WORKDIR /app/

CMD ["java", "-jar", "target/wallet-balance.jar"]
#CMD ["tail", "-f", "/dev/null"]