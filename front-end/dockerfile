FROM oven/bun

WORKDIR /app

COPY . .

EXPOSE 2207

RUN bun i --verbose

CMD ["bun", "dev"]