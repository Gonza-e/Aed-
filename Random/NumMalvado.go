accion Malvado es 
 Ambiente 
	Funcion malvado(num,cont: entero): booleano
		si num = 0 entonces 
			si cont mod 2 = 0 entonces 
				malvado:= verdadero
			sino 
				malvado:= falso
			fsi 
		sino 
			si num mod 10 = 1 entonces 
				malvado:= malvado(num div 10,cont+1)
			sino 
				malvado:= malvado(num div 10,cont)
			fsi 
		fsi 
	FFuncion