# Wired Sensors

Wired Sensors is a web application that will help users make decisions based on real-time and historical environment sensor data.  The project will start with data from the NOAA Tides and Currents API (https://tidesandcurrents.noaa.gov/api/) and the World Air Quality Index API (http://aqicn.org/api/).

### Getting Started 
    
- Download Docker for your OS: https://www.docker.com/community-edition#/download

- Compile the application

```
    # Build the application w/ Docker
    docker build -t wired-sensors .
```

- Run the application

```
    # Run the application w/ Docker
    docker run -it --rm -p 6060:8080 --name wired-sensors wired-sensors
```

- Browse to localhost:6060


