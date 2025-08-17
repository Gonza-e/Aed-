Accion ListaDoble es 
 Ambiente 
	nodo = registro 
		letra: AN(1)
		prox, ant: puntero a nodo 
	freg 
	prim,p,q,ult,u,s: puntero a nodo 

	Funcion capicua(proximo,anterior: puntero a nodo): booleano 
		si (proximo = anterior) entonces 
			capicua:= verdadero 
		sino 
			si *proximo.letra = *anterior.letra entonces 
				capicua:= capicua(*proximo.prox,*anterior.ant)
			sino 
				capicua:= falso 
			fsi 
		fsi 
	FFuncion 

	procedimiento insertar() es
		Nuevo(q); *q.letra:= letra 	
		si prim = nil entonces 
			prim:= q
			ult:= q 
			*q.ant:= nil 
			*q.prox:= nil 
		sino 
			*q.prox:= nil 
			*q.ant:= ult 
			*ult.prox:= q 
			ult:= q 
		fsi 
	FProcedimiento

	letra, salir: AN 

	Proceso 
		letra:= ' '; salir:= ' '

		Esc("Presione ENTER entre cada letra, presione 0 y ENTER para terminar")
		repetir 
			Esc("Ingrese una letra: "); Leer(letra)
			insertar()	
		hasta que (Letra = "0")

		s:= prim; u:= ult
		si capicua(verdadero,s,u) entonces 
			Esc("La palabra es capicua")
		sino 
			Esc("La palabra no es capicua")
		fsi 
	FProceso 
FAccion 
		