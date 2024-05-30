# Ultimate Debugging with Go - Metrics module
This repo is created for [The Ultimate Debugging with Go course](https://www.bytesizego.com/the-ultimate-guide-to-debugging-with-go)

## Getting started
You should simply be able to run `docker compose up` to get everything running (assuming docker is installed).
Once that is done, you should have:
- Prometheus available on http://localhost:9090
- Grafana available on http://localhost:3000
- The app on http://localhost:8080

The following endpoints should be available to you:
- http://localhost:8080/metrics
- http://localhost:8080/tx
- http://localhost:8080/duration
- http://localhost:8080/size
- http://localhost:8080/blood

Have a play around and build your own dashboard!

The username for Grafana is defined in the docker compose file. It defaults to admin/yourpassword
