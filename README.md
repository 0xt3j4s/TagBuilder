# TagBuilder

TagBuilder is a microservice-based application that enables the generation of tags for sentences. In this application, we utilize Docker containers to encapsulate and deploy each microservice independently. 

## Application Demo
This application has one functionality. It takes a sentence as input. Using Natural Language Toolkit (NLTK), generates tags for the sentence.

The application consists of three microservices, each one of specific functionality. The microservices are:
1. sa-frontend: a Nginx web server that serves our static ReactJS application.
2. sa-webapp: A Golang server that serves as the backend for handling the requests from the frontend.
3. sa-logic:A Flask Application that performs tag generation.

## Application Architecture
The application architecture is shown in the figure below:
![TagBuilder Architecture](/images/microservice_architecture.png)

## Containerization
To enable easy deployment and scalability, each microservice in this application has been containerized using Docker. 
![Containerized Microservices](/images/containerized_microservices.png)

## Pre-requisites

Before running the application, ensure that you have the following installed:

1. Nginx
2. Golang
3. Python

## Run

### Locally
To run the application locally, follow the steps below:
1. Clone the repository:
```bash
$ git clone https://github.com/0xt3j4s/TagBuilder.git
```
2. Change directory to the repository:
```bash
$ cd TagBuilder
```
3. Run the following commands to start the microservices: <br>


    
```bash
# a. Flask Application:
$ cd sa-logic
$ cd sa
$ pip install -r requirements.txt
$ python tag_builder.py
$ cd ..

# b. Golang Server:
$ cd sa-webapp
$ go build
$ cd ..

# c. ReactJS Application:
$ cd sa-frontend
$ npm install
$ npm run build
$ cd ..
```


### Docker
If you don't have Docker installed, you can follow the instructions [here](https://docs.docker.com/get-docker/).

Pull the container images from Docker Hub using the following commands:

```bash
$ sudo docker pull $DOCKER_USER_ID/tag-builder-logic

$ sudo docker pull $DOCKER_USER_ID/tag-builder-webapp

$ sudo docker pull $DOCKER_USER_ID/tag-builder-frontend
```

This will pull the latest images from Docker Hub and store them in your local machine.

To run the application, run the following commands:

```bash
$ sudo docker run -d -p 5010:5010 $DOCKER_USER_ID/tag-builder-logic

$ sudo docker run -d -p 8081:8081 -e PYTHON_API_URL='http://<container_ip or docker machine ip>:5010' $DOCKER_USER_ID/tag-builder-webapp

$ sudo docker run -d -p 80:80 $DOCKER_USER_ID/tag-builder-frontend
```



## Contributing
If you would like to contribute to TagBuilder, please follow the guidelines in CONTRIBUTING.md. We welcome contributions from the community.