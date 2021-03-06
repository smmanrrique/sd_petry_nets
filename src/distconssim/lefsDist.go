package distconssim

import (
	"log"

	u "sd_petry_nets/src/utils"
)

// Incidence is a dictionary of Global ID AND connection of Pre, Post
type Incidence map[IndGlobalTrans]u.Connect

// Active check if get all lookout before init
var Active bool = false

// LefsDist es el tipo de datos principal que gestiona el disparo de transiciones.
type LefsDist struct {
	Post, Pre     Incidence            //
	Lookout       map[string]TypeClock //
	SubNet        TransitionList       // Slice de transiciones de esta subred
	IsTransSensib StackTransitions     // Identificadores de las transiciones sensibilizadas para
	IlEvents      EventList            //Lista de eventos a procesar

}

/*
-----------------------------------------------------------------
   METODO: NewLefsDist
   RECIBE: Lista de transiciones
   DEVUELVE: Nada
   PROPOSITO: crear nueva estructura LEF
-----------------------------------------------------------------
*/
func NewLefsDist(listaTransiciones TransitionList) LefsDist {
	l := LefsDist{}
	l.SubNet = listaTransiciones
	l.IsTransSensib = nil
	l.IlEvents = nil
	return l
}

/*
-----------------------------------------------------------------
   METODO: AddEvents
   RECIBE: Evento a a�adir
   DEVUELVE: OK si todo va bien o ERROR en caso contrario
   PROPOSITO: A�ade a la lista de eventos
-----------------------------------------------------------------
*/
func (self *LefsDist) AddEvents(ae_evento EventDist) bool {
	self.IlEvents.Insert(ae_evento)
	return true
}

/*
-----------------------------------------------------------------
   METODO: SetLookout
   RECIBE: IP who send msm and Type of subred
   DEVUELVE: Nothing
   PROPOSITO: Set to subred lookahead value
-----------------------------------------------------------------
*/
func (self *LefsDist) SetLookout(str string, tim TypeClock) {
	if self.Lookout == nil {
		self.Lookout = map[string]TypeClock{}
	}

	self.Lookout[str] = tim
}

/*
-----------------------------------------------------------------
   METODO: AddSensitive
   RECIBE: Transicion sensibilizada a a�adir
   DEVUELVE: OK si todo va bien o ERROR en caso contrario
   PROPOSITO: A�ade a la lista de transiciones sensibilizadas
-----------------------------------------------------------------
*/
func (self *LefsDist) AddSensitive(ai_transicion IndGlobalTrans) bool {
	self.IsTransSensib.push(ai_transicion)
	return true // OK
}

/*
-----------------------------------------------------------------
   METODO: TimeFirstEvent
   RECIBE: Nada
   DEVUELVE: El valor del tiempo del primer evento de la lista de eventos.
	  -1 si ocurrio un error o no hay eventos.
   PROPOSITO: Visualizar el valor temporal del primer evento para conocer
	   posteriormente si debemos avanzar el reloj local
-----------------------------------------------------------------
*/
func (self LefsDist) TimeFirstEvent() TypeClock {
	if self.IlEvents.Length() > 0 {
		le_evento := self.IlEvents.GetFirstEvent()
		return le_evento.ITime
	} else {
		return -1
	}
}

/*
-----------------------------------------------------------------
   METODO: CheckLookout
   RECIBE: Nothing
   DEVUELVE: Nothing
   PROPOSITO: Checking if it has all lookahead
-----------------------------------------------------------------
*/
func (self *LefsDist) CheckLookout() {
	Active = true
	for _, lookahead := range self.Lookout {
		if lookahead < 0 {
			Active = false
			break
		}
	}
}

/*
-----------------------------------------------------------------
   METODO: MinTime
   RECIBE: Nothing
   DEVUELVE: TypeClock
   PROPOSITO: Search min time snapshot into lookout
-----------------------------------------------------------------
*/
func (self *LefsDist) MinTime() TypeClock {
	var min TypeClock
	for _, addr := range self.Pre {
		min = self.Lookout[addr.GetIDSubRed()]
		break
	}
	for _, tim := range self.Lookout {
		if tim < min {
			min = tim
		}
	}
	return min
}

/*
-----------------------------------------------------------------
   METODO: TimeDuration
   RECIBE: IdGlbal of transition
   DEVUELVE: TypeClock
   PROPOSITO: calculate shooting time
-----------------------------------------------------------------
*/
func (self *LefsDist) TimeDuration(id IndGlobalTrans) TypeClock {
	for _, trans := range self.SubNet {
		if trans.IDGlobal == id {
			return trans.IiShotDuration
		}
	}
	// Chequear este caso
	// log.Println("******************ERROR***************************")
	// return TypeClock(1)
	return TypeClock(0)
}

/*
-----------------------------------------------------------------
   METODO: ThereEvent
   RECIBE: Tiempo del reloj local
   DEVUELVE: true si quedan eventos para ese tiempo o false en caso contrario
   PROPOSITO: Conocer si restan eventos disponibles para el tiempo dado
-----------------------------------------------------------------
*/
func (self LefsDist) ThereEvent(ai_tiempo TypeClock) bool {
	if self.TimeFirstEvent() == ai_tiempo {
		return true
	} else {
		return false
	}
}

/*
-----------------------------------------------------------------
   METODO: ThereSensitive
   RECIBE: Nada
   DEVUELVE: true si las hay o false en caso contrario
   PROPOSITO: Conocer si tenemos funciones sensibilizadas
	COMENTARIOS: Se supone que previamente a la invocacion a esta funcion
	   se ha tenido que llamar a UpdateSensitive (reloj_local)
-----------------------------------------------------------------
*/
func (self LefsDist) ThereSensitive() bool {
	return !self.IsTransSensib.isEmpty()
}

/*
-----------------------------------------------------------------
   METODO: GetSensitive
   RECIBE: Nada
   DEVUELVE: El identificador de la primera transicion sensibilizada
	 o -1 en caso contrario
   PROPOSITO: Coger el primer identificador de la lista de transiciones
	 sensibilizadas
-----------------------------------------------------------------
*/
func (self *LefsDist) GetSensitive() IndGlobalTrans {
	if (*self).IsTransSensib.isEmpty() {
		return -1
	} else {
		return (*self).IsTransSensib.pop()
	}
}

/*
-----------------------------------------------------------------
   METODO: GetFirstEvent
   RECIBE: Nada
   DEVUELVE: El primer evento de la lista de eventos
   PROPOSITO: Coger el primer evento de la lista de eventos
-----------------------------------------------------------------
*/
func (self *LefsDist) GetFirstEvent() EventDist {
	le_evento := (*self).IlEvents.GetFirstEvent()
	(*self).IlEvents.DeleteFirstEvent()
	return le_evento
}

/*
-----------------------------------------------------------------
   METODO: UpdateSensitive
   RECIBE: Valor del reloj local actual para el que queremos saber las
	  transiciones sensibilizadas
   DEVUELVE: OK si todo fue bien o ERROR en caso contrario
   PROPOSITO: Que esta funcion sirva para recorrerse toda la lista de transiciones
	   e Inserttar aquellas en la pila de transiciones sensibilizadas.
COMENTARIOS: Me recorro todo el array de transiciones, por lo que deberiamos
	   invocar a esta funcion cuando ya hayan sido a�adidas todas las transiciones.
-----------------------------------------------------------------
*/
func (self *LefsDist) UpdateSensitive(ai_relojlocal TypeClock) bool {
	for li_i, t := range (*self).SubNet {
		if t.IiValorLef <= 0 && t.ITime == ai_relojlocal {
			(*self).IsTransSensib.push(IndGlobalTrans(li_i))
		}
	}
	return true
}

/*
-----------------------------------------------------------------
   METODO: UpdateTime
   RECIBE: Codigo de la transicion y nuevo valor del tiempo
   DEVUELVE: true si todo fue bien o false en caso contrario
   PROPOSITO: Modificar el tiempo de la transicion dada
-----------------------------------------------------------------
*/
func (self *LefsDist) UpdateTime(il_tr IndGlobalTrans, ai_ti TypeClock) bool {
	// Algunas comprobaciones...
	if il_tr >= 0 && il_tr < self.SubNet.Length() {
		// Modificacion del tiempo
		self.SubNet[il_tr].ITime = ai_ti
		return true
	} else { // index out of range
		return false
	}
}

/*
-----------------------------------------------------------------
   METODO: UpdateFuncValue
   RECIBE: Codigo de la transicion y valor con el que modificar
		OJO, no es el valor definitivo, sino la CTE a a�adir al valor que tenia
		antes la funcion
   DEVUELVE: OK si todo fue bien o ERROR en caso contrario
   PROPOSITO: Modificar valor de funcion de sensibilizacion de transicion dada
-----------------------------------------------------------------
*/
func (self *LefsDist) UpdateFuncValue(ilTr IndGlobalTrans, aiValLef TypeConst) bool {
	// Algunas comprobaciones...
	if ilTr >= 0 && ilTr < self.SubNet.Length() {
		// Modificacion del valor de la funcion lef
		self.SubNet[ilTr].IiValorLef += aiValLef
		return true
	} else { // Out of range
		return false
	}
}

/*
-----------------------------------------------------------------
   METODO: Shoot
   RECIBE: Indice en el vector de la transicion a Shoot
   DEVUELVE: OK si todo fue bien o ERROR en caso contrario
   PROPOSITO: Shoot una transicion. Esto es, generar todos los eventos
	   ocurridos por el disparo de una transicion
-----------------------------------------------------------------
*/
func (self *LefsDist) Shoot(ilTr IndGlobalTrans) bool {
	// Algunas comprobaciones...
	if ilTr >= 0 && ilTr < self.SubNet.Length() {
		// Prepare 3 local variables
		tiTrans := self.SubNet[ilTr].ITime        // time to spread to new events
		tiDur := self.SubNet[ilTr].IiShotDuration //time length
		listCtes := self.SubNet[ilTr].IiListactes // list of TransCtes

		// La CTE de la primera trans., hace referencia a la cte a mandar a
		// TRANS. QUE SE HA DISPARADO, y va con tiempo igual al de la transicion
		// tiempo, cod_transicion, cte
		self.AddEvents(EventDist{tiTrans, listCtes[0].INextTrans, listCtes[0].Const})

		// Generamos eventos ocurridos por disparo de transicion ilTr
		for _, trCo := range listCtes[1:] {
			// tiempo = tiempo de la transicion + coste disparo
			self.AddEvents(EventDist{tiTrans + tiDur, trCo.INextTrans, trCo.Const})
		}

		return true
	} else {
		return false
	}
}

/*
-----------------------------------------------------------------
   METODO: PrintEventTransitions
   RECIBE: Nada
   DEVUELVE: Nada
   PROPOSITO: Imprimir las transiciones para depurar errores
-----------------------------------------------------------------
*/
func (self LefsDist) PrintEventTransitions() {
	log.Println(" ")
	log.Println("------IMPRIMIMOS LA LISTA DE TRANSICIONES---------")
	for _, tr := range self.SubNet {
		tr.PrintEventValues()
	}
	log.Println("------FINAL DE LA LISTA DE TRANSICIONES---------")
	log.Println(" ")
}

/*
-----------------------------------------------------------------
   METODO: PrintEvent
   RECIBE: Nada
   DEVUELVE: Nada
   PROPOSITO: Imprimir los atributos de la clase para depurar errores
-----------------------------------------------------------------
*/
func (self LefsDist) PrintEvent() {
	u.DistWall()
	log.Println("STRUCT LefsDist")
	log.Println("\tNº transiciones: ", self.SubNet.Length())

	if self.IsTransSensib.isEmpty() {
		log.Println("\tLISTA TRANSICIONES SENSIBILIZADAS VACIA")
	} else {
		log.Println("\tLista transciones sensibilizadas :")
		for _, iTr := range self.IsTransSensib {
			log.Println(iTr, " ")
		}
	}
	log.Println("------Lista transiciones---------")
	for _, tr := range self.SubNet {
		tr.PrintEvent()
	}
	log.Println("------Final lista transiciones---------")

	log.Println("-----------Lista eventos---------")
	self.IlEvents.PrintEvent()
	log.Println("-----------Final lista eventos---------")
	log.Println("FINAL ESTRUCTURA LefsDist")
}
