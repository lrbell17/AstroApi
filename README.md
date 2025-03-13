# AstroApi

AstroApi is a backend service that integrates real data from [NASA's exoplanet archive](https://exoplanetarchive.ipac.caltech.edu/cgi-bin/TblView/nph-tblView?app=ExoTbls&config=PSCompPars) and exposes a REST API to interact with exoplanet and star system data.

This service supports caching with Redis and uses PostgreSQL for persistent data storage.


### Components
- astroapi - The core service responsible for the API logic and interactions.
- redis - Caching layer to optimize response times and reduce load on the database.
- postgres -  Persistent database storing exoplanet and star system data.

### Key packages used:
- gorm: ORM library for persistent data storage and interaction with PostgreSQL.
- gin: Web framework for building the REST API.
- logrus: Structured logging for easy tracking and debugging. 

## How to:
- Start locally: `make start`
- Rebuild: `make clean`
- Stop: `make stop`
- Clean: `make clean`
- Tail logs: `make logs`
