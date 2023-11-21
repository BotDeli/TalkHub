FROM coturn/coturn:latest

COPY turnserver.conf /etc/coturn/turnserver.conf

EXPOSE 3478 5349

CMD ["turnserver", "-c", "/etc/coturn/turnserver.conf"]