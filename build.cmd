cls
set ver=0.1
set geek=D:\dev\gopath\src\github.com\ecdiy\itgeek
del %geek%\ui\geek-%ver%-linux.zip
del %geek%\ui\geek-%ver%-windows.zip
del %geek%\ui\dist\geek.sql
del %geek%\ui\dist\geek.exe
del %geek%\ui\dist\geek
copy %geek%\db\geek.sql %geek%\ui\dist\geek.sql
cd %geek%\gk\cmd
set GOOS=windows
set GOARCH=amd64
go build  -ldflags "-s -w" -o %geek%\ui\dist\geek.exe
echo go build ok,build admin ...
cd %geek%\ui
call npm run build-admin
call npm run build-web
call npm run build-m
%geek%\ui
%geek%\tools\7z.exe  a -t7z %geek%\ui\geek-%ver%-windows.zip "%geek%\ui\dist\*" -r -mx=9 -m0=LZMA2 -ms=10m -mf=on -mhc=on -mmt=on
echo "build windows zip ok. build linux ..."
cd %geek%\gk\cmd
set GOOS=linux
set GOARCH=amd64
go build -ldflags "-s -w" -o %geek%\ui\dist\geek
del %geek%\ui\dist\geek.exe
cd %geek%\ui
%geek%\tools\7z.exe  a -t7z %geek%\ui\geek-%ver%-linux.zip "%geek%\ui\dist\*" -r -mx=9 -m0=LZMA2 -ms=10m -mf=on -mhc=on -mmt=on
cd %geek%