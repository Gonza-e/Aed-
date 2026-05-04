Accion ej2 (Alf: arreglo["A"..."Z"] de entero) es 
	Ambiente
	
		sec: secuencia de caracteres 
		s: caracter 
		sal: secuencia de enteros

		bandera: booleano
		vocal = ("A","E","I","O","U")

		Funcion Enc(c: AN; bandera: booleano): entero 
			si bandera entonces 
				Enc:= Alf[c] + raizCuadrada(Alf[c])
			sino 
				Enc:= Alf[c] + factorial(Alf[c])
			Fsi 
		FFuncion 

		nodo = registro 
			letra: AN(1)
			prox: puntero a nodo 
		Freg 
		prim,p,q,ult: puntero a nodo 
		
		Procedimiento limpiarLista()
			p:= prim
			Mientras p <> nil hacer 
				prim:= *p.prox 
				Disponer(p)
				p:= prim 
			FM 
		FP 

		Procedimiento cargarEncolada()
			Nuevo(q); *q.letra:= s 

			si prim = nil entonces 
				*q.prox:= prim 
				prim:= q
				ult:= q 
			sino 
				*q.prox:= *ult.prox 
				*ult.prox:= q 
				ult:= q 
			Fsi 
		FP 

	Proceso 
		Arr(sec); Avz(sec,s); Crear(sal)

		prim:= nil 
		Mientras NFDS(sec) hacer 
			Mientras s <> "" hacer 
				cargarEncolada()
				Avz(sec,s)
			FM 
			Avz(sec,s)

			p:= prim; bandera:= falso 
			si *ult.letra EN vocal entonces 
				bandera:= verdadero 
			Fsi 

			Mientras p <> nil hacer 
				Grabar(sal,Enc(*p.letra,bandera)) 
				p:= *p.prox 
			FM 
			limpiarLista()
			Grabar(sal,0)
		FM 
		Cerrar(sec); Cerrar(sal)

	FProceso 
FAccion 
