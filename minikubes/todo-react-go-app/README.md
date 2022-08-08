
### CASE1:
  BE
    Deployment
      containers:
        ports:
          containerPort: 32001

    Service
      name: todo-golang-react-be-service
      type: NodePort
      ports:
        - port: 32001
          targetPort: 32001
          protocol: TCP
          nodePort: 32001

    in go code - project will be started at 32001 port !

  UI
    Deployment
      containers
        ports:
          containerPort: 80

    Service
      name: todo-golang-react-ui-service
      type: NodePort
      ports:
        - port: 80
          targetPort: 80
          nodePort: 32000

  ingress
    host: hello-world.info
    paths:
      path: /
      backend:
        name: todo-golang-react-ui-service
        port:
          number: 80
      path: /api
      backend:
        name: todo-golang-react-be-service
        port: 
          number: 

  if click at http://192.168.64.3/ - doesn't working!
  if click at hello-world.info - oki ok!

### how to launch and review ?
  kubectl apply -f ./k8s/be.deployment.yaml
  kubectl apply -f ./k8s/be.service.yaml
  kubectl apply -f ./k8s/ui.deployment.yaml
  kubectl apply -f ./k8s/ui.service.yaml
  kubectl apply -f ./k8s/ingress.yaml

  minikube dashboard
