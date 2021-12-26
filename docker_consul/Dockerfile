FROM hashicorp/consul:1.10.0

LABEL author="haibojiang@bitorobotics.ltd"

RUN apk --update add redis

CMD ["agent", "-dev", "-client", "0.0.0.0"]