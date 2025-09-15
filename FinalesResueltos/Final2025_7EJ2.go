Accion ejercicio2 es 
 ambiente 
	sec: secuencia de caracteres 
	v: caracter 

	nodo = registro 
		letra: AN(1)
		ant, prox: puntero a nodo 
	Freg 
	prim,p,q,s: puntero a nodo 

	nodo2 = registro
		og, des: AN(1)
		prox: puntero a nodo2 
	Freg
	prim2,p2,q2,ant2: puntero a nodo2 

	Procedimiento cargaCircularDoble() 
		*q.prox:= *ult.prox 
		*(*ult.prox).ant:= q 
		*q.ant:= ult 
		*ult.prox:= q 
		ult:= q 
	FProcedimiento 

	Funcion num(n:entero): AN 
		1: num:= "A"
		2: num:= "B"
		3: num:= "C"
		4: num:= "D"
		5: num:= "E"
		6: num:= "F"
		7: num:= "G"
		8: num:= "H"
		9: num:= "I"
		10: num:= "J"
		11: num:= "K"
		12: num:= "L"
		13: num:= "M"
		14: num:= "N"
		15: num:= "O"
		16: num:= "P"
		17: num:= "Q"
		18: num:= "R"
		19: num:= "S"
		20: num:= "T"
		21: num:= "U"
		22: num:= "V"
		23: num:= "W"
		24: num:= "X"
		25: num:= "Y"
		26: num:= "Z"
	FFuncion
		
	Procedimiento cargarLetras()
		para i:= 1 hasta 26 hacer 
			Nuevo(q); *q.letra:= num(i)
			si prim = nil entonces 
				*q.ant:= q 
				*q.prox:= q 
				prim:= q 
				ult:= q 
			sino 
				cargaCircularDoble()
			fsi 		
		fpara 
	FProcedimiento

	Procedimiento cargaEncolada()
		p2:= prim2
		Mientras p2 <> nil hacer 
			ant2:= p2 
			p2:= *p2.prox 
		FM 

		*q2.prox:= *ant2.prox 
		*ant2.prox:= q2 
	FProcedimiento

	Procedimiento desencriptar()
		p:= prim 
		Mientras (*p.letra <> v) hacer
			p:= *p.prox 
		FM 

		si *p.letra = v entonces 
			s:= p 
			para i:= 1 hasta k hacer 
				s:= *s.ant 
			fpara 

			Nuevo(q2); *q2.og:= *p.letra; *q2.des:= *s.letra 
			cargarEncolada()
		fsi 
	FProcedimiento

	Proceso 
		ARR(sec); AVZ(sec,v)
		prim:= nil; prim2:= nil 
		cargarLetras()

		Mientras NFDS(sec) hacer 
			desencriptar()
			AVZ(sec,v)
		FM 
		Cerrar(sec)
	FProceso 
FAccion



			

		


