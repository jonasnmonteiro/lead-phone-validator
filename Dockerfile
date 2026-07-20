FROM node:22-alpine

WORKDIR /app

COPY --chown=node:node package*.json ./

RUN npm install --production

COPY --chown=node:node . .

USER node

EXPOSE 3007

CMD ["node", "server.js"]

