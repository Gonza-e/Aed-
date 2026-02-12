Accion listdelist es 
	nodo1 = registro 
		valor: N(4)
		prox: puntero a nodo1 
	Freg 
	prim,p,q: puntero a nodo 

	nodo2 = registro 
		prim: puntero a nodo1
		prox: puntero a nodo2 
	Freg 
	prim2,p2,q2: puntero a nodo2

	Procedimiento cargarLista1(num)
		Nuevo(q); *q.valor:= num
		*q.prox:= prim 
		prim:= q 
	FProcedimiento 

	Procedimiento cargarLista2()
		Nuevo(q2); *q2.prim:= prim 
		*q2.prox:= prim2 
		prim2:= q2
		prim:= nil  
	FProcedimiento

	Proceso 
		prim:= nil; prim2:= nil 

		Para i:= 1 hasta 9 hacer 
			cargarLista1(i)
		FPara 

		cargarLista2()

		Para i:= 1 hasta 4 hacer 
			cargarLista1(i)
		FPara 

		cargarLista2()

		p2:= prim2 
		Mientras p2 <> nil hacer 
			p:= *p2.prim 
			Mientras p <> nil hacer 
				Esc("Valor: ",*p.valor)
				p:= *p.prox 
			FM 
			p2:= *p2.prox 
		FM 

	FProceso 
FAccion