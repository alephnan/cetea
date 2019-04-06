FROM node:lts-alpine as builder
WORKDIR /app
# Worksaround: https://github.com/npm/npm/issues/20861
RUN npm config set unsafe-perm true
COPY package.json .
COPY package-lock.json* .
RUN npm i

FROM node:lts-alpine
WORKDIR /app
COPY --from=builder /app /app
COPY . .
# Outputs to dist directory
RUN npm run build

RUN mkdir -p /var/www/app && mv dist/* /var/www/app