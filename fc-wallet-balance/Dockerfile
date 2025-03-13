FROM openjdk:17-jdk-slim

RUN apt-get update && apt-get install -y default-mysql-client && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY wait-for-mysql.sh /wait-for-mysql.sh

RUN chmod +x /wait-for-mysql.sh

COPY target/fc-wallet-balance-0.0.1-SNAPSHOT.jar java-service.jar

EXPOSE 3003

CMD ["/wait-for-mysql.sh", "mysql-java", "java", "-jar", "java-service.jar"]
