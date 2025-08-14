accion ejercicio2 es 
 Ambiente 
	Funcion Guardian(A,B,sumA,sumB,resA,resB): booleano
		si resA = 0 y resB = 0 entonces 
			si sumA = A y sumB = B entonces 
				Guardian:= verdadero 
			sino 
				Guardian:= falso 
			fsi 
		sino 
			si resA > 0 entonces // esto en caso de que un numero llegue a 0 primero que el otro
				si (resA mod 2 <> 0) y (A mod resA = 0) entonces 
					Guardian:= Guardian(A,B,sumA+resA,sumB,resA-1,resB)
				sino 
					Guardian:= Guardian(A,B,sumA,sumB,resA-1,resB)
				fsi 
			fsi 

			si resB > 0 entonces 
				si (resB mod 2 = 0) y (B mod resB = 0) entonces
					Guardian:= Guardian(A,B,sumA,sumB+resB,resA,resB-1)
				sino 
					Guardian:= Guardian(A,B,sumA,sumB,resA,resB-1)
				fsi 
			fsi
		fsi 
	FFuncion 

	numA,numB: enteros 
	Proceso 
		Esc("Ingrese extremo inferior: "); Leer(numA)
		Esc("Ingrese extremo superior: "); Leer(numB)

		Para i:= numA hasta numB - 1 hacer 
			Para j:= i + 1 hasta numB hacer 
				si Guardian(numA,numB,0,0,numA-1,numB-1) entonces 
					Esc("Los numeros: ",numA," y ",numB," son guardianes")
				sino 
					Esc("Los numeros no son guardianes")
				fsi 
			FPara
		FPara 

	FProceso 
FAccion



