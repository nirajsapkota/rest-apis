# Build
FROM node:14.17.0-alpine3.10

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . ./

RUN npm run build


# Execute
FROM node:14.17.0-alpine3.10

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install --only=production

COPY --from=0 /usr/src/app/dist ./dist

EXPOSE 8080

CMD [ "npm", "run", "serve" ]