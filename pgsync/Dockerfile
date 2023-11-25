FROM bitnami/python:3.9

RUN mkdir -p /home/nobody

WORKDIR /home/nobody

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh wait-for-it.sh

RUN chown -R nobody:nogroup /home/nobody && chmod +x wait-for-it.sh

USER nobody

ENV HOME=/home/nobody

RUN pip install git+https://github.com/toluaina/pgsync.git

COPY --chown=nobody:nogroup run.sh run.sh

ENTRYPOINT ~/run.sh
