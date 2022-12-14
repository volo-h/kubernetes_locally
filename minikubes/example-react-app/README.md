# Getting Started with Create React App

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.\
You will also see any lint errors in the console.

### `npm test`

Launches the test runner in the interactive watch mode.\
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

### `npm run eject`

**Note: this is a one-way operation. Once you `eject`, you can’t go back!**

If you aren’t satisfied with the build tool and configuration choices, you can `eject` at any time. This command will remove the single build dependency from your project.

Instead, it will copy all the configuration files and the transitive dependencies (webpack, Babel, ESLint, etc) right into your project so you have full control over them. All of the commands except `eject` will still work, but they will point to the copied scripts so you can tweak them. At this point you’re on your own.

You don’t have to ever use `eject`. The curated feature set is suitable for small and middle deployments, and you shouldn’t feel obligated to use this feature. However we understand that this tool wouldn’t be useful if you couldn’t customize it when you are ready for it.

## Learn More

You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).

To learn React, check out the [React documentation](https://reactjs.org/).

### Code Splitting

This section has moved here: [https://facebook.github.io/create-react-app/docs/code-splitting](https://facebook.github.io/create-react-app/docs/code-splitting)

### Analyzing the Bundle Size

This section has moved here: [https://facebook.github.io/create-react-app/docs/analyzing-the-bundle-size](https://facebook.github.io/create-react-app/docs/analyzing-the-bundle-size)

### Making a Progressive Web App

This section has moved here: [https://facebook.github.io/create-react-app/docs/making-a-progressive-web-app](https://facebook.github.io/create-react-app/docs/making-a-progressive-web-app)

### Advanced Configuration

This section has moved here: [https://facebook.github.io/create-react-app/docs/advanced-configuration](https://facebook.github.io/create-react-app/docs/advanced-configuration)

### Deployment

This section has moved here: [https://facebook.github.io/create-react-app/docs/deployment](https://facebook.github.io/create-react-app/docs/deployment)

### `npm run build` fails to minify

This section has moved here: [https://facebook.github.io/create-react-app/docs/troubleshooting#npm-run-build-fails-to-minify](https://facebook.github.io/create-react-app/docs/troubleshooting#npm-run-build-fails-to-minify)

#### some parts reused from 
https://medium.com/geekculture/dockerizing-a-react-application-with-multi-stage-docker-build-4a5c6ca68166
+
https://enlear.academy/deploy-a-react-application-in-kubernetes-69bd07e375ab

```sh
  ✗ docker build -t example-react-app .
  ✗ docker images example-react-app
  ✗ docker run -dit -p 3000:3000 example-react-app:latest
  ✗ docker ps | grep example-react-app
  ✗ docker stop <CONTAINER_ID>

  ✗ kubectl apply -f ./deployment.yaml
  ✗ kubectl get pods
  ✗ kubectl get deployments
  ✗ kubectl get services
  ✗ kubectl get ing

  ✗ minikube service example-react-app --url
```

### CASE1:
  Deployment
    containers:
      ports:
        containerPort: 80

  Service
    name: example-react-app
    type: NodePort
    ports:
      - port: 80
        targetPort: 80
        protocol: TCP
        nodePort: 32000

  ```sh
    minikube service example-react-app --url
  ```
  go to browser - okiok - http://192.168.64.3:32000

  TIPS: minikube run ports on the range of valid ports is 30000-32767
  but we can reassign that
  ```
    minikube start --extra-config=apiserver.service-node-port-range=3000-32767 --ports=127.0.0.1:3000-32767:3000-32767
  ```

### CASE2:
  Deployment
    containers:
      ports:
        containerPort: 80

  Service
    name: example-react-app
    type: NodePort
    ports:
      - port: 80
        targetPort: 80
        protocol: TCP
        nodePort: 32000

  Ingress
    host: hello-world.info
      name: example-react-app
        - pathType: Prefix
          path: /
          backend:
            service:
              name: example-react-app
              port: 
                number: 80

  in browser
    http://hello-world.info/ - okiok
    http://192.168.64.3/ - failed!

