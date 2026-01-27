Funcion buscar(prim: puntero a nodo, dato: entero):entero 
	si prim.dato = dato entonces 
		buscar:= dato 
	fsi 

	si prim = nil entonces 
		buscar:= nil 
	fsi 
	 
	si *prim.dato < dato entonces 
		buscar:= buscar(*prim.der,dato)
	sino 
		si *prim.dato > dato 
			buscar:= buscar(*prim.izq,dato)
		fsi 
	fsi 

FFuncion 

Procedimiento insertar(prim, q: puntero a nodo)
	si prim = nil entonces 
		*q.der:= nil 
		*q.izq:= nil 
		prim:= q 
	sino 
		insertarRecursivo(prim,q)
	fsi 
FProcedimiento 

Procedimiento insertarRecursivo(prim, q: puntero a nodo)
	si *q.dato <= *prim.dato entonces 
		si *prim.izq = nil entonces 
			*prim.izq:= q 
			*q.izq:= nil 
			*q.der:= nil 
		sino 
			insertarRecursivo(*prim.izq,q)
		fsi 
	sino 
		si *q.dato > *prim.dato entonces 
			si *prim.der = nil entonces
				*prim.der:= q 
				*q.izq:= nil 
				*q.der:= nil 
			sino 
				insertarRecursivo(*prim.der,q)
			fsi 
		fsi 
	fsi 
FProcedimiento

Funcion calcNivel(prim: puntero a nodo): entero 
	si prim = nil entonces 
		calcNivel:= 0 
	sino 
		si calcNivel(*prim.izq) > calcNivel(*prim.der) entonces 
			calcNivel:= 1 + calcNivel(*prim.izq)
		sino 
			calcNivel:= 1 + calcNivel(*prim.der)
		Fsi 
	Fsi 
FFuncion

Procedimiento enOrden(prim: puntero a nodo)
	si prim <> nil entonces 
		enOrden(*prim.izq)
		tratarRaiz()
		enOrden(*prim.der)
	fsi 
FProcedimiento

Procedimiento preOrden(prim: puntero a nodo)
	si prim <> nil entonces 
		tratarRaiz()
		preOrden(*prim.izq)
		preOrden(*prim.der)
	fsi 
FProcedimiento

Procedimiento postOrden(prim: puntero a nodo)
	si prim <> nil entonces 
		postOrden(*prim.izq)
		postOrden(*prim.der)
		tratarRaiz()
	fsi 
FProcedimiento