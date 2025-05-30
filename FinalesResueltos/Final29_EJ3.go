Accion final29 es 
	
	Funcion numerosPerfectos(num1,num2,suma: enteros): booleano
		si num2 = 0 entonces 
			si num1 = suma entonces 
				numerosPerfectos:= verdadero 
			sino 
				numerosPerfectos:= falso 
			fsi 
		sino 
			si num1 mod num2 = 0 entonces 
				numerosPerfectos:= numerosPerfectos(num1,num2-1,suma+num2)
			sino 
				numerosPerfectos:= numerosPerfectos(num1,num2-1,suma)
			fsi 
		fsi 
	FFuncion
 