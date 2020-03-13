## openmetrics-mocker
Simple go server that serves a textfile from disk for its metrics endpoint.
Allows for manual testing of OpenMetrics scrapers and visualizers.

### Development and Testing

Run locally:
```bash
go run cmd/server.go
```
Build locally:
```bash
docker build -t openmetrics-mocker -f build/Dockerfile .
```

Deploy to Kubernetes:
```bash
# NOTE: modify the configmap to update the metrics whenever you need.
#       you may also want to modify the image if testing your own
#       version of the image.
kubectl apply -f deployments/deploy.yaml
```
