package centralsim

import (
	"log"
)

//--------------------------------------------------------------------------
// TransitionList is a list of transitions themselves
type TransitionList []Transition //Slice de transiciones como Lista

// length return length of TransitionList with type adapted to IndLocalTrans
func (self TransitionList) Length() IndLocalTrans {
	return IndLocalTrans(len(self))
}

//--------------------------------------------------------------------------

// IndGlobalTrans is a index of a transition in the global list
type IndGlobalTrans int32

// IndLocalTrans is a index of a transition in the local lefs list
type IndLocalTrans int32

//TypeConst is the constant to propagate in lefs
type TypeConst int32

type TransitionConstant struct {
	INextTrans IndLocalTrans
	Cnstnt     TypeConst
}

//------------------------------------------------------------------------

// -----------------------------------------------------------------------
// Tipo abstracto Transition para guardar la informacion de una transicion
// -----------------------------------------------------------------------
type Transition struct {
	// indice en la tabla global de transiciones
	IdLocal IndLocalTrans

	// iiValorLef es el valor que tiene la funcion de
	// sensibilizacion en el instante de tiempo que nos da
	// la variable ITime
	IiValorLef TypeConst
	ITime      TypeClock

	// tiempo que dura el disparo de la transicion
	IiShotDuration TypeClock

	// vector de transiciones a las que tengo que propagar cte
	// cuando se dispare esta transicion, junto con la cte que
	// tengo que propagar
	IiListactes []TransitionConstant
}

/*
	-----------------------------------------------------------------
	   METODO: PrintEvent
	   RECIBE: Nada
	   DEVUELVE: Nada
	   PROPOSITO: Imprimir los atributos de la clase para depurar errores
		-----------------------------------------------------------------
*/
func (self *Transition) PrintEvent() {
	log.Println("Dato Transicion:")
	log.Println("IDGLOBAL: ", self.IdLocal)
	log.Println(" VALOR LEF: ", self.IiValorLef)
	log.Println(" TIEMPO: ", self.ITime)
	log.Println(" DURACION DISPARO: ", self.IiShotDuration)
	log.Println(" LISTA DE CTES: ")
	for _, v := range self.IiListactes {
		log.Println("\tTRANSICION: ", v.INextTrans, "\t\tCTE: ", v.Cnstnt)
	}
}

/*
	-----------------------------------------------------------------
   METODO: PrintEventValues
   RECIBE: Nada
   DEVUELVE: Nada
   PROPOSITO: Imprimir simplemente el valor de la transicion
	COMENTARIO : es solo de lectura
	-----------------------------------------------------------------
*/
func (self Transition) PrintEventValues() {
	log.Println("Transicion -> ")
	//	log.Println("\tIDGLOBAL: ", self.Ii_idglobal)
	log.Println("\t\tVALOR LEF: ", self.IiValorLef)
	log.Println("\t\tTIEMPO: ", self.ITime)
}

//----------------------------------------------------------------------

// Stack Transition is a Stack of transitions indices
type StackTransitions []IndLocalTrans

//Push transition id to stack
func (self *StackTransitions) push(i_tr IndLocalTrans) {
	*self = append(*self, i_tr)
}

//Pop transition id from stack
func (self *StackTransitions) pop() IndLocalTrans {
	if (*self).isEmpty() {
		return -1
	} else {
		i_tr := (*self)[len(*self)-1]  // obtener dato de lo alto de la pila
		*self = (*self)[:len(*self)-1] //desempilar
		return i_tr
	}
}

func (self StackTransitions) isEmpty() bool {
	return len(self) == 0
}
