# KRR UI

## Description

...

## Prerequisites
- Running prometheus

## Notes
- KRR-UI and KRR-UI api must be publicly accessible for correct work
- KRR-UI api service account must have rights for pod creation
- Prometheus must be accessible from pod with KRR ran by KRR-UI api

## Steps to install

### Build required images:
```
cd ui && docker build -t krr-ui . && cd -
cd api && docker build -t krr-ui-api . && cd -
```

### Load images to k8s:
For minikube:
```
minikube image load krr-ui
minikube image load krr-ui-api
```

### Apply deployment manifests:
```
kubectl apply -f ./deploy/ui.yaml
kubectl apply -f ./deploy/api.yaml
```

Done! Now you can access KRR-UI.