FROM nginx:alpine

COPY ./views/ /usr/share/nginx/html/

COPY ./css/ /usr/share/nginx/html/css/
COPY ./js/ /usr/share/nginx/html/js/
COPY ./images/ /usr/share/nginx/html/images/
COPY ./assets/ /usr/share/nginx/html/assets/

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
