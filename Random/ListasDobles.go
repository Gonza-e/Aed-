Accion ListasDobles es 
 Ambiente 
	nodo = registro 
		dato: entero 
		prox: puntero a nodo 
		ant: puntero a nodo 
	FinReg
	//aclaracion: en listas doblemente enlazadas vamos a tener dos punteros internos en vez de uno
	//aclaracion: "ant" va a apuntar al nodo anterior y "prox" como siempre apunta al proximo
	//aclaracion: vamos a tener un nuevo puntero externo el cual va a ser "ult", el cual marca el ultimo elemento de la lista
	prim,p,q,ult: puntero a nodo 
	Proceso 
		Procedimiento cargar_nodo() es 
			Nuevo(q)
		 	*q.dato:= valor 
		 	Si (prim = nil) entonces
				*q.ant:= nil 
				*q.prox:= nil 
				prim:= q  //hacemos que tanto prim como ult apunten a q (el unico elemento en la lista ya que recien lo ingresamos)
				ult:= q 
			Sino 
				p:= prim 
				Mientras (p<>nil) y (*p.dato > *q.dato) hacer  
					p:= *p.prox 
				FM 
				Si (prim = p) entonces  //aca insertamos "q" antes del primer elemento el cual es "p"
					*q.prox:= p 
					*q.ant:= nil  
					*p.ant:= q 
					prim:= q 
				Sino 
					Si (p = nil) entonces // si p = nil entonces quiere decir que estamos al final de la lista, por lo tanto tenemos que insertar "q" frente al ultimo elemento que es "ult"
						*ult.prox:= q 
						*q.ant:= ult 
						*q.prox:= nil 
						ult:= q  //"ult" tiene que apuntar a "q" para que asi este sea el ultimo elemento
					Sino //en este caso lo que se hace es insertar el elemento "q" en medio de otros nodos
						*(*p.ant).prox:= q //esto se lee como que el puntero "prox" del nodo al que estaba apuntando *p.ant "*(*p.ant)", apunte a "q"
						*q.prox:= p 
						*q.ant:= *p.ant 
						*p.ant:= q
					FinSi
				FinSi
			FinSi
		FProcedimiento

	 	Procedimiento eliminar_nodo() es 
	 	 	Si (prim = nil) entonces 
			 	Esc("Lista vacia")
			Sino 
				Esc("Ingrese valor"); Leer(valor)
				p:= prim 
				Mientras (p<>nil) y (*p.dato<>valor) hacer
					p:= *p.prox
				FM 
				Si (prim = p) entonces  //como estamos en el primer elemento hacemos que prim apunte al proximo nodo 
					prim:= *p.prox 
					*prim.ant:= nil  //como queremos eliminar el primer elemento hacemos que el puntero "ant" de prim apunte a nil
				Sino 
					Si (prim = ult) entonces  //si prim y ult son iguales, es decir, el unico elemento en la lista, solo hacemos que ambos apunten a nil 
						prim:= nil  
						ult:= nil 
						//Disponer(p)
					Sino 
						Si (p = ult) entonces  //si estamos parados en el ultimo elemento hacemos que "ult" apunte al nodo anterior de "p"
							ult:= *p.ant 
							*ult.prox:= nil  //una vez que "ult" apunte al nodo anterior a "p", hacemos que el puntero "prox" de "ult" apunte a nil nada mas
							//Disponer(p)
						Sino 
							*(*p.ant).prox:= *p.prox //en caso de estar en el medio el nodo a eliminar tenemos que hacer que tanto el nodo anterior a "p" y el nodo proximo a "p" se apunten entre si 
							*(*p.prox).ant:= *p.ant 
							//Disponer(p)
						FinSi
					FinSi
				FinSi
			 	Disponer(p)
			FinSi
		FProcedimiento

		Procedimiento cargaApilada() es //este es un invento o intento de carga apilada para listas dobles
			si prim = nil entonces //en caso de estar la lista vacia colocamos el primer nodo
				prim:= q 
				ult:= q 
				*q.ant:= nil 
				*q.prox:= nil 
			sino 
				*q.prox:= prim 
				*q.ant:= nil 
				si *q.prox <> nil entonces //en caso de que el puntero prox apunte a otro nodo hacemos que el puntero "ant" de dicho nodo apunte a "q"
					*(*q.prox).ant:= q 
				fsi 
				prim:= q 
			fsi 
		FProcedimiento
	FinProceso
FinAccion

			