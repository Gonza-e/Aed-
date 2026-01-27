// sumatoria hasta N de (3i + 1), con i = 1 

Funcion Sumatoria(n,i,suma: enteros): entero 
	si n = 0 entonces
		Sumatoria:= suma
	sino
		Sumatoria(n-1,i+1,suma+((3*i)+1)) 
	fsi 
FFuncion