Accion ej2 (prim1: puntero a nodo1; prim2: puntero a nodo2) es 
	Ambiente 

	nodo1 = registro 
		letra: AN(1)
		prox: puntero a nodo1 
	Freg 
	p1: puntero a nodo1 

	nodo2 = registro 
		letra: AN(1)
		prox,ant: puntero a nodo2 
	Freg 
	p2: puntero a nodo2 

	Procedimiento desencriptrar(K: entero; primero1,primero2,ac1,ac2: puntero a nodo, )
		ac1:= primero1; ac2:= primero2
		Mientras ac2 <> nil hacer 
			Mientras *ac1.letra <> *ac2.letra hacer 
				ac1:= *ac1.prox 
			FM 
			
			si *ac1.letra = *ac2.letra entonces 
				Para i:= 1 hasta K hacer 
					ac1:= *ac1.ant 
				FPara 

				Esc(*ac1.letra)
			Fsi 

			ac2:= *ac2.prox 
		FM 
	FProcedimiento 

	numero: entero 

	Proceso 
		Esc("Ingrese el numero de desplazamiento: ")
		Leer(numero)
		desencriptrar(numero,prim1,prim2,p1,p2) 
	FProceso 
FAccion
