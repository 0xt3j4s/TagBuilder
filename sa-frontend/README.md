## Setting up React for Local deployment
` $ npm install `

## Starting the Web App Locally
` $ npm start `

## Building the application [Making the app production ready]
` $ npm run build `

## Building the container 
` $ sudo docker build -t $DOCKER_USER_ID/tag-builder-frontend . `

## Running the container
` $ sudo docker run -d -p 80:80 $DOCKER_USER_ID/tag-builder-frontend `

## Pushing the container
` $ sudo docker push $DOCKER_USER_ID/tag-builder-frontend`

Note: All the commands are supposed to be executed in the sa-frontend directory