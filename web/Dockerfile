FROM php:8.2-fpm

RUN docker-php-ext-install pdo pdo_mysql sockets 
RUN curl -sS https://getcomposer.org/installer​ | php -- \      
    --install-dir=/usr/local/bin --filename=composer  
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer 

WORKDIR /var/www/html
COPY . .

COPY default.conf /etc/nginx/conf.d/default.conf

RUN chown -R www-data:www-data /var/www/html/storage /var/www/html/bootstrap/cache

EXPOSE 8080

CMD ["php-fpm"]