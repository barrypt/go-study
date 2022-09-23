package server

import (
	"context"
	"fmt"
	"sync"
	"time"

     Pb "server/chartroom"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// ^ 实现服务
type Service struct {
	Pb.UnimplementedChatRoomServer
	chatMessageCache []*Pb.ChatMessage
	userMap          sync.Map
	L                sync.RWMutex
}

var (
	workers map[Pb.ChatRoom_ChatServer]Pb.ChatRoom_ChatServer = make(map[Pb.ChatRoom_ChatServer]Pb.ChatRoom_ChatServer)
)

// ^ 实现login用户注册方法
func (s *Service) Login(ctx context.Context, in *Pb.User) (*wrappers.StringValue, error) {
	in.Id = uuid.New().String()
	if _, ok := s.userMap.Load(in.Id); ok {
		return nil, status.Errorf(codes.AlreadyExists, "已有同名用户,请换个用户名")
	}
	s.userMap.Store(in.Id, in)
	go s.sendMessage(nil, &Pb.ChatMessage{Id: "server", Content: fmt.Sprintf("%v 加入聊天室", in.Name), Time: uint64(time.Now().Unix())})
	// some work...
	return &wrappers.StringValue{Value: in.Id}, status.New(codes.OK, "").Err()
}

// ^ 实现聊天室
func (s *Service) Chat(stream Pb.ChatRoom_ChatServer) error {
	if s.chatMessageCache == nil {
		s.chatMessageCache = make([]*Pb.ChatMessage, 0, 1024)
	}
	workers[stream] = stream
	for _, v := range s.chatMessageCache {
		stream.Send(v)
	}
	s.recvMessage(stream)
	return status.New(codes.OK, "").Err()
}

func (s *Service) recvMessage(stream Pb.ChatRoom_ChatServer) {
	md, _ := metadata.FromIncomingContext(stream.Context())
	for {
		mes, err := stream.Recv()
		if err != nil {
			s.L.Lock()
			delete(workers, stream)
			s.L.Unlock()
			s.userMap.Delete(md.Get("uuid")[0])
			fmt.Println("某个用户掉线,目前用户在线数量", len(workers))
			break
		}
		s.chatMessageCache = append(s.chatMessageCache, mes)
		v, ok := s.userMap.Load(md.Get("uuid")[0])
		if !ok {
			fmt.Println("致命错误,用户不存在")
			return
		}
		mes.Name = v.(*Pb.User).Name
		mes.Time = uint64(time.Now().Unix())
		s.sendMessage(stream, mes)
	}
}

func (s *Service) sendMessage(stream Pb.ChatRoom_ChatServer, mes *Pb.ChatMessage) {
	s.L.Lock()
	for _, v := range workers {
		if v != stream {
			err := v.Send(mes)
			if err != nil {
				// err handle
				continue
			}
		}
	}
	s.L.Unlock()
}
