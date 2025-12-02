Accion digitoPrimo es 
 Ambiente 
	funcion esPrimo(num1,num2: entero): entero 
		si num1 < 2 entonces 
			esPrimo:= Falso 
		sino 
			si num2 = 0 entonces 
				esPrimo:= Verdadero 
			sino 
				si num1 mod num2 = 0 entonces 
					esPrimo:= Falso
				sino 
					esPrimo:= esPrimo(num1,num2-1)
				fsi 
			fsi 
		fsi 
	ffuncion 

	k, iteraciones, numero, digito: entero 

	nodo_circ = registro 
		numero: N(12)
		prox: puntero a nodo_circ
	freg
	p_circ: puntero a nodo_circ 

	Proceso 
		p_circ:= prim_circ
		
		Esc("Ingrese un numero"); Leer(iteraciones)

		Mientras *p_circ.prox <> prim_circ hacer 
			k:= iteraciones 
			numero:= *p_circ.numero
			Mientras k > 1 hacer 
				k:= k - 1 
				numero:= numero div 10 
			FM 
			digito:= numero mod 10 

			si esPrimo(digito, digito - 1) entonces 
				Esc("El numero es primo")
			sino 
				Esc("El numero no es primo")
			fsi 
		FM

		si *p_circ.numero = prim_circ entonces 
			Mientras iteraciones > 1 hacer 
				iteraciones:= iteraciones - 1 
				numero:= numero div 10 
			FM 
			digito:= numero mod 10 

			si esPrimo(digito, digito - 1) entonces 
				Esc("El numero es primo")
			sino 
				Esc("El numero no es primo")
			fsi 
		fsi 
	
	FProceso 
FAccion
