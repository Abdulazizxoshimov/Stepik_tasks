type War struct{
    On bool
    Ammo int
    Power int 
}
func (s *War) Shoot() bool{
    if s.On == false{
        return false
    }else if s.Ammo>0{
        s.Ammo -= 1
        return true
    }
    return false
}
func (s *War) RideBike() bool{
    if s.On == false{
        return false
    }else if s.Power>0{
        s.Power -= 1
        return true
    }
    return false
}
func main() {
    testStruct := new(War)
}