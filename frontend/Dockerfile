FROM node:lts-alpine

WORKDIR /app
# Worksaround: https://github.com/npm/npm/issues/20861
RUN npm config set unsafe-perm true
COPY . .
RUN npm i

# Outputs to dist directory
RUN npm run build

RUN mkdir -p /var/www/app && mv dist/* /var/www/app