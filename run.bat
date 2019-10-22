cd .\src\github.com\mykelangelo\plainRockScissorsPaper
cd .\src\main
set name=plainRockScissorsPaper.exe
set app=".\..\..\..\..\..\..\bin\%name%"
taskkill /F /IM %name%
go build -o %app%
%app%
exit
