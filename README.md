# fme - future me

This is a simple commandline utility that asks and saves to a text file a list of your goals.

# Docker image

To build Docker image, run in the project directory:
`docker build -t fme:v1.0.0 .`

To create a new container named my-fme-app:

`docker run --name my-fme-app -it fme:v1.0.0`

To run `my-fme-app` again, start the container in interactive mode:
`docker container start -i my-fme-app`
