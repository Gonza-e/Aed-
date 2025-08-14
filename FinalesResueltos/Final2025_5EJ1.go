Accion ejercicio1 es 
 Ambiente 
	sec: secuencia de caracteres 
	v: caracter 

	A: arreglo[1...2,1...4,1...5] de experiencia 
	rol_men, len_men, len_may_est, len_may_gen, len_men_gen, men, i, j: entero 
	exp, est: logico
	cont: entero 

	Proceso 
		ARR(sec); AVZ(sec,v)
		rol_men:= 0; len_men:= HV; len_may_est:= 0; len_may_est1:= 0; len_may_gen:= 0; len_men_gen:= 0; men:= 0; i:= 0; j:= 0
		exp:= falso; est:= falso 

		Para k:= 1 a 2 hacer 
			Para i:= 1 a 4 hacer 
				Para j:= 1 a 5 hacer 
					A[i,j,k]:= 0
				FPara
			FPara
		FPara 

		Mientas NFDS(sec) hacer 
			Mientras v <> "-" hacer 
				AVZ(sec,v)
			FM 
			AVZ(sec,v)

			Segun v hacer 
				= "E": i:= 1
				= "P": i:= 2 
				= "D": i:= 3 
			FSegun

			Mientras v <> "?" hacer 
				AVZ(sec,v)
			FM 
			AVZ(sec,v)

			Segun v hacer 
				= "P": j:= 1 
				= "C": j:= 2
				= "J": j:= 3 
				= "G": j:= 4 
			FSegun 

			Mientras v <> "?" hacer 
				AVZ(sec,v)
			FM 
			AVZ(sec,v)

			Si v = S entonces 
				k:= 1
			Sino 
				k:= 2
			FSi 

			Para i:= 1 a 3 hacer 
				AVZ(sec,v)
			FPara 

			A[i,j,k]:= A[i,j,k] + 1 
			A[i,5,k]:= A[i,5,k] + 1 
			A[4,j,k]:= A[4,j,k] + 1 
			A[4,5,k]:= A[4,5,k] + 1 
	
			Si exp y est entonces 
				cont:= cont + 1
				exp:= falso; est:= falso  
			FSi 
		FM 
		Para k:= 1 a 2 hacer 
			Para i:= 1 a 4 hacer 
				Para j:= 1 a 5 hacer 
					Si (i = 1) y (j <> 5) entonces 
						Si len_may_est < A[i,j,k] entonces
							len_may_est:= A[i,j,k]
							len_may_est1:= j
						Fsi
					FSi 

					Si (i <> 4) o (j <> 5) entonces 
						Si len_men > A[i,j,k] entonces
							len_men:= A[i,j,k]
							rol_men:= i 
							len_men1:= j 
						Fsi 
					Fsi 

					Si ((i <> 4) o (j <> 5)) y (k = 1) entonces
						Si len_men_doc > A[i,j,k] entonces 
							len_men_doc:= A[i,j,k]
							len_men_doc1:= j 
						Fsi 
					Fsi 
					
					Si i = 5 entonces 
						Si len_may_gen < A[i,j,k] entonces 
							len_may_gen:= A[i,j,k] 
							len_may_gen1:= j 
						Fsi
						Si len_men_gen > A[i,j,k] entonces 
							len_men_gen:= A[i,j,k] 
							len_men_gen1:= j 
						Fsi 
					Fsi 
				FPara 
			FPara
		FPara
					




						

