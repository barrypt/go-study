package main

import (
	"fmt"
	"log"
	"strconv"
)

type LoadBalance interface {
    //选择一个后端Server
    //参数remove是需要排除选择的后端Server
    Select(remove []string) *Server
    //更新可用Server列表
    UpdateServers(servers []*Server)
}

type Server struct {
    //主机地址
    Host string
    //主机名		
    Name string		
    Weight int	
    //主机是否在线	
    Online bool		
}
type Weighted struct {
	Server 			*Server
	Weight          int
	CurrentWeight   int
	EffectiveWeight int
}
 
func (this *Weighted) String() string {
	return fmt.Sprintf("[%s][%d]",this.Server.Host,this.Weight)
}
 
type LoadBalanceWeightedRoundRobin struct{
	servers []*Server
	weighted []*Weighted
}
 
func NewLoadBalanceWeightedRoundRobin(servers []*Server) *LoadBalanceWeightedRoundRobin{
	new:=&LoadBalanceWeightedRoundRobin{}
	new.UpdateServers(servers)
	return new
}
 
func (this *LoadBalanceWeightedRoundRobin) UpdateServers(servers []*Server) {
	if len(this.servers)==len(servers) {
		for _,new:=range servers {
			isEqual:=false
			for _,old:=range this.servers {
				if new.Host==old.Host&&new.Weight==old.Weight&&new.Online==old.Online {
					isEqual=true
					break
				}
			}
			if isEqual==false {
				goto build
			}
		}
		return
	}
 
build:

	weighted:=make([]*Weighted,0)
	for _,v:=range servers {
		if v.Online==true {
			w:=&Weighted{
				Server:v,
				Weight:v.Weight,
				CurrentWeight:0,
				EffectiveWeight:v.Weight,
			}
			weighted=append(weighted,w)
		}
	}
	this.weighted=weighted
	this.servers=servers
	log.Default().Printf("weighted[%v]",this.weighted)
}
 
func (this *LoadBalanceWeightedRoundRobin) Select(remove []string) *Server {
	if len(this.weighted)==0 {
		return nil
	}
	w:=this.nextWeighted(this.weighted,remove)
	if w==nil {
		return nil
	}
	return w.Server
}
 
func (this *LoadBalanceWeightedRoundRobin) nextWeighted(servers []*Weighted,remove []string) (best *Weighted) {
	total := 0
	for i := 0; i < len(servers); i++ {
		w:= servers[i]
		if w == nil {
			continue
		}
		isFind:=false
		for _,v:=range remove {
			if v==w.Server.Host {
				isFind=true
			}
		}
		if isFind==true{
			continue
		}
 
		w.CurrentWeight += w.EffectiveWeight
		total += w.EffectiveWeight
		if w.EffectiveWeight < w.Weight {
			w.EffectiveWeight++
		}
 
		if best == nil || w.CurrentWeight > best.CurrentWeight {
			best = w
		}
	}
	if best == nil {
		return nil
	}
	best.CurrentWeight -= total
	return best
}
 
func (this *LoadBalanceWeightedRoundRobin) String() string {
	return "WeightedRoundRobin"
}


func main() {
	count:=make([]int,4)
	servers:=make([]*Server,0)
	servers=append(servers,&Server{Host:"0",Weight:10,Online:true})
	servers=append(servers,&Server{Host:"1",Weight:20,Online:true})
	servers=append(servers,&Server{Host:"2",Weight:30,Online:true})
	servers=append(servers,&Server{Host:"3",Weight:40,Online:true})
	lb:=NewLoadBalanceWeightedRoundRobin(servers)
 
	for i:=0;i<100000;i++{
		s:=lb.Select(nil)
		id,_:=strconv.Atoi(s.Host)
		count[id]++
	}
	fmt.Println(count)
}