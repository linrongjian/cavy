@echo off&setlocal enabledelayedexpansion
for /f %%i in ('dir /b^|findstr "\<*.proto\>"') do (
	set files=!files!%%i 
)
protoc.exe --go_out=../go %files%
pause