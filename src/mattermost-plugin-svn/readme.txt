该插件根据 https://github.com/bitbackofen/Subversion-Integration-for-Mattermost 的golang版本
具体依赖库自己下载即可。

使用方法:
1.修改对应的post-commit文件
配置指定的路径
2.修改配置文件



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

$ git rm -r --cached aaa           # 本地git删除aaa文件夹
$ git commit -m '删除aaa'        # 提交,添加操作说明
git add .
git commit -m ""
新建一个 github 的仓库，在网页新建，记下https地址，也就是仓库的地址，输入
git remote add github https://自己的github仓库url地址，回车。
输入 git push -u github master, 将代码上传到github 仓库中，
会提示你输入github的账号密码，接着完成就可以了。

git pull github master  拉取远端到本地
git push -f origin master 强制推送到远端




