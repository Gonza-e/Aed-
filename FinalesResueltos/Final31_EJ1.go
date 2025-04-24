Accion final31 es 
 Ambiente 
	sec: secuencia de caracteres 
	v: caracter 
	premium, esigual: booleano
	edad, cant_nulo, cant_prmn, digito_UNO, digito_DOS, total: entero 
	prom, porc_prmn: real 

	nodo = registro 
		codpack: arreglo[1...30] de AN 
		pref: arreglo[1...2] de enteros 
		prox: puntero a nodo 
	freg 
	prim,p,q,ant: puntero a nodo 

	aux: arreglo[1...30] de AN 
	
	funcion verificarCOD(arr1,arr2: arreglo[1...30] de AN): booleano es
		Mientras ((arr1[i] <> nil) y (arr2[i] <> nil)) y (arr1[i] = arr2[i]) hacer
			i:= i + 1 
		FM 

		si arr1[i] <> arr2[i] entonces 
			verificarCOD:= falso
		sino 
			verificarCOD:= verdadero 
		fsi 
	FFuncion

	procedimiento crearNodo() es 
		nuevo(q); *q.codpack:= aux; sumarPref(q)
	FP

	procedimiento sumarPref(nodo: puntero a nodo) es 
		segun v hacer 
			"A": *nodo.pref[1]:= *nodo.pref[1] + 1
			"D": *nodo.pref[2]:= *nodo.pref[2] + 1
		fsegun
	FP

 	Proceso
	 	Arr(sec); Avz(sec,v)
		prim:= nil 
		premium:= falso; esigual:= falso
		edad:= 0, cant_nulo:= 0, total:= 0; cant_prmn:= 0; prom:= 0; porc_prmn:= 0

		Mientras NFDS(sec) hacer 
			Mientras v <> "/" hacer 
				si v = "P" entonces 
					premium:= verdadero 
				fsi 
				Avz(sec,v)

				para i:= 1 a 6 hacer 
					Avz(sec,v)
				fpara 

				digito_UNO:= v
				Avz(sec,v)
				digito_DOS:= v 

				i:= 1
				Mientras v <> "-" hacer 
					aux[i]:= v 
					i:= i + 1 
					Avz(sec,v)
				FM 
				Avz(sec,v)

				si v <> "N" entonces 
					p:= prim 
					si prim = nil entonces
						crearNodo()
						*q.prox:= prim 
						prim:= q 
					sino 
						Mientras p <> nil hacer
							si verificarCOD(*p.codpack[i], aux[i]) entonces 
								sumarPref(p)
								esigual:= verdadero
							fsi 

							p:= *p.prox 
						FM 

						si (p = nil) y (esigual = falso) entonces 
							crearNodo()
							*q.prox:= prim 
							prim:= q
						sino 
							esigual:= falso 
						fsi 
					fsi 

					si premium entonces 
						cant_prmn:= cant_prmn + 1 
						premium:= falso 
					fsi 
				sino 
					cant_nulo:= cant_nulo + 1 
					edad:= edad + AGE(digito_UNO, digito_DOS)
				fsi 
			FM 
			total:= total + 1 
		FM 
		cerrar(sec)

		prom:= edad/cant_nulo
		porc_prmn:= (cant_prmn/total)*100

		p:= prim 
		Mientras p <> nil hacer  // Consigna 1 y 3
			Esc("Codigo: ")
			para i:= 1 a 30 hacer 
				Esc(*p.codpack)
			fpara 

			Esc("Preferencias A: ",*p.pref[1])
			Esc("Preferencias D: ",*p.pref[2])

			si *p.pref[1] > *p.pref[2] entonces 
				Esc("El pack cumple con la condicion de A mayor que D")
			fsi 
		FM 

		Esc("Porcentaje de clientes premium con preferencias: ",porc_prmn:0:2,"%") // Consigna 4
		Esc("Promedio de edad de clientes con preferencias nulas: ",prom:0:2) // Consigna 2
	FProceso  
FAccion





