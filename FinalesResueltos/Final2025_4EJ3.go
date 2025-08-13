accion ejercicio3 es 
 Ambiente 
	Funcion casiPerfecto(num1,suma,num2): booleano
		si num2 = 0 entoncecs 
			si suma = (num1 - 1) entonces 
				casiPerfecto:= verdadero 
			sino 
				casiPerfecto:= falso 
			fsi 
		sino 
			si (num1 mod num2) = 0 entonces
				casiPerfecto:= casiPerfecto(num1,suma+num2,num2-1)
			sino 
				casiPerfecto:= casiPerfecto(num1,suma,num2-1)
			fsi 
		fsi 
	FFuncion

	A,B,cant: entero 

	Proceso

		Esc("Ingrese primer valor del rango"); Leer(A)
		Esc("Ingrese segundo valor del rango"); Leer(B)

		si A < B entonces 
			Para i:= A hasta B hacer 
				Si casiPerfecto(i,0,i-1) entonces 
					cant:= cant + 1 
				fsi 
			FPara 
		sino 
			Esc("El rango no es valido")
		fsi 

		Esc("La cantidad de numeros casi perfectos es de: ",cant)
	FProceso 
FAccion 

