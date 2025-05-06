FROM node:22-alpine AS build
WORKDIR /app

# Copy package.json and package-lock.json and install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy the entire project
COPY . .

# Build the React app
RUN npm run build

# Use Nginx for serving the static files
FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html

# Expose the port the app runs on
EXPOSE 3000

CMD ["nginx", "-g", "daemon off;"]