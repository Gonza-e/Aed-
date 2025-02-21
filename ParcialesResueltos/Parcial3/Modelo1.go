Accion ejercicio1 (prim: puntero a nodo) es 
 Ambiente 
 	nodo = registro 
	 legajo: N(5)
	 apeynom: AN(30)
	 comision: AN(3)
	 notas: arreglo[1...5] de enteros 
	 prox, ant: puntero a nodo 
	FReg
 p,e: puntero a nodo 
 	nodo2 = registro 
	 cursoNro: N(1)
	 cant: N(3)
	 legajo: N(5)
	FReg
 prim1, prim2, prim3, prim4, prim5, prim6, q, c: puntero a nodo2
 
 	Procedimiento cargar_datos(x: puntero a nodo2, n: entero) es 
	 *x.cursoNro:= n
	 *x.cant:= 0 
	 *x.legajo:= 0 
	FProcedimiento

 	Procedimiento iniciar_lista2() es 
	 Nuevo(prim1); cargar_datos(prim1,1)
	 Nuevo(prim2); cargar_datos(prim2,2) 
	 Nuevo(prim3); cargar_datos(prim3,3) 
	 Nuevo(prim4); cargar_datos(prim4,4)
	 Nuevo(prim5); cargar_datos(prim5,5)
	 Nuevo(prim6); cargar_datos(prim6,6)
	 *prim1.ant:= nil; *prim1.prox:= prim2
	 *prim2.ant:= prim1; *prim2.prox:= prim3
	 *prim3.ant:= prim2; *prim3.prox:= prim4
	 *prim4.ant:= prim3; *prim4.prox:= prim5
	 *prim5.ant:= prim4; *prim5.prox:= prim6
	 *prim6.ant:= prim5; *prim6.prox:= nil
	FProcedimiento 

 	Procedimiento cargar_lista(primero: puntero a nodo, nuevo: puntero a nodo)
	 *nuevo.prox:= *primero.prox 
	 	Si *primero.prox <> nil entonces
		 *(*primero.prox).ant:= nuevo
		FSi 
	 *nuevo.ant:= primero 
	 *primero.prox:= nuevo
	 *primero.cant:= *primero.cant + 1
	FProcedimiento 

	Procedimiento eliminar() es 
	 	Si p = prim entonces 
		 *prim.prox:= *p.prox 
		Sino 
		 *ant.prox:= *p.prox 
		FSi 
	 e:= p 
	 p:= *p.prox 
	 Disponer(e)
	FProcedimiento

	Procedimiento promedio(x1, x2, x3, x4, x5, x6: puntero a nodo2)
	 prom:= (*x1.cant + *x2.cant + *x3.cant + *x4.cant + *x5.cant + *x6.cant)/6 
	FProcedimiento
 prom, mayor, aula_mayor: entero 

 Proceso 
	 prom:= 0 
	 aula_mayor:= 0 
	 mayor:= 0
	 iniciar_lista2()
	 p:= prim 
	 Mientras (p <> nil) hacer 
	 	  	Si EsAlumnoLibre(*p.legajo) entonces 
			 p:= *p.prox 
			Sino 
			 Nuevo(q); *q.cursoNro:= AsignarAula(); *q.legajo:= *p.legajo; *q.cant:= 1
			 	Segun *q.cursoNro hacer 
				 = 1: cargar_lista(prim1,q) 
				 = 2: cargar_lista(prim2,q) 
				 = 3: cargar_lista(prim3,q) 
				 = 4: cargar_lista(prim4,q) 
				 = 5: cargar_lista(prim5,q) 
				 = 6: cargar_lista(prim6,q) 
				FSegun 
			 eliminar()
			FSi 
		FM 
	 promedio(prim1,prim2,prim3,prim4,prim5,prim6)
	
	FinProceso 
FinAccion


Accion Ejercicio2 (prim: puntero a nodo) es 
 Ambiente 
 	nodo = registro 
	 legajo: N(5)
	 apeynom: AN(30)
	 comision: AN(3)
	 notas: arreglo[1...5] de enteros 
	 prox, ant: puntero a nodo 
	FReg
 p, p2, ant, ant2, prim2, q: puntero a nodo 

 	Procedimiento cargar_lista() es 
	 	Si prim2 = nil entonces
		 *q.prox:= prim2 
		 prim2:= q 
		Sino 
		 p2:= prim2; ant2:= nil 
		 Mientras (p2 <> nil) Y (*q.legajo < *p.legajo) hacer 
		 	 ant2:= p2 
			 p2:= *p2.prox 
			FM 
			Si prim2 = p2 ent 
			 *q.prox:= *p2.prox 
			 *p2.prox:= q 
			Sino 
			 ant2.prox:= q 
			 *q.prox:= *p2.prox 
			FSi 
		FSi
	FProcedimiento

	Calcular(cont: entero, i: entero, A: arreglo[1...5] de enteros): logico
		Si i = 0 entonces 
			Si cont >= 3 entonces 
			 Calcular:= verdadero 
			Sino 
			 Calcular:= falso
			FSi 
		Sino 	
		 	Si A[i] >= 6 entonces 
			 Calcular:= Calcular(cont + 1, i - 1 )
			Sino 
			 Calcular:= Calcular(cont, i - 1)
			FSi 
		FSi
	 cont:= 0
	FFuncion
 Proceso 
	 prim2:= nil
	 p:= prim, ant:= nil 
	 Mientras (p <> nil) hacer 
	 		Si Calcular(0,5,*p.notas) entonces
			 Nuevo(q); *q.legajo:= *p.legajo; *q.apeynom:= *p.apeynom; *q.comision:= *p.comision
	 		 cargar_lista()
			FSi
		FM 
	FinProceso
FinAccion 
		       