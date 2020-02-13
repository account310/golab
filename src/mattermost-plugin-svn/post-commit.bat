@ECHO OFF
REM *************************************************************
REM * Edit the following lines to suit your environment         *
REM *************************************************************
REM Set full path to svnlook.exe
SET SVNLOOK="D:\Program Files\visualsvn\bin\svnlook.exe"
REM Set full path to svnplus.exe
SET POSTEXE="D:\projectLab\golab\src\mattermost-plugin-svn\bin\svnplus.exe"
SET POSTCONFIG="D:\projectLab\golab\src\mattermost-plugin-svn\conf\config.toml"
REM *************************************************************
REM * This sets the arguments supplied by Subversion            *
REM *************************************************************
SET REPOS=%1
SET REV=%2


REM *************************************************************
REM * Get Author and comment                                    *
REM *************************************************************
setlocal EnableDelayedExpansion

rem get comments
SET COMMENT=
for /f "tokens=*" %%i in ('%SVNLOOK% log -r %REV% %REPOS%') do set COMMENT=!COMMENT! %%i
rem get author
SET AUTHOR=
for /f "delims=" %%t in ('%SVNLOOK% author %REPOS% -r %REV%') do set AUTHOR=%%t

rem get file_list
SET FILELIST=
for /f "tokens=*" %%i in ('%SVNLOOK% changed %REPOS% -r %REV% ') do set FILELIST=!FILELIST! %%i

REM *************************************************************
REM * Hand it to commit                                       *
REM *************************************************************
"%POSTEXE%" -conf="%POSTCONFIG%" -author="%AUTHOR%" -comments="%COMMENT%" -filelist="%FILELIST%" -projectpath="%REPOS%" -sendtype="" -rev="%REV%"
exit 0