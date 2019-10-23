#run local instance (Windows 10 script only yet)
cd .\src\github.com\mykelangelo\plainRockScissorsPaper
cd .\src\main
set port=80
set name=plainRockScissorsPaper.exe
set app=".\..\..\..\..\..\..\bin\%name%"
taskkill /F /IM %name%
go build -o %app%
%app%
exit
