FROM node:16

ARG VITE_API_URL

RUN npm install -g pnpm

WORKDIR /usr/src/dew-web
COPY . .

RUN pnpm install && pnpm run build

EXPOSE 5000
ENV HOST=0.0.0.0
ENTRYPOINT ["pnpm", "run", "serve", "--", "--host"]
