Accion Ejercicio2 es 
 Ambiente 
	EJEMPLARES = registro 
		id_ejemplar: N(8)
		id_libro: AN(30)
		sucursal: (1...5)
		digital: ("SI","NO")
		disponible: logico 
	FinRegistro
	arch: archivo de EJEMPLARES 
	reg: EJEMPLARES
	A: arreglo[1...5,1...3,1...3] de entero
	mayor, suc_mayor: entero

	Funcion obtener_digital(i: entero): AN
	 	Segun i hacer 
			1: obtener_digital:= "Si"
			2: obtener_digital:= "No"
		FinSegun 
	FinFuncion

	Funcion obtener_disponible(j: entero): AN 
		Segun j hacer 
			1: obtener_disponible:= "Si"
			2: obtener_disponible:= "No"
		FinSegun 
	FinFuncion

	Proceso 
		Abrir E/(arch)
		Leer(arch,reg)
		mayor:=0; suc_mayor:=0
		Para k:= 1 a 5 hacer 	
			Para i:= 1 a 3 hacer 
				Para j:= 1 a 3 hacer 
					A[i,j,k]:=0
				FinPara
			FinPara
		FinPara

		Mientras NFDA(arch) hacer 
			k:= reg.sucursal
			Segun reg.digital hacer 
				"SI": i:=1
				"NO": i:=2
			FinSegun

			Segun reg.disponible hacer 
				verdadero: j:=1
				falso: j:=2
			FinSegun

			A[i,j,k]:= A[i,j,k] + 1
			A[3,j,k]:= A[3,j,k] + 1 
			A[i,3,k]:= A[i,3,k] + 1 
			A[3,3,k]:= A[3,3,k] + 1 
			Leer(arch,reg)
		FinMientras

		Para k:= 1 a 5 hacer 
			Esc("En la sucursal",k)
			Para i:= 1 a 2 hacer 
				Para j:= 1 a 2 hacer 
					Esc("Digital:",obtener_digital(i),"Disponible:",obtener_disponible(j),"Cantidad:",A[i,j,k])
				 	Si A[3,1,k] > mayor entonces 
						mayor:= A[3,1,k] 
						suc_mayor:= k
					FinSi
				FinPara
			FinPara
		FinPara
		//consigna 3 y 4 
		Esc("digitales:",A[1,3,k],"fisicos:",A[2,3,k])
		Esc("En la sucursal 2 hay",A[3,3,2])
		Cerrar(arch)
	FinProceso 
FinAccion
			 	 


		   
		 
	 


 