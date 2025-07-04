# Use official PHP image with FPM
FROM php:8.2-fpm

# Install system dependencies and PHP extensions needed for Laravel
RUN docker-php-ext-install pdo pdo_mysql sockets 
RUN curl -sS https://getcomposer.org/installer​ | php -- \      
    --install-dir=/usr/local/bin --filename=composer  
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer 

# Copy application code
WORKDIR /var/www/html
COPY . .

# Copy Nginx config
COPY default.conf /etc/nginx/conf.d/default.conf

# Set permissions for Laravel storage and cache
RUN chown -R www-data:www-data /var/www/html/storage /var/www/html/bootstrap/cache

# Expose port 80 for nginx
EXPOSE 80

CMD ["php-fpm"]