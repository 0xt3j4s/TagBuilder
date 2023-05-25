# Start from a base Go image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . /app/

# Set the build arguments
ARG PYTHON_API_URL=http://172.17.0.2/generate-tags/

# Set the environment variable for the Python API URL
ENV PYTHON_API_URL=$PYTHON_API_URL

# Build the Go application with the build argument
# RUN go build -o tagbuilder -ldflags "-X main.pythonURL=$PYTHON_API_URL" .
RUN go build -o tagbuilder .

# Expose the port that the web application will listen on
EXPOSE 8081

# Define the command to run the web application
CMD ["./tagbuilder"]