# k8s-helm-grafana
This repo contains an example setup of Grafana and Prometheus on Kubernetes using Skaffold and Helm. It also contains a very simple service (written in Go) which Prometheus is scraping metrics from.

The purpose is to show how these components connect and how you can make that work in a local setup or dev environment. 

Enjoy, and may the force be with you! 🖖

## Usage
Make sure you have a Docker and a Kubernetes cluster at your disposal. You will also need Skaffold and Helm installed. Then simply: 

```
$ skaffold run
```

To clean things up when you're done:

```
$ skaffold delete
```

## Details
`skaffold.yaml` contains the project components: a Go service called `prom` and Prometheus and Grafana based on remote Helm charts.

`The Go service` consists of main.go and a Dockerfile. Simple as that. The docker image is built as part of `skaffold run`.

The service exposes two endpoints. `/metrics` allows Prometheus to scrape metrics from the service, while `/hello` is a simple endpoint which counts each invocation on a Prometheus Counter from the Go Prometheus SDK. 

As this is our service, and our Helm chart, ports and other values are simply declared in `./helm-chart/values.yaml`. 

An important piece of the puzzle is somewhat hidden. Check `./helm-chart/templates/service.yaml` and you'll notice a bunch of Prometheus related properties. These are necessary to instruct Prometheus on how and where to scrape metrics from the service.

From your browser, try http://localhost:30003/hello and http://localhost:30003/metrics! 

`Prometheus` is added by using a remote Helm chart. Some values are then overridden to change the service type to NodePort and set a specifc port to make it easy to access the service from your browser. 

You can access the Prometheus dashboard on http://localhost:30001

`Grafana` is added by using a remote Helm chart. Some values are then overridden to change the service type to NodePort and set a specifc port to make it easy to access the service from your browser. 

Also, to make life easier, the admin account credentials are declared explicitly in skaffold.yaml. 

Try it out at http://localhost:30002. Once logged in, you can add a Prometheus data source. The internal cluster URL to use is http://prometheus-server.default.svc.cluster.local:80

