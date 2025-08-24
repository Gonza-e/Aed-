accion esPrimo es 
 Ambiente 
	Funcion esPrimo(num1,num2): booleano
		si num1 = 1 entonces 
			esPrimo:= falso 
		sino  
			si num2 = 0 entonces
				esPrimo:= verdadero 
			sino 
				si (num1 mod num2 = 0) o (num1 = 1) entonces 
					esPrimo:= falso 
				sino 
					esPrimo:= esPrimo(num1,num2-1)
				fsi 
			fsi 
		fsi
	FFuncion