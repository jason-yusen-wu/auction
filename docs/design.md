# Design Document 

## Project Description
An extensible, scalable platform for high-contention eCommerce

The goal is to handle a wide variety of workloads and the systems challenges they pose:
* General Retail: a large product catalog, e.g. Amazon as you know it 
* Limited inventory: prevent double booking and maintain idempotency of transactions
* Viral flash sales: scaling to massive traffic spikes when your sale goes live 
* Auctions: real-time price updates with sensitive deadlines

The platform also comes with business/operational capabilities:
* Fraud detection: anti-bot systems and payment idempotency
* Efficient catalog: search and recommendations 
* Resilience: fault-tolerant and observable
* Security: meeting compliance

## Project Timeline 

### *v0.1*
* Project is containerized, deployed, and hooked up to CI pipeline 

## System Design
[Excalidraw diagram](https://excalidraw.com/#json=zwttAOUhkksmz74yneHcI,m9l2OLFWReFZYujy2GNmXQ)

