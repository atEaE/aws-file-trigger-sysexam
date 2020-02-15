FROM nginx:1.17.7-alpine
RUN rm -f /etc/nginx/conf.d/*
ADD nginx.conf /etc/nginx/conf.d/nginx.conf
CMD /usr/sbin/nginx -g "daemon off;"

EXPOSE 80