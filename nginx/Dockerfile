FROM nginx:alpine
COPY ./nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf
COPY front/*  /var/www/web/front/
CMD ["nginx", "-g", "daemon off;"]
EXPOSE 80
