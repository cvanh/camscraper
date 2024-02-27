## camscraper
a scraper to find the cheapest climbing cams


## structure
The scraper works with multible seprate scrapers which send data to the api via apache kafka. I a diagram this would look like this:
```mermaid
---
title: setup
---
flowchart 
    API[backend]
    
    scraperKL[klim winkel scraper]
    scraperBF[bergfreunde scraper]
    scraperBV[bever scraper]

    DB[(mysql)]

    user
    frontend

    scraperKL -->|kafka| API
    scraperBV -->|kafka| API
    scraperBF -->|kafka| API

    API --> DB

    frontend --> API
    user --> frontend
```

### folder structure
The folder structure for this project is the following:
```mermaid
---
Title: folder setup
---

flowchart 
    root[./]

    scrapers[scrapers/]
    klimwinkel_scraper[klimwinkel.nl scraper]
    bergfreunde_scraper[bergfreunde.eu scraper]
    bever_scraper[bever.nl scraper]

    infra[infra/]

    backend[backend/]

    root --> scrapers
    scrapers --> klimwinkel_scraper
    scrapers --> bergfreunde_scraper
    scrapers --> bever_scraper

    root --> infra

    root --> backend
```

