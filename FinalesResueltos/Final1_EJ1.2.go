Accion ej1 es 
 Ambiente 
	nodo = registro 
		num: N(7)
		con: N(4)
		pos: N(4)
		prox: puntero a nodo 
	Freg 
	prim,p,q,nil: puntero a nodo 

	Procedimiento carga()
		si prim = nil entonces 
			*q.prox:= prim 
			prim:= q 
		sino 
			p:= prim; ant:= nil 
			Mientras p <> nil hacer 
				si *p.num = *q.num entonces 
					k:= k + 1 
				Fsi 
				p:= *p.prox 
			FM 

			si k <= 1 entonces 
				*ant.prox:= q 
				*q.prox:= p 
			sino 
				k:= 0
			Fsi 
		Fsi 

	FProcedimiento 

	conjunto,pos,k,numero,min,max,conMin,conMax,contSub,contTotal: entero
	prom: real 
	
	Procedimiento cargarLista()
		si numero = 0 entonces 
			si prom <> 0 y contSub <> 0 entonces 
				prom:= prom/contSub 
			sino 
				prom:= 0 
			Fsi 
			Esc("El promedio del subconjunto ",conjunto," es de:",prom)
			conjunto:= conjunto + 1 
			contTotal:= contTotal + 1
			pos:= 0 
			contSub:= 0 
			prom:= 0
		sino 
			contSub:= contSub + 1 
			prom:= prom + numero 
			pos:= pos + 1
		Fsi 
		Nuevo(q); *q.num:= numero; *q.con:= conjunto; *q.pos:= pos 
		carga()
	FProcedimiento

	Procedimiento buscarMayMen()
		si *p.num > max entonces 
			conMax:= *p.con
			pos:= *p.pos 
			max:= *p.num 
		Fsi 
		si *p.num < min entonces 
			min:= *p.num 
		Fsi 
		p:= *p.prox 
	FProcedimiento

	Procedimiento init()
		prim:= nil
		conjunto:= 1; pos:= 0; k:= 0; min:= HV; max:= 0
		conMin:= 0; conMax:= 0; contSub:= 0; contTotal:= 0
		prom:= 0
	FProcedimiento

	Proceso 
		init()

		repetir 
			Esc("Ingrese un numero: "); Leer(numero)
			si numero > 0 entonces 
				cargarLista()
			Fsi 
		hasta que (numero < 0)

		p:= prim 
		Mientras p <> nil hacer 
			Mientras *p.num <> 0 y p <> nil hacer 
				buscarMayMen()
			FM 
			Esc("Valor minimo del subconjunto: ",min)
			min:= HV 
			p:= *p.prox 
		FM  

		si contSub = 0 y contTotal = 0 entonces 
			Esc("El conjunto está vacio")
		sino 
			Esc("El valor maximo encontrado:",max,"Subconjunto:",conMax,"Posicion relativa:",pos)
			Esc("Subconjuntos totales:",contTotal)
		Fsi 
	FProceso 
FAccion

				