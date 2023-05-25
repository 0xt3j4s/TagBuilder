## Building the Docker Container

```
$ sudo docker build -t $DOCKER_USER_ID/tag-builder-logic .
```

## Running the Docker Container

```
$ sudo docker run -d -p 5010:5010 $DOCKER_USER_ID/tag-builder-logic
```

The app is listening by default on port 5010. The 5010 port of the host machine is mapped to the port 5010 of the container.

-p 5010:5010 i.e.

``` -p <hostPort>:<containerPort>```

### Verifying that it works

Execute a POST on endpoint 

-> `localhost:5010/generate-tags/` or 

-> `<docker-machine ip>:5010/generate-tags` Docker-machine ip has to be used if your OS doesn't provide native docker support. 

Request body:

```
{
    "sentence": "I enjoy playing flute in my free time."
}
```

## Pushing to Docker Hub

```
$ sudo docker push $DOCKER_USER_ID/tag-builder-logic
```