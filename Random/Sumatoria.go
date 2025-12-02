// sumatoria hasta N de (3i + 1), con i = 1 

Funcion Sumatoria(n,i: enteros): entero 
	si n = 0 entonces
		Sumatoria:= i 
	sino
		Sumatoria(n - 1, (3*i) + 1) 
	fsi 
FFuncion