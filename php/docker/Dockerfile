FROM php:8.0-cli

#项目地址
ARG FILE=${FILE}
#composer命令  有：运行  空：不运行
ARG COMPOSER_UPDATE=${COMPOSER_UPDATE}
#端口
ARG PORT=${PORT}
# Use ustc source insteads of debian official source
RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list \
&& sed -i 's|security.debian.org/debian-security|mirrors.ustc.edu.cn/debian-security|g' /etc/apt/sources.list
#安装linux依赖包
#安装php扩展
#安装composer
RUN apt-get update && apt-get install -y libzip-dev && apt-get install -y zip \
    && apt-get install -y bash \
    && docker-php-ext-install bcmath mysqli pdo_mysql pcntl zip \
    && pecl install redis-5.3.2 \
    && pecl install swoole-4.5.9 \
    && docker-php-ext-enable redis swoole \
    && php -r "copy('https://getcomposer.org/download/2.0.7/composer.phar', 'composer.phar');" \
    && mv composer.phar composer \
    && chmod +x composer \
    && mv composer /usr/local/bin \
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

#拷贝项目 正式环境使用
COPY . ${FILE}
#工作目录
WORKDIR ${FILE}
RUN mv php.ini $PHP_INI_DIR/php.ini
#    && ${COMPOSER_UPDATE} || composer update --no-plugins
#暴露端口
EXPOSE ${PORT}
