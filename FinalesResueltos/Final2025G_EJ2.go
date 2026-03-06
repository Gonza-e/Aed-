Accion ej2 (A: arreglo[0...9] de AN) es 
	Ambiente 

	R = registro 
		pais: N(2)
		ame: N(1)
		inte: N(4)
		ipO: AN(15)
		ipD: AN(15)
		nivel: ("B","M","A")
		f: Fecha 
		hora: N(2)
	Freg 
	arch: archivo de R ordenado por pais y ame 
	reg: R 

	am: arreglo[0...9] de entero

	ipMay: AN 
	intPais,intAme,cant,resPais,resAme,maxAme,ame: entero 

	Procedimiento corteAme()
		Esc("Amenaza:",A[resAme],"Intentos:",intAme)
		Esc("IP:",ipMay,"Amenaza:",A[resAme])
		intPais:= intPais+intAme 
		intAme:= 0 
		resAme:= reg.ame 
	FProcedimiento

	Procedimiento cortePais()
		corteAme()
		Esc("Porc de ataques de nivel alto o critico:",cant/intPais,"%")
		si intPais > maxPais entonces 
			maxPais:= intPais 
			pais:= resPais 
		Fsi 
		cant:= 0
		intPais:= 0 
		resPais:= reg.pais 
	FProcedimiento

	Procedimiento tratarReg()
		intAme:= intAme+reg.inte 

		si reg.nivel = "A" o reg.nivel = "M" entonces 
			cant:= cant+reg.inte 
		Fsi 

		am[reg.ame]:= am[reg.ame]+reg.inte 

		si reg.inte > maxIp entonces 
			maxIp:= reg.inte
			ipMay:= reg.ipO 
		Fsi  
	FProcedimiento 

	Procedimiento init()
		Para i:= 0 hasta 9 hacer 
			am[i]:= 0 
		FPara 
		intAme:= 0; intPais:= 0; maxIp:= 0; cant:= 0; ame:= 0; maxAme:= 0
		ipMay:=""; resAme:= reg.ame; resPais:= reg.pais 
	FProcedimiento 

	Proceso 
		Abrir E/(arch); Leer(arch,reg)
		init()

		Mientras NFDA(arch) hacer 
			si resPais <> reg.pais entonces 
				cortePais()
			sino 
				si resAme <> reg.ame entonces 
					corteAme()
				Fsi 
			Fsi 

			tratarReg()
			Leer(arch,reg)
		FM 
		cortePais()
		Cerrar(arch)

		Para i:= 0 hasta 9 hacer 
			si maxAme < am[i] entonces 
				maxAme:= am[i]
				ame:= i 
			Fsi 
		FPara 
		Esc("Amenaza mas frecuente:",A[ame])
		Esc("Pais con mas amenazas:",pais)
	FProceso 
FAccion






		