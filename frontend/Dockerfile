FROM node:lts-alpine

WORKDIR /app
COPY package.json .
COPY src/ src/

# Outputs to dist directory
RUN npm run build

RUN mkdir -p /var/www/app && mv dist/* /var/www/app