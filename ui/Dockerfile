FROM node:22 AS builder

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx AS production-stage

RUN mkdir /app

COPY nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=builder /app/build /app

EXPOSE 80