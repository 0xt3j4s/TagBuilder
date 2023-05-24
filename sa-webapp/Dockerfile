# Start from a base Go image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . /app/

# Build the Go application
RUN go build -o tagbuilder .

# Set the environment variable for the Python API URL
ENV PYTHON_API_URL=http://localhost:5000/generate-tags/

# Expose the port that the web application will listen on
EXPOSE 8080

# Define the command to run the web application
CMD ["./tagbuilder"]
