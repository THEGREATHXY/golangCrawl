## git分支开发及合并流程

### 分支开发
1. 拉取项目代码：git clone git@github.com:THEGREATHXY/golangCrawl.git

2. 本地建立分支：git checkout -b branch_name

    1. 功能分支命名规则：feature_username_featurename

    2. Bug分支命名规则：bug_username_bugname

    3. 例如：feature_heyifei_userauthority

3. 分支开发+本地提交

4. 将本地新建分支同步到远程：git push origin branch_name

5.  分支开发+本地提交

6.  常规同步：git push 即可

7. 循环5和6

### 分支合并
1. 切换到主干：git checkout master

2. 拉取最新主干代码：git pull

3. 切换到开发分支：git checkout branch_name

4. 合入最新主干代码：git merge master

5. 如有冲突，解决冲突；如无冲突，下一步

6. commit整合后的分支代码并push到远程

7. 登录gitee.com，进入项目：

    1. 例如：https://gitee.com/he_123_admin/metas

8. 在+New的下拉菜单中选择New Merge request，填入需要CR的人，即可发起MR

1. 其中有一人approve之后，MR进入可以Merge的状态，点击Merge，系统自动完成Merge操作，流程结束。

1. 本地切换到master分支，pull代码。

1. 新建分支，开始下一次搬砖

1. 对于无用的分支可以定期删除或者勾选合并后删除提交的分支，注意在本地也要删除
    1. 可以参考 https://blog.csdn.net/xjzdr/article/details/104694367
