Accion ej3 es 
 Ambiente 
	nodo = registro 
		letra: AN(1)
		prox,ant: puntero a nodo 
	Freg 
	prim,ult,p,q: puntero a nodo 

	Procedimiento cargar()
		si prim = nil entonces 
			prim:= q 
			ult:= q 
			*q.ant:= nil 
			*q.prox:= nil 
		sino 
			*q.prox:= *ult.prox 
			*q.ant:= ult 
			*ult.prox:= q 
			ult:= q 
		Fsi 
	FProcedimiento

	Funcion palindromo(p,a: puntero a nodo): booleano 
		si (p = a) o (*p.ant = a) entonces 
			palindromo:= verdadero 
		sino
			si *p.letra = *a.letra entonces 
				palindromo:= palindromo(*p.prox,*a.ant)
			sino 
				palindromo:= falso 
			Fsi 
		Fsi 
	FFuncion

	letra: AN 
	cont: entero 

	Proceso 
		prim:= nil 
		cont:= 16 

		Esc("Ingrese su frase letra por letra")
		Mientras cont > 0 hacer 
			Leer(letra)
			si letra <> " " entonces 
				Nuevo(q); *q.letra:= letra 
				cargar()
				cont:= cont - 1 
			Fsi 
		FM 

		si palindromo(prim,ult) entonces 
			Esc("Es palindromo")
		sino 
			Esc("No es palindromo")
		Fsi 
	FProceso 
FAccion
	