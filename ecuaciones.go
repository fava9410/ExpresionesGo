package main

import (
	"fmt";
  "bufio";
  "os";
  "strings";
	"strconv"
)

type Arbol struct {
	Izquierda  *Arbol
	Valor string
	Derecha *Arbol
}

type NodeA struct {
	variable string
	arbol *Arbol
	valor string
}

func (n *NodeA) String() string {
	return fmt.Sprint(n.arbol)
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*NodeA
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *NodeA) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *NodeA {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

type Node struct {
	Variable string
  Ecuacion []string
}

func (n *Node) String() string {
	return fmt.Sprint(n.Variable , "->" ,n.Ecuacion)
}

func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func Operaciones(t *Arbol) int {
	if t == nil {
		return 0
	}
	switch t.Valor {
		case "+":
			return Operaciones(t.Izquierda) + Operaciones(t.Derecha)
		case "-":
			return Operaciones(t.Izquierda) - Operaciones(t.Derecha)
		case "*":
			return Operaciones(t.Izquierda) * Operaciones(t.Derecha)
		case "/":
			return Operaciones(t.Izquierda) / Operaciones(t.Derecha)
		default:
			numero,_ := strconv.Atoi(t.Valor)
			return numero
	}
	return 0
}

func sintaxis(t * Arbol) int{
	if t != nil{
		_,er := strconv.Atoi(t.Valor)
		if t.Izquierda == nil && t.Derecha == nil{
			if er != nil {
				fmt.Println("Se esperaba un numero en lugar de ",t.Valor)
				return sintaxis(t.Izquierda) + sintaxis(t.Derecha) + 1
			}
			return sintaxis(t.Izquierda) + sintaxis(t.Derecha)
		}	else if t.Izquierda != nil && t.Derecha != nil {
			if er != nil{
				if t.Valor=="*"||t.Valor=="/"||t.Valor=="+"||t.Valor=="-"{
					return sintaxis(t.Izquierda) + sintaxis(t.Derecha)
				} else {
					fmt.Println("Se esperaba un signo en lugar de ",t.Valor)
					return sintaxis(t.Izquierda) + sintaxis(t.Derecha) +1
				}
			} else {
					fmt.Println("Se esperaba un signo en lugar de un numero, en ",t.Valor)
					return sintaxis(t.Izquierda) + sintaxis(t.Derecha) + 1
			}
		}
	}
	return 0
}

func evaluarEcuacion(terminos [] string, variable string, arboles *Stack)  int {
	s := NewStack()
  for i := 0; i < len(terminos); i++ {
    if terminos[i] == "+" || terminos[i] == "-" || terminos[i] == "*" || terminos[i] == "/" {
      if len(s.nodes) < 2 {
        fmt.Println("La expresion de la variable "+ variable +" esta mal formada, por favor revisela e intente de nuevo")
        return 1
      } else {
        a := s.Pop()
        b := s.Pop()
          s.Push(&NodeA{variable, &Arbol{b.arbol,terminos[i],a.arbol}, "0"})
      }
    } else {
			cambiovariable := false
			for j := 0; j < len(arboles.nodes); j++ {
					if (terminos[i] == arboles.nodes[j].variable){
						s.Push(&NodeA{variable, &Arbol{nil,arboles.nodes[j].valor,nil},"0"})
						cambiovariable = true
					}
			}
			if cambiovariable == false{
				s.Push(&NodeA{variable, &Arbol{nil,terminos[i],nil},"0"})
			}
		}
  }
	arbolfinal := s.Pop()
	arboles.Push(arbolfinal)
  return 0
}

func main() {
	arboles := NewStack()
	ecuaciones := NewQueue(1)
	for (true){
		sc := bufio.NewScanner(os.Stdin)
	  sc.Scan()
	  ecuacion := sc.Text()
		if ecuacion == "" {
			break
		} else {
			terminos := strings.Split(ecuacion," ")
			ecuaciones.Push(&Node{terminos[len(terminos)-1], append(terminos[:len(terminos)-1])})
		}
	}

	for i := 0; i < len(ecuaciones.nodes); i++ {
		ee := evaluarEcuacion(ecuaciones.nodes[i].Ecuacion, ecuaciones.nodes[i].Variable, arboles)
		if ee == 0{
	    t2 := arboles.Pop()
	    arbolt2 := t2.arbol
	    sin := sintaxis(arbolt2)
	  	if (sin == 0){
				valor := Operaciones(arbolt2)
				valors := strconv.Itoa(valor)
	  		fmt.Println("El valor de la variable "+ t2.variable +" es: "+ valors)
	  		arboles.Push(&NodeA{t2.variable, arbolt2, valors})
	  	}
	  }
	}
}
