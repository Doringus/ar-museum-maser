# AR museum backend
## Prerequirements

- JDK 17
- Maven
- Docker

## Build 
`mvn  package`

## Before usage

- Run mongodb docker image `docker run -d -p 27017:27017 --name mongo mongo:latest`
- Run museum backend `java -jar target/armaster-0.0.1-SNAPSHOT.jar`
- Navigate to `http://localhost:8080/mock_post?path=`, where path - path to Archive.zip(in repo root)

## Usage

- Navigate to `http://localhost:8080/main_qr`
- Scan qr code via ArMuseum app