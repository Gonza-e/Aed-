Accion ej1 es 
	Ambiente 

	CINE = registro 
		codsuc: AN(3)
		codpel: AN(20)
		fchsuc: N(5)
		turno: (1...3)
		nombre: AN(30)
		cant: N(3)
		tipo: ("3D","2D")
	Freg 
	arch: archivo de CINE ordenado por codsuc, codpel, fchsuc y turno
	reg: CINE 

	SALIDA = registro 
		codsuc: AN(3)
		codpel: AN(20)
		cant: N(7)
	Freg 
	sal: archivo de SALIDA 
	aux,reg_s: SALIDA

	A: arreglo[1...3] de enteros
	total,mayor,cant,cantSuc: entero 
	resPel,resSuc: AN

	Procedimiento tratarReg()
		si reg.tipo = "3D" entonces 
			cant:= cant + reg.cant 
		Fsi 
		si reg.cant > 100 entonces 
			A[reg.turno]:= A[reg.turno] + 1 
		Fsi 
	FProcedimiento

	Procedimiento buscarMayor()
		si cant > mayor entonces 
			aux.codsuc:= resSuc
			aux.codpel:= resPel 
			aux.cant:= cant 
			mayor:= cant 
		Fsi 
	FProcedimiento

	Procedimiento cortePel()
		Esc("Pelicula:",resPel,"Vistas:",cant)
		buscarMayor()
		cantSuc:= cantSuc + cant 
		cant:= 0 
		resPel:= reg.codpel
	FProcedimiento

	Procedimiento corteSuc()
		Esc("Sucursal:",resSuc,"Vistas:",cantSuc)
		total:= total + cantSuc 
		reg_s:= aux; Grabar(sal,reg_s)
		cantSuc:= 0 
		resSuc:= reg.codsuc
	FProcedimiento
	
	Procedimiento init()
		total:= 0; mayor:= 0; cant:= 0; cantSuc:= 0
		resPel:= reg.codpel; resSuc:= reg.codsuc 
		Para i:= 1 hasta 3 hacer 
			A[i]:= 0
		FPara 
	FProcedimiento

	Proceso 
		Abrir E/(arch); Leer(arch,reg)
		Abrir S/(sal)

		init()

		Mientras NFDA(arch) hacer 
			si resSuc <> reg.codsuc entonces 
				corteSuc()
			sino 
				si resPel <> reg.codpel entonces 
					cortePel()
				Fsi 
			Fsi 

			tratarReg()
			Leer(arch,reg)
		FM 
		corteSuc()
		Cerrar(arch); Cerrar(sal)

		Para i:= 1 hasta 3 hacer 
			Esc("Turno:",i,"Peliculas con mas de 100 vistas:",A[i])
		FPara 
	FProceso 
FAccion
