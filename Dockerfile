##########
## This is an example if you would use node.js to solve the challenge.
#FROM node:10-alpine
#WORKDIR /usr/src/app
#COPY package*.json ./
#RUN npm install
#COPY . .

###########
# Keep this port exposed to simplify the acceptance test.
# Make sure your implementation works on this port.
EXPOSE 3000

#CMD ["node", "src/app.js"]