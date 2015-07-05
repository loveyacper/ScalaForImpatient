Last login: Wed Jul  1 23:42:51 on ttys002
BertdeMacBook-Pro:Debug bertyoung$ LS
Qedis			qedis.conf		sendthread_log
cmd_stat_info		qedis.conf.bak		slowlog.qedis
dump.qdb.rdb		qedis_appendonly.aof	tmp_qdb_file_19411
dump.rdb		qedis_log		tmp_qdb_file_19419
libqbaselib.a		qedislog		tmp_qdb_file_19424
libqedislib.a		recvthread_log		tmp_qdb_file_19457
BertdeMacBook-Pro:Debug bertyoung$ cd
BertdeMacBook-Pro:~ bertyoung$ ls
Applications   Downloads      Music          VirtualBox VMs xcodedebug
Desktop        Library        Pictures       study
Documents      Movies         Public         ttt.cc
BertdeMacBook-Pro:~ bertyoung$ cd Do
-bash: cd: Do: No such file or directory
BertdeMacBook-Pro:~ bertyoung$ cd Downloads/
BertdeMacBook-Pro:Downloads bertyoung$ ls
AppCleaner.app                        dmg
BattleCity1990                        jdk-8u45-macosx-x64.dmg
BattleCity1990.zip                    leveldb实现解析.pdf
CocosStudioForMac-v2.0.0.0-Beta0.dmg  mario-resource.zip
MySuperMario                          marioresource
MySuperMario.zip                      redis-2.6.17
TileGame3                             scala-2.11.6
TileGame3.zip                         scala-2.11.6.tar
adt-bundle-mac-x86_64-20140702        skiplist_bert.cc
android-ndk-r9c                       tiled-0.10.2.dmg
back_up                               toBeDeleted
books                                 健身
clienttest.cc                         即通平台部okochayang(杨强).pdf
cocos2d-x-3.1
BertdeMacBook-Pro:Downloads bertyoung$ cd scala-2.11.6
BertdeMacBook-Pro:scala-2.11.6 bertyoung$ ls
bin doc lib man
BertdeMacBook-Pro:scala-2.11.6 bertyoung$ cd bin/
BertdeMacBook-Pro:bin bertyoung$ ls
ch1      fsc      scala    scalac   scaladoc scalap
BertdeMacBook-Pro:bin bertyoung$ ./scala
Welcome to Scala version 2.11.6 (Java HotSpot(TM) 64-Bit Server VM, Java 1.8.0_45).
Type in expressions to have them evaluated.
Type :help for more information.

scala> println()


scala> println("fdfds")
fdfds

scala> quit
<console>:8: error: not found: value quit
              quit
              ^

scala> quit()
<console>:8: error: not found: value quit
              quit()
              ^

scala> :quit
BertdeMacBook-Pro:bin bertyoung$ vi 2.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 2.scala 
825152896
BertdeMacBook-Pro:bin bertyoung$ vi 2.scala

KNo manual entry for Hello

shell returned 1

Press ENTER or type command to continue
No manual entry for Hello

shell returned 1

Press ENTER or type command to continue
BertdeMacBook-Pro:bin bertyoung$ ./scala 2.scala 
/Users/bertyoung/Downloads/scala-2.11.6/bin/2.scala:7: error: type mismatch;
 found   : Long
 required: Int
    res
    ^
one error found
BertdeMacBook-Pro:bin bertyoung$ vi 2.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 2.scala 
9415087488
BertdeMacBook-Pro:bin bertyoung$ vi 2.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 2.scala 
9415087488
9415087488
BertdeMacBook-Pro:bin bertyoung$ vi 2.scala
BertdeMacBook-Pro:bin bertyoung$ vi 2.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 2.scala 
9415087488
9415087488
Hello
BertdeMacBook-Pro:bin bertyoung$ vi 2.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 2.scala 
9415087488
9415087488
9415087488
BertdeMacBook-Pro:bin bertyoung$ vi 2.scala
BertdeMacBook-Pro:bin bertyoung$ vi 2.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 2.scala 
9415087488
9415087488
9415087488
25.0
0.04
BertdeMacBook-Pro:bin bertyoung$ vi 2.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 2.scala 
9415087488
9415087488
9415087488
25.0
0.04
BertdeMacBook-Pro:bin bertyoung$ vi 2.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 2.scala 
9415087488
9415087488
9415087488
1.0
0.04
25.0
0.008
125.0
BertdeMacBook-Pro:bin bertyoung$ vi 2.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 2.scala 
9415087488
9415087488
9415087488
1.0
0.25
4.0
0.125
8.0
BertdeMacBook-Pro:bin bertyoung$ vi 2.scala
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
java.lang.ArrayIndexOutOfBoundsException: 5
	at Main$$anon$1$$anonfun$randArray$1.apply$mcVI$sp(3.scala:6)
	at scala.collection.immutable.Range.foreach$mVc$sp(Range.scala:166)
	at Main$$anon$1.randArray(3.scala:5)
	at Main$$anon$1.<init>(3.scala:10)
	at Main$.main(3.scala:1)
	at Main.main(3.scala)
	at sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)
	at sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)
	at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.lang.reflect.Method.invoke(Method.java:497)
	at scala.reflect.internal.util.ScalaClassLoader$$anonfun$run$1.apply(ScalaClassLoader.scala:70)
	at scala.reflect.internal.util.ScalaClassLoader$class.asContext(ScalaClassLoader.scala:31)
	at scala.reflect.internal.util.ScalaClassLoader$URLClassLoader.asContext(ScalaClassLoader.scala:101)
	at scala.reflect.internal.util.ScalaClassLoader$class.run(ScalaClassLoader.scala:70)
	at scala.reflect.internal.util.ScalaClassLoader$URLClassLoader.run(ScalaClassLoader.scala:101)
	at scala.tools.nsc.CommonRunner$class.run(ObjectRunner.scala:22)
	at scala.tools.nsc.ObjectRunner$.run(ObjectRunner.scala:39)
	at scala.tools.nsc.CommonRunner$class.runAndCatch(ObjectRunner.scala:29)
	at scala.tools.nsc.ObjectRunner$.runAndCatch(ObjectRunner.scala:39)
	at scala.tools.nsc.ScriptRunner.scala$tools$nsc$ScriptRunner$$runCompiled(ScriptRunner.scala:175)
	at scala.tools.nsc.ScriptRunner$$anonfun$runScript$1.apply(ScriptRunner.scala:192)
	at scala.tools.nsc.ScriptRunner$$anonfun$runScript$1.apply(ScriptRunner.scala:192)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1$$anonfun$apply$mcZ$sp$1.apply(ScriptRunner.scala:161)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1.apply$mcZ$sp(ScriptRunner.scala:161)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1.apply(ScriptRunner.scala:129)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1.apply(ScriptRunner.scala:129)
	at scala.tools.nsc.util.package$.trackingThreads(package.scala:43)
	at scala.tools.nsc.util.package$.waitingForThreads(package.scala:27)
	at scala.tools.nsc.ScriptRunner.withCompiledScript(ScriptRunner.scala:128)
	at scala.tools.nsc.ScriptRunner.runScript(ScriptRunner.scala:192)
	at scala.tools.nsc.ScriptRunner.runScriptAndCatch(ScriptRunner.scala:205)
	at scala.tools.nsc.MainGenericRunner.runTarget$1(MainGenericRunner.scala:67)
	at scala.tools.nsc.MainGenericRunner.run$1(MainGenericRunner.scala:87)
	at scala.tools.nsc.MainGenericRunner.process(MainGenericRunner.scala:98)
	at scala.tools.nsc.MainGenericRunner$.main(MainGenericRunner.scala:103)
	at scala.tools.nsc.MainGenericRunner.main(MainGenericRunner.scala)
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
java.lang.ArrayIndexOutOfBoundsException: 5
	at Main$$anon$1$$anonfun$randArray$1.apply$mcVI$sp(3.scala:6)
	at scala.collection.immutable.Range.foreach$mVc$sp(Range.scala:166)
	at Main$$anon$1.randArray(3.scala:5)
	at Main$$anon$1.<init>(3.scala:10)
	at Main$.main(3.scala:1)
	at Main.main(3.scala)
	at sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)
	at sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)
	at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.lang.reflect.Method.invoke(Method.java:497)
	at scala.reflect.internal.util.ScalaClassLoader$$anonfun$run$1.apply(ScalaClassLoader.scala:70)
	at scala.reflect.internal.util.ScalaClassLoader$class.asContext(ScalaClassLoader.scala:31)
	at scala.reflect.internal.util.ScalaClassLoader$URLClassLoader.asContext(ScalaClassLoader.scala:101)
	at scala.reflect.internal.util.ScalaClassLoader$class.run(ScalaClassLoader.scala:70)
	at scala.reflect.internal.util.ScalaClassLoader$URLClassLoader.run(ScalaClassLoader.scala:101)
	at scala.tools.nsc.CommonRunner$class.run(ObjectRunner.scala:22)
	at scala.tools.nsc.ObjectRunner$.run(ObjectRunner.scala:39)
	at scala.tools.nsc.CommonRunner$class.runAndCatch(ObjectRunner.scala:29)
	at scala.tools.nsc.ObjectRunner$.runAndCatch(ObjectRunner.scala:39)
	at scala.tools.nsc.ScriptRunner.scala$tools$nsc$ScriptRunner$$runCompiled(ScriptRunner.scala:175)
	at scala.tools.nsc.ScriptRunner$$anonfun$runScript$1.apply(ScriptRunner.scala:192)
	at scala.tools.nsc.ScriptRunner$$anonfun$runScript$1.apply(ScriptRunner.scala:192)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1$$anonfun$apply$mcZ$sp$1.apply(ScriptRunner.scala:161)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1.apply$mcZ$sp(ScriptRunner.scala:161)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1.apply(ScriptRunner.scala:129)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1.apply(ScriptRunner.scala:129)
	at scala.tools.nsc.util.package$.trackingThreads(package.scala:43)
	at scala.tools.nsc.util.package$.waitingForThreads(package.scala:27)
	at scala.tools.nsc.ScriptRunner.withCompiledScript(ScriptRunner.scala:128)
	at scala.tools.nsc.ScriptRunner.runScript(ScriptRunner.scala:192)
	at scala.tools.nsc.ScriptRunner.runScriptAndCatch(ScriptRunner.scala:205)
	at scala.tools.nsc.MainGenericRunner.runTarget$1(MainGenericRunner.scala:67)
	at scala.tools.nsc.MainGenericRunner.run$1(MainGenericRunner.scala:87)
	at scala.tools.nsc.MainGenericRunner.process(MainGenericRunner.scala:98)
	at scala.tools.nsc.MainGenericRunner$.main(MainGenericRunner.scala:103)
	at scala.tools.nsc.MainGenericRunner.main(MainGenericRunner.scala)
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:12: error: ';' expected but '#' found.
#println(randArray(5).mkString("(", "-", ")"))
^
one error found
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
java.lang.ArrayIndexOutOfBoundsException: 5
	at Main$$anon$1$$anonfun$randArray$1.apply$mcVI$sp(3.scala:8)
	at scala.collection.immutable.Range.foreach$mVc$sp(Range.scala:166)
	at Main$$anon$1.randArray(3.scala:7)
	at Main$$anon$1.<init>(3.scala:11)
	at Main$.main(3.scala:1)
	at Main.main(3.scala)
	at sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)
	at sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)
	at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.lang.reflect.Method.invoke(Method.java:497)
	at scala.reflect.internal.util.ScalaClassLoader$$anonfun$run$1.apply(ScalaClassLoader.scala:70)
	at scala.reflect.internal.util.ScalaClassLoader$class.asContext(ScalaClassLoader.scala:31)
	at scala.reflect.internal.util.ScalaClassLoader$URLClassLoader.asContext(ScalaClassLoader.scala:101)
	at scala.reflect.internal.util.ScalaClassLoader$class.run(ScalaClassLoader.scala:70)
	at scala.reflect.internal.util.ScalaClassLoader$URLClassLoader.run(ScalaClassLoader.scala:101)
	at scala.tools.nsc.CommonRunner$class.run(ObjectRunner.scala:22)
	at scala.tools.nsc.ObjectRunner$.run(ObjectRunner.scala:39)
	at scala.tools.nsc.CommonRunner$class.runAndCatch(ObjectRunner.scala:29)
	at scala.tools.nsc.ObjectRunner$.runAndCatch(ObjectRunner.scala:39)
	at scala.tools.nsc.ScriptRunner.scala$tools$nsc$ScriptRunner$$runCompiled(ScriptRunner.scala:175)
	at scala.tools.nsc.ScriptRunner$$anonfun$runScript$1.apply(ScriptRunner.scala:192)
	at scala.tools.nsc.ScriptRunner$$anonfun$runScript$1.apply(ScriptRunner.scala:192)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1$$anonfun$apply$mcZ$sp$1.apply(ScriptRunner.scala:161)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1.apply$mcZ$sp(ScriptRunner.scala:161)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1.apply(ScriptRunner.scala:129)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1.apply(ScriptRunner.scala:129)
	at scala.tools.nsc.util.package$.trackingThreads(package.scala:43)
	at scala.tools.nsc.util.package$.waitingForThreads(package.scala:27)
	at scala.tools.nsc.ScriptRunner.withCompiledScript(ScriptRunner.scala:128)
	at scala.tools.nsc.ScriptRunner.runScript(ScriptRunner.scala:192)
	at scala.tools.nsc.ScriptRunner.runScriptAndCatch(ScriptRunner.scala:205)
	at scala.tools.nsc.MainGenericRunner.runTarget$1(MainGenericRunner.scala:67)
	at scala.tools.nsc.MainGenericRunner.run$1(MainGenericRunner.scala:87)
	at scala.tools.nsc.MainGenericRunner.process(MainGenericRunner.scala:98)
	at scala.tools.nsc.MainGenericRunner$.main(MainGenericRunner.scala:103)
	at scala.tools.nsc.MainGenericRunner.main(MainGenericRunner.scala)
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(0-3-0-4-0)
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:16: error: not found: value a
    for (i <- 0 until (a.length, 2)) {
                       ^
one error found
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ls
2.scala  3.scala  ch1      fsc      scala    scalac   scaladoc scalap
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ vi 2.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:16: error: not found: value a
    for (i <- 0 until (a.length, 2)) {
                       ^
one error found
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(3-1-2-4-2)
5
[I@58372a00
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:23: error: not found: value *
println(*arr)
        ^
one error found
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:23: error: not found: value arr_*
println(arr_*)
        ^
one error found
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:23: error: no `: _*' annotation allowed here
(such annotations are only allowed in arguments to *-parameters)
println(arr : _*)
            ^
one error found
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(4-1-2-0-4)
5
[I@58372a00
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:23: error: type mismatch;
 found   : Char(',')
 required: String
println(arr.mkString(','))
                     ^
one error found
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(0-2-4-3-3)
5
0,4,0,2,0
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(3-4-3-3-4)
5
0 - 3 - 0 - 1 - 0
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(1-2-4-1-1)
5
java.lang.ArrayIndexOutOfBoundsException: 5
	at Main$$anon$1$$anonfun$swapElem$1.apply$mcVI$sp(3.scala:18)
	at scala.collection.immutable.Range.foreach$mVc$sp(Range.scala:166)
	at Main$$anon$1.swapElem(3.scala:16)
	at Main$$anon$1.<init>(3.scala:24)
	at Main$.main(3.scala:1)
	at Main.main(3.scala)
	at sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)
	at sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)
	at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.lang.reflect.Method.invoke(Method.java:497)
	at scala.reflect.internal.util.ScalaClassLoader$$anonfun$run$1.apply(ScalaClassLoader.scala:70)
	at scala.reflect.internal.util.ScalaClassLoader$class.asContext(ScalaClassLoader.scala:31)
	at scala.reflect.internal.util.ScalaClassLoader$URLClassLoader.asContext(ScalaClassLoader.scala:101)
	at scala.reflect.internal.util.ScalaClassLoader$class.run(ScalaClassLoader.scala:70)
	at scala.reflect.internal.util.ScalaClassLoader$URLClassLoader.run(ScalaClassLoader.scala:101)
	at scala.tools.nsc.CommonRunner$class.run(ObjectRunner.scala:22)
	at scala.tools.nsc.ObjectRunner$.run(ObjectRunner.scala:39)
	at scala.tools.nsc.CommonRunner$class.runAndCatch(ObjectRunner.scala:29)
	at scala.tools.nsc.ObjectRunner$.runAndCatch(ObjectRunner.scala:39)
	at scala.tools.nsc.ScriptRunner.scala$tools$nsc$ScriptRunner$$runCompiled(ScriptRunner.scala:175)
	at scala.tools.nsc.ScriptRunner$$anonfun$runScript$1.apply(ScriptRunner.scala:192)
	at scala.tools.nsc.ScriptRunner$$anonfun$runScript$1.apply(ScriptRunner.scala:192)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1$$anonfun$apply$mcZ$sp$1.apply(ScriptRunner.scala:161)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1.apply$mcZ$sp(ScriptRunner.scala:161)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1.apply(ScriptRunner.scala:129)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1.apply(ScriptRunner.scala:129)
	at scala.tools.nsc.util.package$.trackingThreads(package.scala:43)
	at scala.tools.nsc.util.package$.waitingForThreads(package.scala:27)
	at scala.tools.nsc.ScriptRunner.withCompiledScript(ScriptRunner.scala:128)
	at scala.tools.nsc.ScriptRunner.runScript(ScriptRunner.scala:192)
	at scala.tools.nsc.ScriptRunner.runScriptAndCatch(ScriptRunner.scala:205)
	at scala.tools.nsc.MainGenericRunner.runTarget$1(MainGenericRunner.scala:67)
	at scala.tools.nsc.MainGenericRunner.run$1(MainGenericRunner.scala:87)
	at scala.tools.nsc.MainGenericRunner.process(MainGenericRunner.scala:98)
	at scala.tools.nsc.MainGenericRunner$.main(MainGenericRunner.scala:103)
	at scala.tools.nsc.MainGenericRunner.main(MainGenericRunner.scala)
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(3-2-0-0-3)
5
java.lang.ArrayIndexOutOfBoundsException: 5
	at Main$$anon$1$$anonfun$swapElem$1.apply$mcVI$sp(3.scala:18)
	at scala.collection.immutable.Range.foreach$mVc$sp(Range.scala:166)
	at Main$$anon$1.swapElem(3.scala:16)
	at Main$$anon$1.<init>(3.scala:24)
	at Main$.main(3.scala:1)
	at Main.main(3.scala)
	at sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)
	at sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)
	at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.lang.reflect.Method.invoke(Method.java:497)
	at scala.reflect.internal.util.ScalaClassLoader$$anonfun$run$1.apply(ScalaClassLoader.scala:70)
	at scala.reflect.internal.util.ScalaClassLoader$class.asContext(ScalaClassLoader.scala:31)
	at scala.reflect.internal.util.ScalaClassLoader$URLClassLoader.asContext(ScalaClassLoader.scala:101)
	at scala.reflect.internal.util.ScalaClassLoader$class.run(ScalaClassLoader.scala:70)
	at scala.reflect.internal.util.ScalaClassLoader$URLClassLoader.run(ScalaClassLoader.scala:101)
	at scala.tools.nsc.CommonRunner$class.run(ObjectRunner.scala:22)
	at scala.tools.nsc.ObjectRunner$.run(ObjectRunner.scala:39)
	at scala.tools.nsc.CommonRunner$class.runAndCatch(ObjectRunner.scala:29)
	at scala.tools.nsc.ObjectRunner$.runAndCatch(ObjectRunner.scala:39)
	at scala.tools.nsc.ScriptRunner.scala$tools$nsc$ScriptRunner$$runCompiled(ScriptRunner.scala:175)
	at scala.tools.nsc.ScriptRunner$$anonfun$runScript$1.apply(ScriptRunner.scala:192)
	at scala.tools.nsc.ScriptRunner$$anonfun$runScript$1.apply(ScriptRunner.scala:192)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1$$anonfun$apply$mcZ$sp$1.apply(ScriptRunner.scala:161)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1.apply$mcZ$sp(ScriptRunner.scala:161)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1.apply(ScriptRunner.scala:129)
	at scala.tools.nsc.ScriptRunner$$anonfun$withCompiledScript$1.apply(ScriptRunner.scala:129)
	at scala.tools.nsc.util.package$.trackingThreads(package.scala:43)
	at scala.tools.nsc.util.package$.waitingForThreads(package.scala:27)
	at scala.tools.nsc.ScriptRunner.withCompiledScript(ScriptRunner.scala:128)
	at scala.tools.nsc.ScriptRunner.runScript(ScriptRunner.scala:192)
	at scala.tools.nsc.ScriptRunner.runScriptAndCatch(ScriptRunner.scala:205)
	at scala.tools.nsc.MainGenericRunner.runTarget$1(MainGenericRunner.scala:67)
	at scala.tools.nsc.MainGenericRunner.run$1(MainGenericRunner.scala:87)
	at scala.tools.nsc.MainGenericRunner.process(MainGenericRunner.scala:98)
	at scala.tools.nsc.MainGenericRunner$.main(MainGenericRunner.scala:103)
	at scala.tools.nsc.MainGenericRunner.main(MainGenericRunner.scala)
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(1-1-2-4-4)
5
1 - 3 - 4 - 2 - 4
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(3-2-1-1-3)
2 - 1 - 4 - 3 - 5
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(1-3-2-4-4)
2-1-4-3-5
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:34: error: illegal start of simple expression
                yield a(i + 1)
                ^
one error found
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:32: error: illegal start of statement
        yield {
        ^
one error found
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:33: error: illegal start of statement
            yield {
            ^
one error found
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:36: error: not found: value a
                    a(i + 1)
                    ^
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:38: error: not found: value a
                    a(i - 1)
                    ^
two errors found
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(1-1-3-3-3)
2-1-4-3-5
2-1-4-3-()
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(1-3-0-4-0)
2-1-4-3-5
2-1-4-3-5
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(2-0-3-2-0)
2-1-4-3-5
2-1-4-3-5
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:78: error: ')' expected but eof found.

^
one error found
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:75: error: arr2 is already defined as value arr2
val arr2 = Array(1, 2, 3, 4, 5)
    ^
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:50: error: not found: value ArrayBuffer
    val posIndexes = ArrayBuffer[Int]()
                     ^
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:51: error: not found: value ArrayBuffer
    val otherIndexes = ArrayBuffer[Int]()
                       ^
three errors found
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:50: error: not found: value ArrayBuffer
    val posIndexes = ArrayBuffer[Int]()
                     ^
/Users/bertyoung/Downloads/scala-2.11.6/bin/3.scala:51: error: not found: value ArrayBuffer
    val otherIndexes = ArrayBuffer[Int]()
                       ^
two errors found
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(1-0-1-4-1)
2-1-4-3-5
2-1-4-3-5
1-2-3-4-5
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(1-0-0-0-2)
2-1-4-3-5
2-1-4-3-5
1-3--2-0--5
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(3-3-1-0-2)
2-1-4-3-5
2-1-4-3-5
1,3,-2,0,-5
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(4-1-1-3-2)
2-1-4-3-5
2-1-4-3-5
1,3,-2,0,-5
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(0-0-1-2-3)
2-1-4-3-5
2-1-4-3-5
1,3,-2,0,-5
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ ./scala 3.scala 
5
(0-4-2-3-1)
2-1-4-3-5
2-1-4-3-5
1,3,-2,0,-5
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 
BertdeMacBook-Pro:bin bertyoung$ vi 3.scala 

// 3.1 
import util.Random

def randArray(n : Int) = { 
    val a = new Array[Int](n)
    println(a.length)
    for (i <- 0 until a.length)
        a(i) = Random.nextInt(n)
    a   
}

println(randArray(5).mkString("(", "-", ")"))

// 3.2 
def swapElem(arr : Array[Int]) = { 
    for (i <- 0 until (arr.length, 2)) {
        if (i + 1 < arr.length) {
            val tmp = arr(i)
            arr(i) = arr(i + 1)
            arr(i + 1) = tmp 
        }   
    }   
}
:set nonumber                                                 1,1           Top
