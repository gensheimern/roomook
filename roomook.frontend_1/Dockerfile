FROM node:11.13.0 AS build
WORKDIR /app
COPY . .
RUN   npm install \
   && npm run build

FROM nginx:1.13.3-alpine

COPY nginx/default.conf /etc/nginx/conf.d/

COPY --from=build /app/dist /usr/share/nginx/html
RUN rm -rf /usr/shar/nginx.html/*
ENTRYPOINT ["nginx", "-g", "daemon off;"]
