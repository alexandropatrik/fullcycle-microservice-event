FROM maven:3.6.3-jdk-11
WORKDIR /app
COPY . .
RUN mvn clean install -DskipTests

CMD ["java", "-jar", "target/wallet-balance.jar"]
#CMD ["tail", "-f", "/dev/null"]