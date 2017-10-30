package main

type USB interface {
	Name() string
	Connter
}

type Connter interface{
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func Disconnect(usb interface{}) {
	v:= usb.(type){
	case PhoneConnecter:
		fmt.Println("Disconnect",v.name)
	default:
		fmt.Println("UNKNOW Device")
		}
	
}

func main() {
	var a interface{}
	fmt.Println(a == nil)
	
	var p *int = nil
	a = p
	fmt.Println(a == nil)
	
	var p USB
	p = PhoneConnecter{"ddd"}
	p.Connect()
	Disconnect(p)
}
