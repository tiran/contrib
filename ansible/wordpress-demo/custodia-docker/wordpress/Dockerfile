# Based on https://github.com/docker-library/wordpress
FROM wordpress
VOLUME /var/www/html
ENTRYPOINT ["/custodia/go-secret"]
CMD ["/entrypoint.sh", "apache2-foreground"]

