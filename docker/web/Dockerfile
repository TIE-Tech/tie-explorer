FROM node:latest as build-stage

WORKDIR /web

COPY ./web ./

RUN npm install

RUN npm run build

FROM nginx:1.15.3-alpine as production-stage

RUN rm /etc/nginx/conf.d/default.conf
RUN rm /etc/nginx/nginx.conf

COPY ./docker/web/nginx.conf /etc/nginx/nginx.conf

COPY --from=build-stage /web/dist /usr/share/nginx/html

EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]