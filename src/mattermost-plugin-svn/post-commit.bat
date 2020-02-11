@ECHO OFF
REM *************************************************************
REM * Edit the following lines to suit your environment         *
REM *************************************************************
REM Set full path to svnlook.exe
SET SVNLOOK="D:\Program Files\visualsvn\bin\svnlook.exe"
REM Set full path to svnplus.exe
SET POSTEXE="D:\projectLab\matterpluginsvn\bin\svnplus.exe"
SET POSTCONFIG="D:\projectLab\matterpluginsvn\conf\config.toml"
REM *************************************************************
REM * This sets the arguments supplied by Subversion            *
REM *************************************************************
SET REPOS=%1
SET TXN=%2

REM *************************************************************
REM * Get Author and comment                                    *
REM *************************************************************
setlocal EnableDelayedExpansion
rem get comments
for /f "tokens=*" %%i in ('%SVNLOOK% log -r %TXN% %REPOS%') do set COMMENT=%%i
rem get author
for /f "tokens=*" %%i in ('%SVNLOOK% author -r %TXN% %REPOS%') do set AUTHOR=%%i

REM *************************************************************
REM * Hand it to commit                                       *
REM *************************************************************
"%POSTEXE%" -conf="%POSTCONFIG%" -text="%AUTHOR% committed revision %TXN% to %REPOS%: %COMMENT%"
pause
rem exit 0