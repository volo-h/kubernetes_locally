#### oki ok
#### how to

```sh
kubectl apply -f k8s1/deployment.yaml
  kubectl get deployment

kubectl expose deployment klaus-hello-world --type=NodePort --port=8080
  kubectl get service klaus-hello-world

minikube service klaus-hello-world --url
```

### delete if need
```sh
  kubectl delete deployment klaus-hello-world 
  kubectl delete svc klaus-hello-world
```

### statistic, working okiok
```sh
  kubectl get deployments | grep klaus-hello-world
```
  klaus-hello-world   2/2     2            2           14m

```sh
  kubectl get svc | grep klaus-hello-world
```
  klaus-hello-world      NodePort    10.96.159.164    <none>        8080:30108/TCP   99s

```sh
  minikube service klaus-hello-world --url
```
  http://192.168.64.3:30108

#### show existing endpoints of service klaus-hello-world
```sh
  kubectl get endpoints | grep klaus-hello-world
```
  klaus-hello-world      172.17.0.2:8080,172.17.0.3:8080   27s

request from IP in browser -> (port of "NodePort" service)
or by curl

ex.:

```sh
  curl http://192.168.64.3:32036
```
  Hello Klaus!!%

```sh
  kubectl get nodes -o wide
```

deployment.yaml + expose deployment with NodePort and at all !
