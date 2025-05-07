FROM node:24

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

EXPOSE 3000

VOLUME /usr/src/app/data

CMD [ "npm", "run", "run" ]
