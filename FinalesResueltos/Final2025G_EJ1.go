Accion ej1 es 
	Ambiente 

	Fecha = registro 
		dia: N(2)
		mes: N(2)
		año: N(4)
	Freg 

	CALIF = registro 
		est: N(6)
		escuela: (1...12)
		mat: ("M","L","C")
		calif: N(2)
		fec: Fecha 
		valido: ("S","N")
	Freg 
	arch: archivo de CALIF
	reg: CALIF 

	A: arreglo[1...2,1...13,1...4] de entero
	prom,prom1,prom2,prom3,prom4,max,min: real
	mejMat,peoMat,i,j,k: entero  

	Funcion obtMat(x:entero): AN 
		segun x hacer 
			1: obtMat:= "Matematica"
			2: obtMat:= "Lengua"
			3: obtMat:= "Ciencia"
		Fsegun 
	FFuncion 

	Procedimiento ind()
		segun reg.mat hacer 
			"M": j:= 1
			"L": j:= 2
			"C": j:= 3 
		Fsegun
	FProcedimiento
	
	Procedimiento cargarAlumno()
		A[1,i,j]:= A[k,i,j]+1
		A[1,13,j]:= A[1,13,j]+1
		A[1,i,4]:= A[1,i,4]+1
		A[1,13,4]:= A[1,13,4]+1
	FProcedimiento

	Procedimiento cargarNota()
		A[2,i,j]:= A[2,i,j]+reg.calif
		A[2,13,j]:= A[2,13,j]+reg.calif
		A[2,i,4]:= A[2,i,4]+reg.calif
		A[2,13,4]:= A[2,13,4]+reg.calif
	FProcedimiento

	Procedimiento con2()
		prom2:= A[2,i,4]/A[1,i,4]

		Esc("Escuela:",i,"Promedio:",prom2,"Promedio nacional:",prom)
	FProcedimiento

	Procedimiento con1()
		prom1:= A[2,13,j]/A[1,13,j]

		Esc("Escuela:",i,"Materia:",obtMat(j),"Promedio:",prom1,"Promedio nacional:",prom)
	FProcedimiento

	Procedimiento con3()
		prom3:= A[2,i,j]/A[1,i,j]

		si prom3 > max entonces 
			max:= prom3 
			mejMat:= j 
		Fsi 

		si prom3 < min entonces 
			min:= prom3 
			peoMat:= j 
		Fsi 
	FProcedimiento

	Procedimiento con4()
		prom4:= A[2,i,j]/A[2,i,j]
		
		si prom4 > prom entonces
			cont:= cont+1 
		Fsi 

		si cont >= 2 entonces 
			Esc("La escuela:",i,"supera al promedio nacional en dos materias")
		Fsi 
	FProcedimiento

	Proceso 
		Abrir E/(arch); Leer(arch,reg)
		prom:= 0; prom1:= 0; prom2:= 0; prom3:= 0; prom4:= 0
		mejMat:= 0; peoMat:= 0; max:= 0; min:= 0

		Mientras NFDA(arch) hacer 
			si reg.valido = "S" entonces 
				ind()
				i:= reg.escuela
				cargarAlumno()
				cargarNota()
			Fsi 
			Leer(arch,reg)
		FM 
		Cerrar(arch)

		prom:= A[2,13,4]/A[1,13,4]
		Para i:= 1 hasta 12 hacer 
			Para j:= 1 hasta 3 hacer  
				con1()
				con3()
				con4()
			FPara 
			Esc("Escuela:",i,"Materia con mejor promedio:",obtMat(mejMat),"Materia con peor promedio:",obtMat(peoMat))
			max:= 0; min:= HV 
			con2()
		FPara 

	FProceso 
FAccion
				
