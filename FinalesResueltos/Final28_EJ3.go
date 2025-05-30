Accion final28 (prim: puntero a nodo) es 

	Arbol = registro 
		dato: entero 
		izq: puntero a nodo 
		der: puntero a nodo 
	freg 
	prim: puntero a nodo 

	Funcion buscar(prim: puntero a nodo, dato: entero):entero 
		si prim.dato = dato entonces 
			buscar:= dato 
		sino 
			si *prim.dato < dato entonces 
				buscar:= buscar(*prim.der,dato)
			sino 
				si *prim.dato > dato 
					buscar:= buscar(*prim.izq,dato)
				fsi 
			fsi 
		fsi 
	FFuncion 
FAccion
