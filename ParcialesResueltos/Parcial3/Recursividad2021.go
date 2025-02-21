Accion binario (prim: puntero a nodo) es 
 Ambiente 

 	nodo = registro 
	 dato: entero 
	 prox: puntero a nodo 
	FReg
 p: puntero a nodo

	Funcion Imparidad(numero: entero, unos: entero): logico
	 	Si numero = 0 entonces
		 	Si (unos MOD 3) = 0 entonces
			 Imparidad:= verdadero 
			Sino 
			 Imparidad:= falso
			FSi 
		Sino 
		 	Si (numero MOD 10) = 1 entonces 
			 Imparidad:= Imparidad(numero DIV 10, unos + 1)
			Sino 
			 Imparidad:= Imparidad(numero DIV 10, unos + 1)
			FSi
		FSi 
	FFuncion 

 Proceso 
	 p:= prim 
	 Mientras p <> nil hacer 
	 	 	Si Imparidad(*p.dato,0) = verdadero entonces 
			 accion_verdadero()
			Sino 
			 accion_falso()
			FSi 
		 p:= *p.prox
		FM 
	FProceso 
FAccion