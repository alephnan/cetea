FROM node:lts-alpine

WORKDIR /app
RUN npm i -g @vue/cli 
# TODO: exclude node_modules
COPY . .

# Outputs to dist directory
RUN npm run build

RUN mkdir -p /var/www/app && mv dist/* /var/www/app