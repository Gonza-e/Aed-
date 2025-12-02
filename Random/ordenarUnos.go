procedimiento ordenar(num: entero)
	si num > 0 entonces 
		si (num mod 10) = 1 entonces 
			Esc(num mod 10)
		fsi 

		ordenar(num div 10)

		si (num mod 10) = 0 entonces 
			Esc(num mod 10)
		fsi 
	fsi 
FP

