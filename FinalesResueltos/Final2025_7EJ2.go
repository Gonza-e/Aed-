Accion ejercicio2 es 
 ambiente 
	sec: secuencia de caracteres 
	v: caracter 

	nodo = registro 
		letra: AN(1)
		prox: puntero a nodo 
	Freg 
	prim,p,q,s: puntero a nodo 

	Procedimiento cargaEncolada() 
		p:= prim 
		Mientras *p.prox <> prim hacer 
			p:= *p.prox 
		FM 

		*q.prox:= *p.prox 
		*p.prox:= q 
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
			cargaEncolada()
		fpara 
	FProcedimiento

	Procedimiento desencriptar()
		p:= prim 
		Mientras (*p.prox <> prim) y (*p.letra <> v) hacer
			k:= k - 1 
			p:= *p.prox 
		FM 

		si *p.letra = v entonces 
			s:= p 
			

		


