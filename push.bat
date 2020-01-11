@echo off
git status
pause
echo Input you commit ......
set /p commit=

git add .
git commit -m %commit%
git push -u origin master
