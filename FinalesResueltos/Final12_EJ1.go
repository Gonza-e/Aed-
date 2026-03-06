Accion ej1 es 
	Ambiente

	sec: secuencia de caracter 
	s: caracter 

	nodoP = registro 
		prin: puntero a nodo
		compras: N(4)
		nu,ma,de: N(4)
		prox: puntero a nodoP
	Freg 
	primP,pP,qP: puntero a nodoP 

	nodo = registro 
		l: AN(1)
		prox: puntero a nodo
	Freg 
	prim,p,q,ult: puntero a nodo

	edad: entero 
	prom: real 

	Procedimiento cargarNodo()
		Nuevo(q); *q.l:= s 
		si prim = nil entonces 
			*q.prox:= prim 
			prim:= q 
			ult:= q 
		sino 
			*ult.prox:= q 
			*q.prox:= nil 
			ult:= q 
		Fsi 
	FProcedimiento

	Procedimiento cargarNodoP()
		si primP = nil entonces 
			crearNodoP()
			*qP.prox:= primP
			primP:= qP 
		sino 
			pP:= primP 
			Mientras pP<>nil y NO iguales(*pP.prin,prim) hacer 
				pP:= *pP.prox 
			FM 

			si iguales(*pP.prin,prim) entonces 
				actualizar()
				borrar()
			sino 
				crearNodoP()
				*qP.prox:= primP
				primP:= qP 
			Fsi 
		Fsi
	Fsi 

	Funcion iguales(a,n: puntero a nodo): booleano
		si a = nil y n = nil entonces 
			iguales:= verdadero
		sino 
			si *a.l = *n.l entonces 
				iguales:= iguales(*a.prox,*n.prox)
			sino 
				iguales:= falso 
			Fsi 
		Fsi 
	FFuncion 

	Procedimiento verificar(nodo: puntero a nodoP)
		segun pref hacer 
			"M": *nodo.ma:= *nodo.ma+1
			"N": *nodo.nu:= *nodo.nu+1
			"D": *nodo.de:= *nodo.de+1
		Fsegun 
	FProcedimiento

	Procedimiento crearNodoP()
		Nuevo(qP); *qP.prin:= prim; *qP.nu:= 0; *q.de:= 0; *q.ma:= 0; *qP.compras:= 1
		verificar(qP)
	FProcedimiento

	Procedimiento actualizar()
		*pP.compras:= *pP.compras+1
		verificar(pP)
	FProcedimiento

	Procedimiento borrar()
		p:= prim 
		Mientras p<>nil hacer 
			prim:= *prim.prox 
			Disponer(p)
			p:= prim 
		FM 
	FProcedimiento

	Proceso 
		Arr(sec); Avz(sec,s)
		edad:= 0; prom:= 0

		Mientras NFDS(sec) hacer 
			Para i:= 1 hasta 7 hacer 
				Avz(sec,s)
			FPara 

			Para i:= 1 hasta 2 hacer 
				edad:= edad*10 + conv(s)
				Avz(sec,s)
			FPara 

			Mientras s <> "-" hacer 
				cargarNodo()
				Avz(sec,s)
			FM 
			Avz(sec,s)
			cargarNodoP()

			pref:= s 
			si s = "N" entonces 
				cant:= cant+1
				prom:= prom+edad 
			Fsi 
			Avz(sec,s)
		FM 
		Cerrar(sec)

		prom:= prom/cant 

		Esc("El promedio de edad es de: ",prom)

	FProceso 
FAccion 
