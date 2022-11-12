/* ----------------------------------
*  @author suyame 2022-11-09 15:54:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package server

import "sync"

type Server struct {
	sync.RWMutex
	host    string
	pageNum int
	isAlive bool
	QPS     float64
}

func NewServer(host string) *Server {
	return &Server{
		host:    host,
		isAlive: true,
	}
}

func (s *Server) GetHost() string {
	return s.host
}

func (s *Server) Add() {
	s.Lock()
	defer s.Unlock()
	s.pageNum++
}
func (s *Server) Sub() {
	s.Lock()
	defer s.Unlock()
	s.pageNum--
}

func (s *Server) IsAlive() bool {
	s.RLock()
	defer s.RUnlock()
	return s.isAlive
}

func (s *Server) Crush() {
	s.Lock()
	defer s.Unlock()
	s.isAlive = false
}
