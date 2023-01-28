package stack

// IsValid ()匹配,s 取值只能为 (){}[]
func (a *ArrayStack)IsValid(s string) bool {
	temp := []byte(s)
	if len(temp)%2 != 0{
		return false
	}
	for _,v := range temp{
		switch v {
		case '(':
			a.Push(')')
		case '{':
			a.Push('}')
		case '[':
			a.Push(']')
		case ')':
			if a.Pop() != ')'{
				return false
			}
		case ']':
			if a.Pop() != ']'{
				return false
			}
		case '}':
			if a.Pop() != '}'{
				return false
			}
		default:
			return false
		}
	}
	if a.top != -1{
		return false
	}
	return true
}

// RemoveDuplicates 删除字符串中的所有相邻重复项
func (a *ArrayStack)RemoveDuplicates(s string) string {
	temp := []rune(s)
	for _,v := range temp{
		if a.Top() == nil || v != a.Top().(rune){
			a.Push(v)
		}else{
			a.Pop()
		}
	}
    if a.IsEmpty(){
    	return ""
	}
	return a.GetAllDataInString()
}



