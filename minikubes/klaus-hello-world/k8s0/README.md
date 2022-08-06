```sh
  kubectl apply -f k8s0/deployment.yaml
  kubectl expose deployment klaus-hello-world --type=LoadBalancer --name=klaus-hello-world-service
```

##### https://minikube.sigs.k8s.io/docs/handbook/accessing/#using-minikube-tunnel
```sh
  minikube tunnel
```
after type password in terminal that should created "Services" with accessible "External Endpoints" which you can see in runned dashboard

```sh
  minikube dashboard
```
