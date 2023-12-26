# Hello HTTP Server

Simple golang app that exposes test endpoint, so we can simulate different type of traffic

## Endpoints

- /hello returns 200
- /error returns 500
- /bad returns 400
- /random/latency returns 200 within 1s to 2s second duration
- /metrics exposes prometheus metrics
