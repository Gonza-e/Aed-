Accion ej2 es 
	Ambiente 

	sec: secuencia de caracteres 
	s: caracter 

	Funcion obtIndice(A: arreglo[1...27] de AN;let: AN;izq,der,medio: entero): entero 
		Mientras izq < der y let <> A[medio] hacer 
			si let < A[medio] entonces 
				der:= medio 
			sino 
				izq:= medio 
			Fsi 

			medio:= (izq+der) div 2
		FM 

		si A[medio] = let entonces 
			obtIndice:= medio 
		sino 
			obtIndice:= -1 
		Fsi 

	FFuncion 

	Funcion insertarLetra(x: entero): AN
		segun x hacer 
			1: insertarLetra:= "A" 
			2: insertarLetra:= "B"
			3: insertarLetra:= "C"
			4: insertarLetra:= "D"
			5: insertarLetra:= "E"
			6: insertarLetra:= "F"
			7: insertarLetra:= "G"
			8: insertarLetra:= "H"
			9: insertarLetra:= "I"
			10: insertarLetra:= "J"
			11: insertarLetra:= "K"
			12: insertarLetra:= "L"
			13: insertarLetra:= "M"
			14: insertarLetra:= "N"
			15: insertarLetra:= "Ñ"
			16: insertarLetra:= "O"
			17: insertarLetra:= "P"
			18: insertarLetra:= "Q"
			19: insertarLetra:= "R"
			20: insertarLetra:= "S"
			21: insertarLetra:= "T"
			22: insertarLetra:= "U"
			23: insertarLetra:= "V"
			24: insertarLetra:= "W"
			25: insertarLetra:= "X" 
			26: insertarLetra:= "Y"
			27: insertarLetra:= "Z"
		Fsegun 
	FFuncion

	A: arreglo[1...27] de AN 

	Procedimiento cargarArreglo()
		Para i:= 1 hasta 27 hacer 
			A[i]:= insertarLetra(i)
		FPara 
	FProcedimiento

	Procedimiento desencriptar()
		indice:= obtIndice(A,*p.og,1,27,14)
		indice:= (indice - desplazamiento) mod 27
		*p.cod:= A[indice]
	FProcedimiento

	nodo = registro 
		og: AN(1)
		cod: AN(1)
		prox: puntero a nodo 
	Freg 
	prim,p,q: puntero a nodo 

	Procedimiento cargaEncolada()
		si prim = nil entonces 
			*q.prox:= prim 
			prim:= q 
		sino 
			p:= prim; ant:= nil 
			Mientras p <> nil hacer 
				ant:= p 
				p:= *p.prox 
			FM 

			*ant.prox:= q 
			*q.prox:= nil 
		Fsi 
	FProcedimiento

	Proceso 
		Arr(sec); Avz(sec,s)
		prim:= nil 

		cargarArreglo()

		Mientras NFDS(sec) hacer 
			Nuevo(q); *q.og:= s 
			si s <> " " entonces 
				desencriptar()
			sino 
				*q.cod:= " "
			Fsi 
			cargaEncolada()

			Avz(sec,s)
		FM 
		Cerrar(sec)

	FProceso 
FAccion 
			
			
