Accion ejercicio2 (a: arreglo[1...60] de enteros) es 
 ambiente
	Procedimiento verificar(a: arreglo[1...60] de enteros, i:entero)
		si i = 61 entonces 
			Esc("Fin del proceso")
		sino 
			si i = a[i] entonces 
				Esc("Coinciden")
			sino 
				Esc("No coinciden")
			fsi 

			verificar:= verificar(a,i+1)
		fsi 
	fprocedimiento 

	Proceso 

		Esc(verificar(a,0))

	FProceso 
FAccion

