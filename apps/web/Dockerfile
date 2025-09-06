# syntax=docker/dockerfile:1
FROM oven/bun:1.2.21-alpine
WORKDIR /app
COPY package.json bun.lock* ./
RUN bun install
COPY . .
CMD ["bun", "dev", "--host", "0.0.0.0"]
