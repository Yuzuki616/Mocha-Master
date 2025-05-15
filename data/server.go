package data

import "github.com/goccy/go-json"

type Server struct {
	Id        int64           `xorm:"pk autoincr"`
	Name      string          `xorm:"varchar(255) notnull unique"`
	PortRange [2]int          `xorm:"varchar(255) notnull unique"`
	Config    json.RawMessage `xorm:"json"`
}

type ServerFunc struct {
	d *Data
}

func (s *ServerFunc) Create(nd *Server) error {
	_, err := s.d.e.Insert(nd)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerFunc) Update(nd *Server) error {
	_, err := s.d.e.Update(nd)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerFunc) Delete(nd *Server) error {
	_, err := s.d.e.Delete(nd)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerFunc) Get(nd *Server) error {
	_, err := s.d.e.Get(nd)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerFunc) List() ([]Server, error) {
	var servers []Server
	err := s.d.e.Find(&servers)
	if err != nil {
		return nil, err
	}
	return servers, nil
}

func (s *ServerFunc) IsExist(sv *Server) (bool, error) {
	return s.d.e.Exist(sv)
}
