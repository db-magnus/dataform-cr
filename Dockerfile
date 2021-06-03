FROM node:14-slim
WORKDIR /dataform
COPY index.js ./
COPY dataform.json ./
COPY definitions ./
COPY package.json ./
COPY script.sh ./

RUN yarn global add @dataform/cli
RUN dataform install

CMD [ "node", "index.js" ]
