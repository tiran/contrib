# Based on https://github.com/docker-library/mysql
FROM mysql
VOLUME /var/lib/mysql
ENTRYPOINT ["/custodia/go-secret"]
CMD ["/entrypoint.sh", "mysqld"]
