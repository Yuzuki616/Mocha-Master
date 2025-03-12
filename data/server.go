package data

import "xorm.io/xorm"

type Server struct {
	Id   int64  `xorm:"pk autoincr"`
	Name string `xorm:"varchar(255) notnull unique"`
	Ext  map[string]interface{}
}

type ServerFunc struct {
	*xorm.Engine
}

func (s *ServerFunc) Create(nd *Server) error {
	_, err := s.Engine.Insert(nd)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerFunc) Update(nd *Server) error {
	_, err := s.Engine.Update(nd)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerFunc) Delete(nd *Server) error {
	_, err := s.Engine.Delete(nd)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerFunc) Get(nd *Server) error {
	_, err := s.Engine.Get(nd)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerFunc) List() ([]Server, error) {
	var servers []Server
	err := s.Engine.Find(&servers)
	if err != nil {
		return nil, err
	}
	return servers, nil
}

func (s *ServerFunc) IsExist(sv *Server) (bool, error) {
	return s.Engine.Exist(sv)
}
