进入到 src/mattermost-plugin-svn/main
go build -o ../bin/svnplus.exe main.go
会在同目录生成 main.exe

需要创建 目录
mattermostsvnplus
-bin
  svnplus.exe
-conf
  config.toml

创建完成
cmd 到该目录bin下 执行 svnplus.exe -text=hello 
好使了.


2.将 post-commit.bat 放到 仓库目录的 hook目录下 D:\svnLocalRepo\localsvn\hooks
3.在提交一个commit 试试


