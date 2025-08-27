accion ejercicio2 es 
 ambiente 
	Funcion numeroPrometido(num1,num2,num3,num4,suma1,suma2):booleano
		si (num2 = 0) y (num4 = 0) entonces 
			si (suma1 + 1 = num3) y (suma2 + 1 = num1)
				numeroPrometido:= verdadero
			sino 
				numeroPrometido:= falso 
			fsi 
		sino 
			si num2 > 0 entonces 
				si num1 mod num2 = 0 entonces 
					numeroPrometido:= numeroPrometido(num1,num2-1,num3,num4,suma1+num2,suma2)
				sino 
					numeroPrometido:= numeroPrometido(num1,num2-1,num3,num4,suma1,suma2)
				fsi 
			fsi 
			si num4 > 0 entonces 
				si num3 mod num4 = 0 entonces 
					numeroPrometido(num1,num2,num3,num4-1,suma1,suma2+num3)
				sino 
					numeroPrometido(num1,num2,num3,num4-1,suma1,suma2)
				fsi 
			fsi 
	FFuncion 

	A,B: entero 

	Proceso 
		Esc("Ingrese extremo inferior"); Leer(A)
		Esc("Ingrese extremo superior"); Leer(B)

		Para i:= A hasta B - 1 hacer 
			Para j:= i + 1 hasta B hacer 
				si numeroPrometido(i,i-1,j,j-1,0,0) entonces 
					Esc("El par de numeros es prometido: ","(",i,",",j,")")
				fsi 
			fpara 
		fpara 
	FProceso 
FAccion

