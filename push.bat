@echo off
git init
git add .
git commit -m %1
git push origin master
php artisan serve