FROM node:16-alpine


WORKDIR /app
COPY pong.wasm index.html wasm_exec.js /app/
RUN npm install -g npm && npm i -g http-server

EXPOSE 8080
CMD ["http-server"]