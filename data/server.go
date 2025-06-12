package data

type Server struct {
	Id     int64  `xorm:"pk autoincr" json:"id"`
	Name   string `xorm:"varchar(255) notnull" json:"name"`
	Config string `json:"config"`
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
	_, err := s.d.e.ID(nd.Id).Update(nd)
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
	err = s.d.Rule.Delete(&Rule{ServerId: nd.Id})
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
