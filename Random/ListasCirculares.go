Accion ListasCirculares (prim: puntero a nodo) es //en caso de que el ejercicio te de la lista la pones aca como parametro
 Ambiente 
	nodo = registro 
		dato: entero 
		prox: puntero a nodo 
	FinReg
	p,q,k,ant: puntero a nodo 
	Proceso 
		Procedimiento insertar_elemento() es 
			Nuevo(q)
			*q.dato:= valor
			Si (prim = nil) entonces  //primero preguntamos si la lista está vacia
				prim:= q 
				*q.prox:= q  //en caso de que este vacia entonces hacemos que prim apunte a "q" y el proximo de "q" apunte a si mismo (es decir que el proximo de "q" apunte a q)
			Sino 
				p:= prim; ant:= nil 
				Mientras (*p.prox<>prim) hacer  //en caso de la lista tenga elementos se va a recorrer 
					ant:= p 
					p:= *p.prox 
				FM 
				Si (p = prim) /*o (ant = nil), esta condicion se aplica tambien cuando (p = prim)*/ entonces //en este caso como (p = prim) es decir, el primer elemento, vamos a hacer que el puntero prox de prim apunte a "q"
					*q.prox:= *prim.prox 
					*prim.prox:= q 
				Sino 
					Si (*p.prox = prim) entonces  //en este caso si (*p.prox = prim) es decir que estamos parados en el nodo que esta antes del nodo al que esta apuntando prim
						*p.prox:= q 
						*q.prox:= prim  //hacemos que el proximo de "p" apunte a "q" y el proximo de "q" apunte a donde apunte prim
					Sino 
						*ant.prox:= q  //el caso anterior es cuando queremos insertar un nodo despues del ultimo elemento, en este caso estamos insertando el elemento en el medio de otros nodos
						*q.prox:= p
					FinSi
				FinSi
			FinSi
		FProcedimiento

	 	Procedimiento eliminar_elemento() es 
			Si prim = nil entonces 
				Esc("Lista vacia, no se puede eliminar nada")
			Sino 
				Esc("Ingrese elemento a eliminar")
				Leer(valor)
				p:= prim
				Mientras (*p.prox<>prim) y (*p.dato<>valor) hacer  //busqueda lineal con centinela
					ant:= p 
					p:= *p.prox 
				FM 
				Si (*p.dato <> valor) entonces 
				    Esc("No se encontro el elemento")
				Sino 
					Si (p = prim) entonces    //esto es lo mismo que en el caso de insercion
						prim:= *prim.prox 
						//disponemos P y luego hacemos que apunte a prim nuevamente ----> Disponer(p); p:= prim 
					Sino 
						Si (*p.prox = prim) entonces  //este seria el caso de eliminar el ultimo elemento 
							*ant.prox:= *p.prox  //al proximo de "ant" tenemos que hacer que apunte al proximo nodo al que este apuntando "p"
							//disponemos P y luego hacemos que apunte a al proximo de "ant" ----> Disponer(p); p:= *ant.prox
						Sino 
							Si (*p.prox = p) entonces  //este seria el caso cuando hay solo un elemento, lo que quiere decir (*p.prox = p) es que el proximo de "p" está apuntando a donde está apuntando "p", es decir, a si mismo
								prim:= nil 
								//Disponer(p)
							FinSi
						FinSi
					FinSi
					e:= p; p:= *p.prox 
					Disponer(e) //esta seria otra forma de disponer un nodo sin la necesidad de borrar "p" y volver a hacer que señale a prim
				FinSi
			FinSi
		FProcedimiento
	FinProceso
FinAccion
		 	 
		 
