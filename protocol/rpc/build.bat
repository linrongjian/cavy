@echo off&setlocal enabledelayedexpansion
for /f %%i in ('dir /b^|findstr "\<*.proto\>"') do (
	set files=!files!%%i 
)
protoc --gofast_out=plugins=grpc:./ %files% 

pause 