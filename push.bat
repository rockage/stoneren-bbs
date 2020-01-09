@echo off
echo Input you commit ......
set /p commit=

git add .
git commit -m %commit%
git push -u origin master
