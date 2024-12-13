# Example

## Description
Instructions below show how to set up the required environment in minikube for KRR-UI from scratch (for testing purposes only).

## Prerequisites
- Running Minikube cluster
- Helm

## Steps

### Install metrics-server:
```
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```

### Skip tls check in metrics-server:
``kubectl edit deployment metrics-server -n kube-system``

add ``--kubelet-insecure-tls``:

```yaml
spec:
  containers:
  - name: metrics-server
    image: <metrics-server-image>
    args:
    - --kubelet-insecure-tls
    - --kubelet-preferred-address-types=InternalIP,Hostname,ExternalIP
```

### Install Prometheus via Helm:
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm search repo prometheus
helm install prometheus prometheus-community/prometheus
```

### Allow Krr to access Prometheus data:
```
kubectl apply -f ./example/deploy/allow-pod-view.yaml
```

### Allow KRR-UI api to run krr pods:
```
kubectl apply -f ./example/deploy/allow-pod-creation.yaml
```


Done! Now you can install and use KRR-UI