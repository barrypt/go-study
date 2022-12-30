package main

import (
   "log"
   "net"
)

func checkErr(err error) {
   if err != nil {
      log.Fatal(err)
   }
}

func handleConn(c net.Conn) {
   defer c.Close()
   for {
      var buf = make([]byte, 10)
      log.Println("开始从conn读取：")

      n, err := c.Read(buf)
      checkErr(err)
      log.Printf("read %d bytes ,content is : %s\n", n, string(buf[:n]))
   }
}

func main() {
   //监听8080端口
   l, err := net.Listen("tcp", ":8080")
   checkErr(err)
   defer l.Close()
   log.Println("服务启动 成功")

   for {
      c, err := l.Accept()
      checkErr(err)
      //开启新协程处理连接
      go handleConn(c)
   }

}