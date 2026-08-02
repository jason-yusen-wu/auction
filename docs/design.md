# Design Document 

An eCommerce backend designed for drop culture's engineering challenges:
 - Extreme distributed contention
 - Idempotency guarantees
 - Limited inventory
 - Fairness in presence of scalper bots 

## Architecture

Built on top of `chi`, a lightweight HTTP router. Data layer comprised of Postgres and Redis.
