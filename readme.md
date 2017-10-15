# Wired Sensors

### Getting Started 
    
- Download Docker for your OS: https://www.docker.com/community-edition#/download

- Compile the application

    # Build the application w/ Docker
    docker build -t wired-sensors .

- Run the application

    # Run the application w/ Docker
    docker run -it --rm -p 6060:8080 --name wired-sensors wired-sensors

- Browse to localhost:6060
