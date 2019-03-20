# Use the standard Nginx image from Docker Hub
FROM nginx
# The Dockerfile's author

# Copy the configuration file from the current directory and paste
# it inside the container to use it as Nginx's default config.
COPY nginx.template /etc/nginx/nginx.template

# Port 8080 of the container will be exposed and then mapped to port
# 8080 of our host machine via Compose. This way we'll be able to
# access the server via localhost:8080 on our host.
ARG port=8080
ARG port_app=8081
ENV PORT=$port
ENV PORT_APP=$port_app
EXPOSE $port

RUN envsubst < /etc/nginx/nginx.template > /etc/nginx/nginx.conf
ENTRYPOINT [ "nginx", "-g", "daemon off;" ]