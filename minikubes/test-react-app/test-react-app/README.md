https://www.codementor.io/@skynettechltd/running-a-react-app-on-local-kubernetes-cluster-on-windows-10-1gjrkcwi66

```sh
  minikube docker-env
  eval $(minikube -p minikube docker-env)
```

```sh
docker build -t test-react-app .

kubectl apply -f ./k8s0/deployment.yaml
kubectl get pods
```

```sh
  minikube service test-react-app --url
```
and when you visit into browser http://192.168.64.3:32000 - should see okiok default react page
