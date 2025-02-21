Accion ejercicio1 (prim,ult: puntero a nodo) es 
 Ambiente 
 	Fecha = registro
	 	dia: N(2)
	 	mes: N(2)
	 	anio: N(4)
	FReg

 	nodo = registro 
	 	fecha: Fecha
	 	cant: N(5)
	 	cod: AN(20)
	 	dni: N(8)
	 	prox,ant: puntero a nodo
	FReg
 	p,k,e: puntero a nodo 

 	nodo2 = registro
	 	cod: AN(20)
		dni: N(8)
	 	prox: puntero a nodo
	FReg
 	prim2,q,p2,ant2: puntero a nodo2

	Procedimiento eliminar() es 
		Si (p = ult) entonces 
			ult:= *p.ant 
			*ult.prox:= nil //(o *p.prox)
		Sino 
			*(*p.ant).prox:= *p.prox 
			*(*p.prox).ant:= *p.ant
		FinSi
	 e:= p 
	 p:= *p.ant
	 disponer(e)
	FP

 	Procedimiento cargar() es 
	 	Si (prim = nil) entonces
		 	*q.prox:= nil 
		 	prim:= q 
		Sino 
		 	p2:= prim2
		 	Mientras (p2<>nil) y (*q.cod > *p2.cod) hacer 
		 	 	ant2:= p2
			 	p2:= *p2.prox 
			FM 
			Si (prim2 = p2) entonces
			 	*q.prox:= p2 
			 	prim2:= q 
			Sino 
			 	ant2.prox:= q 
			 	*q.prox:= p2 
			FinSi
		FinSi
	FP
 	mes,dia: entero
 	Proceso 
	 	k:= prim
		p:= *k.prox
	 	prim2:= nil 
	 	Mientras (p<>nil) hacer 
		    Si (*p.fecha = " ") entonces 
				Si swiftieHabilitada(*p.cod,*p.dni) entonces
				 	eliminar()
					*k.cant:= *k.cant - 1
				Sino 
				 	Nuevo(q)
				 	*q.cod:= funcionDesencriptarLugarEnLaFila(*p.cod)
				 	*q.dni:= *p.dni
				 	cargar()
				FinSi
				p:= *p.prox
			Sino 	
			 k:= p 
			FinSi
		FM 
		
		Esc("Ingrese mes y dia del show"); Leer(dia); Leer(mes)
	 	p:= prim
		Mientras (p<>nil) hacer 
	 	 	Si (*p.fecha.mes = mes) y (*p.fecha.dia = dia) entonces
			 	Esc(*p.cant) 
			Sino 
			 	p:= *p.prox
			FinSi
		FM
	FProceso
FAccion

Accion ejercicio2 (prim: puntero a nodo) es 
 Ambiente
 	nodo = registro 
	 	dato: entero 
	 	prox: puntero a nodo 
	FReg
 	p: puntero a nodo

 	nodo2 = registro 
	 	dato: entero 
	 	prox,ant: puntero a nodo2
	FReg
 	prim2,q: puntero a nodo2

 	Funcion multiplo3(x: entero, cant: entero): logico 
	 	Si (x = 0) entonces
			Si (cant MOD 3 = 0) entonces 
			 	multiplo3:= verdadero 
			Sino 
			 	multiplo3:= falso 
			FinSi
		Sino 
			Si (x MOD 10 = 0) entonces 
			 	multiplo3:= multiplo3(x DIV 10,cant + 1)
			Sino 
			 	multiplo3:= multiplo3(x DIV 10,cant)
			FinSi
		FinSi
	FFuncion
 Proceso 
	 	p:= prim 
	 	prim2:= nil 
	 	Mientras (*p.prox<>prim) hacer 
	 	 	Si multiplo3(*p.dato,0) entonces 
			 	Esc("Los ceros son multiplos de 3")
			Sino 
			 	Nuevo(q); *q.dato:= *p.dato
			 	Si (prim2 = nil) entonces 
				 	*q.prox:= q 
				 	prim2:= q 
				Sino 
				 	*q.prox:= prim2 
				 	prim2:= q 
				FinSi
			FinSi
		FM 
	FProceso
FAccion
 
	
	 
	 
	 
			


		 
		 

	 	 
