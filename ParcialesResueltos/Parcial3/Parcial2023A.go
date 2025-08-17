Accion ejercicio1 (prim: puntero a nodo; A: arreglo[1...12,1...3] de entero) es
 Ambiente 
	Fecha = registro 
		dia: N(2)
		mes: N(2)
		anio: N(4)
	FinReg
	nodo2 = registro 
		fecha: Fecha
		cant: N(3)
		cod: AN(20)
		dni: N(8)
		prox: puntero a nodo2
		ant: puntero a nodo2
	FinReg
	prim9,prim10,prim11,q,s: puntero a nodo2
	nodo = registro 
		dni: N(8)
		n_fila: N(3)
		f_fila: Fecha
		f_recital: Fecha
		prox: puntero a nodo
	FinReg
	p,e,ant: puntero a nodo
	Procedimiento iniciar_segunda_lista() es 
		Nuevo(prim9); *prim9.fecha.dia:= 9; *prim9.fecha.mes:= 11; *prim9.cant:= 0
		*prim9.ant:= nil; *prim9.prox:= nil 
		Nuevo(prim10); *prim10.fecha.dia:= 10; *prim10.fecha.mes:= 11; *prim10.cant:= 0 
		*prim10.ant:= prim9; *prim10.prox:= nil; *prim9.prox:= prim10
		Nuevo(prim11); *prim11.fecha.dia:= 9; *prim11.fecha.mes:= 11; *prim11.cant:= 0
		*prim11.ant:= prim10; *prim11.prox:= nil; *prim10.prox:= prim11
	FP
	Procedimiento borrar_swiftie() es 
	 	Si (p = nil) entonces 
			Esc("La lista esta vacia")
		Sino 
			Si (prim = p) entonces 
			 	prim:= nil 
			Sino 
			 	*ant.prox:= *p.prox
			FinSi
		FinSi
		e:= p 
		p:= *p.prox
		disponer(e)
	FP
	Procedimiento cargar_swiftie() es 
	 	*q.prox:= *s.prox 
		Si *s.prox <> nil 
		 	*(*s.prox).ant:= q 
		FSi 
		*q.ant:= s 
		*s.prox:= q 
	FP
	cod: AN
	mayor, f_mayor: entero 
	Proceso 
		p:= prim 
		iniciar_segunda_lista()
		Mientras (p<>nil) hacer 
			Segun *p.f_recital.dia hacer 
				9: s:= prim9
				10: s:= prim10
				11:= s:= prim11
			FSegun
			Nuevo(q)
			*q.cod:= swiftieEncriptada(*p.n_fila,A[*p.f_fila.mes,*p.f_recital.dia])
			*q.dni:= *p.dni 
			cargar_swiftie()
			borrar_swiftie()
		FM
		s:= prim9
		mayor:= LV
		Mientras (s<>nil) hacer
	 		Si *s.fecha <> " " entonces  
	 	 		Si *s.cant > mayor entonces 
					mayor:= *s.cant 
					f_mayor:= *s.fecha.dia
				FinSi
			FinSi
		 	s:= *s.prox
		FM
	 	Esc("El dia con mas swifties es el",f_mayor,"con",mayor,"swifties")
	FProceso 
FAccion


	

 

	 

	 



	 	 
			 
	
	 
