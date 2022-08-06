
```sh
  kubectl get deployments
  kubectl delete deployment klaus-hello-world

  kubectl get svc | grep klaus-hello-world-service
  kubectl delete svc klaus-hello-world-service

  kubectl get ing | grep main-ingress
  kubectl delete ing main-ingress
```

####
```sh
  kubectl apply -f ./k8s2/deployment.yaml
  kubectl apply -f ./k8s2/service.yaml
  kubectl apply -f ./k8s2/ingress.yaml
```

#### apply all in one command
```sh
  kubectl apply -f k8s2
```

#### get real ip of minikube; that ip we sill add to /etc/hosts
```sh
  minikube ip
```

add to /etc/hosts

```sh
echo "192.168.64.3  hello-world.info" | sudo tee -a /etc/hosts
```

curl hello-world.info
curl hello-world.info/2

```sh
curl http://<EXTERNAL-IP>:<PORT>
  ex. http://10.101.208.96:8080/ - okiok
```

#### start web ui
```
  minikube dashboard
```

Access over http://192.168.64.3/ crashed by 404 Not Found
but
access over browser http://hello-world.info/
or
```sh
 curl http://hello-world.info/
```
Hello Klaus!!%

woks great

#### if you need investigate occured errors
```sh
  kubectl describe pods klaus-hello-world
  kubectl describe ing
  kubectl describe svc klaus-hello-world
```

#### https://minikube.sigs.k8s.io/docs/commands/service/
Lists the URLs for the services in your local cluster
```sh
minikube service list
```

#### how to with ingress
#### from https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/
https://kubernetes.io/ru/docs/tutorials/hello-minikube/

```sh
  cd mysql
  ✗ kubectl apply -f mysql-pv.yaml
  ✗ kubectl apply -f mysql-deployment.yaml
  ✗ kubectl describe deployment mysql
  ✗ kubectl get pods -l app=mysql
  ✗ kubectl describe pvc mysql-pv-claim
  ✗ kubectl run -it --rm --image=mysql:8.0 --restart=Never mysql-client -- mysql -h mysql -ppassword
```
https://kubernetes.io/docs/tasks/run-application/run-single-instance-stateful-application/

✗ kubectl get pods
✗ kubectl get deployments
✗ kubectl get ing


