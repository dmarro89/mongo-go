# mongo-go
Here's a simple application written in go (1.19) communicating with a mongo-db instance.
It exposes some Rest API (made using GIN) in order to communicate with it.

## Configuration
The Dockerfile must be updated to guarantee that the application communicates with a valid mongodb instance.
Change the **ENV CONNECTION_URI** entry with a valid value.

```docker
ENV CONNECTION_URI mongodb+srv://test@localhost
```
## Build and Run
To Build the docker image, from the project root path:

```
docker build -t mongo-go .
```

Run your container based on the image just built:
```
docker run -p 127.0.0.1:8080:8080 mongo-go
```
## Deploy on k8s
To Deploy your docker image on kubernetes cluster:
- First of all, you should push the docker image on the repository that you prefer.
```
docker login
docker push <your_repo>/mongo-go:latest
```
- You can modify the k8s/deployment.yaml file changing the **image** entry with your entry name. At the moment, the yaml file is using an already pushed image of the last project version on a public repository.
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mongo-go
  labels:
    app: mongo-go
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mongo-go
  template:
    metadata:
      labels:
        app: mongo-go
    spec:
      containers:
        - name: mongo-go
          image: davidemarro/mongo-go:latest
          ports:
            - containerPort: 8080
          imagePullPolicy: Always
```
- Than you can deploy the docker image to your cluster
```
kubectl apply -f deployment.yaml -n <your_namespace>
```
- Apply the service.yaml file too to expose the service on port 8080
```
kubectl apply -f service.yaml -n <your_namespace>
```
## Rest API
- POST http://localhost:8080/user -> {"name":"Davide", "surname":"Marro", "email":"davide@marro.it"}
- PUT http://localhost:8080/user -> {"id": "000000000000000","name":"Davide", "surname":"Marro", "email":"davidemarro@marro.it"}
- GET http://localhost:8080/user/:userId
- DELETE http://localhost:8080/user/:userId
- GET http://localhost:8080/users



