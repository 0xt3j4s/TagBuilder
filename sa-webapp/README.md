## Starting the Web App Locally
` $ go run main.go `

### Build
` $ go build -o tagbuilder .`

### Run
` $ ./tagbuilder `

## Building the container 
` $ sudo docker build -t $DOCKER_USER_ID/tag-builder-webapp . `

## Running the container
` $ sudo docker run -d -p 80:80 $DOCKER_USER_ID/tag-builder-webapp `

## Pushing the container
` $ sudo docker push $DOCKER_USER_ID/tag-builder-webapp`

Note: All the commands are supposed to be executed in the sa-webapp directory