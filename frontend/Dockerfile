# Gunakan base image Node.js
FROM node:18 AS build

# Set working directory
WORKDIR /app

# Copy file package.json dan package-lock.json
COPY package*.json ./

# Update npm dan install dependencies
RUN npm install -g npm@latest
RUN npm install

# Copy seluruh kode sumber
COPY . .

# Build aplikasi Vue.js
RUN npm run build

# Gunakan Nginx untuk melayani aplikasi
FROM nginx:alpine

# Copy hasil build dari stage sebelumnya
COPY --from=build /app/dist /usr/share/nginx/html

# Expose port yang digunakan oleh Nginx
EXPOSE 80

# Jalankan Nginx
CMD ["nginx", "-g", "daemon off;"]
