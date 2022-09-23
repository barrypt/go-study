package main

import (
	"bufio"
	"context"
	"os"
	"strings"
	"time"

	pb "client/chartroom"

	"github.com/pterm/pterm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	address = "localhost:23333"
)

func main() {
	/* ---------------------------------- 连接服务器 --------------------------------- */
	spinner, _ := pterm.DefaultSpinner.Start("正在连接聊天室")
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		spinner.Fail("连接失败")
		pterm.Fatal.Printfln("无法连接至服务器: %v", err)
		return
	}
	c := pb.NewChatRoomClient(conn)
	spinner.Success("连接成功")
	/* ---------------------------------- 注册用户名 --------------------------------- */
	var val *wrapperspb.StringValue
	var user *pb.User
	for {
		result, _ := pterm.DefaultInteractiveTextInput.Show("创建用户名")
		if strings.TrimSpace(result) == "" {
			pterm.Error.Printfln("进入聊天室失败,没有取名字")
			continue
		}
		user = &pb.User{Name: result}
		val, err = c.Login(context.TODO(), user)
		if err != nil {
			pterm.Error.Printfln("进入聊天室失败 %v", err)
			continue
		} else {
			break
		}
	}
	user.Id = val.Value
	pterm.Success.Println("创建成功！开始聊天吧！")
	/* ---------------------------------- 聊天室逻辑 --------------------------------- */
	stream, _ := c.Chat(metadata.AppendToOutgoingContext(context.Background(), "uuid", user.Id))
	go func(pb.ChatRoom_ChatClient) {
		for {
			res, _ := stream.Recv()
			switch res.Id {
			case "server":
				pterm.Success.Printfln("(%[2]v) [服务器] %[1]s ", res.Content, time.Unix(int64(res.Time), 0).Format(time.ANSIC))
			default:
				pterm.Info.Printfln("(%[3]v) %[1]s : %[2]s", res.Name, res.Content, time.Unix(int64(res.Time), 0).Format(time.ANSIC))
			}
		}
	}(stream)
	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		input = strings.TrimRight(input, "\r \n")
		// pterm.Info.Printfln("%s : %s", user.Name, input)
		stream.Send(&pb.ChatMessage{Id: user.Id, Content: input})
	}
}
