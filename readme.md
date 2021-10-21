# go web Demo
打桩代码 TestSystem
- 不影响现有代码，引入桩代码TestSystem
- TestSystem里监听各类中断信号

Demo
本地起来之后访问 http://localhost:9090/?number=0
因为是Demo所以代码不是很严谨，只给了3个分支 number 大于0，小于0，等于0的场景来测试覆盖率

集成步骤
- 1：执行集成测试，并将此函数编译成二进制文件
  ```shell
  go test -covermode=count -coverpkg="./..." -c -o cover.test
  ```
- 2：运行二进制文件，指定运行的测试方法是 TestSystem，启动go的main方法。如果是单测，走现有的执行方法，这里展示执行main方法
  ```
  cover.test -test.run "TestSystem" -test.coverprofile=cover.cov -systemTest=true
  ```
  （这里的systemTest 是在桩代码设置的，如果没有设置，或者是false，TestSystem 不启动main方法）
- 3：跑完接口测试后中断进程 （建议 kill -15 {pid}，目前没有捕捉-9的信号，也可以加），生产cover.cov
- 4：运行时的 cover.cov 和单测的 *cov文件合并
  ``` shell
  mkdir coverage
  echo 'mode: count' > ./coverage/system.cov
  tail -q -n +2 ./coverage/*.cov >> ./coverage/system.out
  ```
- 5：转成html文件
  ```
  go tool cover -html cover.out -o cover.html   
  ```
- 6：Sonar集成
 
