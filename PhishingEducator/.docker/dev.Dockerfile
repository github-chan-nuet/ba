FROM node:22-alpine AS build
WORKDIR /app

# Copy package.json and package-lock.json and install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy the entire project
COPY . .
COPY .docker/phishing_educator.dev.env .env

# Expose the port the app runs on
EXPOSE 5173

CMD ["npm", "run", "dev"]