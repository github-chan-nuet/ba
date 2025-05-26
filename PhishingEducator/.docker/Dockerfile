FROM node:22-alpine AS build
WORKDIR /app

# Copy package.json and package-lock.json and install dependencies
COPY package.json package-lock.json ./
RUN npm ci

# Copy the entire project
COPY . .
COPY .docker/phishing_educator.env .env

# Build the React app
RUN npm run build

# Use Nginx for serving the static files
FROM nginx:alpine
COPY --from=build /app/build/client /usr/share/nginx/html

# Copy custom Nginx config
COPY .docker/nginx.conf /etc/nginx/conf.d/default.conf

# Expose the port the app runs on
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]