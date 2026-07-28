# Technology Selection Rationale 

## Real time communication
Go's concurrency is ideal for handling websocket fanouts with tools like Melody/Centrifugo

## Database
Postgres is the most flexible choice

## Plausible technologies 
> Writing them down here for now, can decide whether they're useful later 
* API Styles: REST vs JSON APIs vs gRPC vs SOAP vs GraphQL
* Caching: Redis vs Memcached
* Web servers: Nginx
* Message brokers: Kafka vs RabbitMQ
* Architectural Patterns: Monolith vs Microservices vs SOA vs Serverless vs Service Mesh vs ... 😵‍💫
* Search Engines: Elasticsearch
* RT Data: Server-sent Events, WebSockets/WebTransport, Polling
* NoSQL: real time, document, KV, column, graph, time series 
* Observability: ??

