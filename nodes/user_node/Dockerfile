FROM python:3.10-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY ./src ./src
COPY ./nodes ./nodes
ENV PYTHONPATH=/app/src/python
CMD ["bash", "nodes/user_node/start.sh"]