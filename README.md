**slow未完成**
感觉这个和跑得机器的性能是不是也有点关系。
1. 我看了下test的代码。它是将slow client的启动时间延后一段时间来进行时延的。此外，在进行slow测试的时候，会在发送正常消息之前额外发送一段largeByteSlice来增加丢包的概率，我觉得他这个是不稳定的时延测试。
2.有人提出在程序中打log的方式。我觉得这个方式实质上和上面发送largeByteSlice的想法是一样。就是通过增加在io方面的消耗（这是一个是屏幕上的io，一个是网络数据的io）来人为的降低程序的性能。我在github上找到的另外三个答案，都在程序运行的时候打log到屏幕上。
3.我在我的win10笔记本上跑博主的代码的时候，在5个客户端的性能测试时就fail了。然后在后面的slow测试中就出现的结果是normal client的收消息出现了丢包。然后我优化了一下性能。在通过前面的性能测试后，也出现了不drop包的情况。而在我的服务器上跑得时候，直接就是出现不drop包的情况。
结合以上三点，我觉得这个问题出现的根本原因应该是性能方面的问题。解决方案可能真的是打log或者其他降低性能的方法，把性能降低到能够被largeByteSlice拖到超时但是又不至于卡性能测试就行了。

---
p0
==

This repository contains the starter code that you will use as the basis of your multi-client
echo server implementation. It also contains the tests that we will use to test your implementation,
and an example 'server runner' binary that you might find useful for your own testing purposes.

If at any point you have any trouble with building, installing, or testing your code, the article
titled [How to Write Go Code](http://golang.org/doc/code.html) is a great resource for understanding
how Go workspaces are built and organized. You might also find the documentation for the
[`go` command](http://golang.org/cmd/go/) to be helpful. As always, feel free to post your questions
on Piazza as well.

## Running the official tests

To test your submission, we will execute the following command from inside the
`src/github.com/cmu440/p0` directory:

```sh
$ go test
```

We will also check your code for race conditions using Go's race detector by executing
the following command:

```sh
$ go test -race
```

To execute a single unit test, you can use the `-test.run` flag and specify a regular expression
identifying the name of the test to run. For example,

```sh
$ go test -race -test.run TestBasic1
```

## Testing your implementation using `srunner`

To make testing your server a bit easier (especially during the early stages of your implementation
when your server is largely incomplete), we have given you a simple `srunner` (server runner)
program that you can use to create and start an instance of your `MultiEchoServer`. The program
simply creates an instance of your server, starts it on a default port, and blocks forever,
running your server in the background.

To compile and build the `srunner` program into a binary that you can run, execute the three
commands below (these directions assume you have cloned this repo to `$HOME/p0`):

```bash
$ export GOPATH=$HOME/p0
$ go install github.com/cmu440/srunner
$ $GOPATH/bin/srunner
```

The `srunner` program won't be of much use to you without any clients. It might be a good exercise
to implement your own `crunner` (client runner) program that you can use to connect with and send
messages to your server. We have provided you with an unimplemented `crunner` program that you may
use for this purpose if you wish. Whether or not you decide to implement a `crunner` program will not
affect your grade for this project.

You could also test your server using Netcat as you saw shortly in lecture (i.e. run the `srunner`
binary in the background, execute `nc localhost 9999`, type the message you wish to send, and then
click enter).

## Using Go on AFS

For those students who wish to write their Go code on AFS (either in a cluster or remotely), you will
need to set the `GOROOT` environment variable as follows (this is required because Go is installed
in a custom location on AFS machines):

```bash
$ export GOROOT=/usr/local/lib/go
```
