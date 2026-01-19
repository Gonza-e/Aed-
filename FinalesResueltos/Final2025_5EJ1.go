Accion ej1 es 
 Ambiente 
	sec: secuencia de caracteres 
	v: caracter 

	A: arreglo[1...2,1...4,1....5] de enteros 
	porc: real 
	len_may1, may1, len_men1, men1, rol, len, men2, len_may2, may2, len_men3, men3: entero 

	Procedimiento inicializar()
		Para k:= 1 hasta 2 hacer 
			Para i:= 1 hasta 4 hacer 
				Para j:= 1 hasta 6 hacer 
					A[k,i,j]:= 0 
				FPara
			FPara 
		FPara 
		porc:= 0; len_may1:= 0; may1:= 0; len_men1:= 0; men1:= 0
		rol:= 0; len:= 0; men2:= 0; len_may2:= 0; may2:= 0
		len_men3:= 0; men3:= 0
	FProcedimiento 

	Procedimiento consigna2()
		si (A[1,1,j] + A[2,1,j]) > may1 entonces 
			len_may1:= j 
			may1:= A[1,1,j] + A[2,1,j]
		Fsi 
	FProcedimiento

	Procedimiento consigna3()
		si (A[1,3,j] < men1) entonces 
			len_men1:= j 
			men1:= A[1,3,j]
		Fsi 
	FProcedimiento

	Procedimiento consigna4()
		si (A[1,i,j] + A[2,j,k]) < men2 entonces 
			rol:= i 
			len:= j 
			men2:= A[1,i,j] + A[2,i,j]
		Fsi 
	FProcedimiento

	Procedimiento consigna5()
		si (A[1,i,j] + A[2,i,j]) < men3 entonces 
			len_men3:= j 
			men3:= A[1,i,j] + A[2,i,j]
		sino 
			si (A[1,i,j] + A[2,i,j]) > may2 entonces 
				len_may2:= j 
				may2:= A[1,i,j] + A[2,i,j]
			Fsi 
		Fsi 
	FProcedimiento

	Funcion obtRol(x: entero): AN 
		segun x hacer 
			1: obtRol:= "Estudiante avamzado"
			2: obtRol:= "Persona graduada"
			3: obtRol:= "Docente"
		Fsegun 
	FFuncion 

	Funcion obtLenguaje(x: entero): AN 
		segun x hacer 
			1: obtLenguaje:= "Python"
			2: obtLenguaje:= "C"
			3: obtLenguaje:= "Java"
			4: obtLenguaje:= "Go"
		Fsegun 
	FFuncion 

	Proceso 
		Arr(sec); Avz(sec,v)
		inicializar()

		Mientras NFDS(sec) hacer 
			Mientras v <> "-" hacer 
				Avz(sec,v)
			FM 
			Avz(sec,v)

			segun v hacer 
				"E": i:= 1 
				"P": i:= 2
				"D": i:= 3 
			Fsegun 
			Mientras v <> "/" hacer 
				Avz(sec,v)
			FM 
			Avz(sec,v)

			segun v hacer 
				"P": j:= 1 
				"C": j:= 2
				"J": j:= 3 
				"G": j:= 4 
			Fsegun 
			Mientras v <> "/" hacer 
				Avz(sec,v)
			FM 
			Avz(sec,v)

			segun v hacer 
				"S": k:= 1 
				"N": k:= 2 
			Fsegun 
			Avz(sec,v); Avz(sec,v)

			A[k,i,j]:= A[k,i,j]
			A[k,4,j]:= A[k,4,j]
			A[k,j,5]:= A[k,j,5]
			A[k,4,5]:= A[k,4,5]
		FM 
		Cerrar(sec)
		Para i:= 1 hasta 4 hacer 
			Para j:= hasta 5 hacer 
				si (i <> 4) y (j <> 5) entonces 
					consigna2()
					consigna3()
					consigna4()
					consigna5()
				Fsi 
			FPara 
		FPara
		
		porc:= ((A[1,4,5] + A[2,4,5])/A[1,2,5]) * 100
		Esc("Porcentaje de personas graduadas con experiencia: ",porc,"%")
		Esc("Lenguaje mas elegido por estudiantes: ",obtLenguaje(len_may1))
		Esc("Lenguaje menos elegido por docentes: ",obtLenguaje(len_men1))
		Esc("Rol menor: ",obtRol(rol),"Lenguaje menos votado: ",obtLenguaje(len))
		Esc("Lenguaje mas elegido: ",obtLenguaje(len_may2),"Lenguaje menos elegido: ",obtLenguaje(len_men3))
	
	FProceso 
FAccion
