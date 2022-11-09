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
	port    string
	pageNum int
	isAlive bool
	QPS     float64
}

func NewServer(host string, port string) *Server {
	return &Server{
		host: host,
		port: port,
	}
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
