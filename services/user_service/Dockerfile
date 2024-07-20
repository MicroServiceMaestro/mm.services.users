FROM python:3.10-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY ./src ./src
COPY ./services ./services
ENV PYTHONPATH=/app/src/python
EXPOSE 8888
CMD ["bash", "services/user_service/start.sh"]