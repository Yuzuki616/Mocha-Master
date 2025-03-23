package data

func (r *RuleFunc) CreateTun(nd *Rule, nd2 *Rule) error {

	// Create out rule
	err := r.Create(nd2)
	if err != nil {
		return err
	}
	// Create in rule
	nd.TargetRule = nd2.Id
	nd.TargetType = TunInType
	nd2.TargetType = TunOutType
	err = r.Create(nd)
	if err != nil {
		return err
	}
	return nil
}

func (r *RuleFunc) UpdateTun(nd *Rule, nd2 *Rule) error {
	// Update out rule
	err := r.Update(nd2)
	if err != nil {
		return err
	}
	// Update in rule
	nd.TargetRule = nd2.Id
	nd.TargetType = TunInType
	nd2.TargetType = TunOutType
	err = r.Update(nd)
	if err != nil {
		return err
	}
	return nil
}
