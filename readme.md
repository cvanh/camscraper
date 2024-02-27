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
