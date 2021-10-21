#/bin/bash

#cd src
SHELL_FOLDER=$(cd "$(dirname "$0")";pwd)
cd ${SHELL_FOLDER}/src

#第一步：执行集成测试，并将此函数编译成二进制文件
go test -covermode=atomic -coverpkg="./..." -c -o cover.test
#第二步：运行二进制文件，指定运行的测试方法是 TestMainStart，并将覆盖率报告输出
./cover.test -test.run "TestSystem" -test.coverprofile=cover.cov -systemTest=true
#第三步：将输出的覆盖率报告转换成 html 文件（html 文件查看效果比较好）
go tool cover -html cover.cov -o cover.html
#第四步：生成 Cobertura 格式的 xml 文件 ，，按需
#gocov convert cover.cov | gocov-xml > cover.xml
