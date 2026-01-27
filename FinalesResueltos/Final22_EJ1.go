Accion ej1 es 
	Ambiente 

	MOVIMIENTO = registro 
		codRub: AN(6)
		codIns: AN(6)
		nroMov: N(4)
		fechMov: Fecha 
		tipo: ("I","E")
		cant: N(5)
		precio: N(6,2)
		nroDepo: N(7)
	Freg
	arch_m: archivo de MOVIMIENTO ordenado por codRub, codIns y nroMov 
	reg_m: MOVIMIENTO 

	DEPOSITO = registro 
		codRub: AN(6)
		codIns: AN(6)
		ultPrecioSug: N(6,2)
		stockAct: N(5)
		nroDepo: N(7)
	Freg 
	arch_d: archivo de DEPOSITO indexado por codRub y codIns 
	reg_d: DEPOSITO 

	res_ins, res_rub: AN 
	stock,depo,i,j: entero 
	precioAct: real 
	A: arreglo[1...5,1...13] de enteros 

	Procedimiento actIns()
		reg_d.codRub:= res_rub; reg_d.codIns:= res_ins 
		Leer(arch_d,reg_d)
		si EXISTE entonces 
			si reg_d.ultPrecioSug < precioAct entonces 
				reg_d.ultPrecioSug:= precioAct 
			Fsi 
			reg_d.stockAct:= stock 
			Regrabar(arch_d,reg_d)
		sino 
			reg_d.codRubro:= res_rub
			reg_d.codIns:= res_ins 
			reg_d.ultPrecioSug:= precioAct
			reg_d.stockAct:= stock 
			reg_d.nroDepo:= depo 
			Grabar(arch_d,reg_d)
		Fsi 
		stock:= 0; depo:= 0; precioAct:= 0
		res_ins:= reg_d.codIns 
	FProcedimiento 

	Procedimiento actRub()
		res_rub:= reg_d.codRub 
		res_ins:= reg_d.codIns 
		stock:= 0; depo:= 0; precioAct:= 0 
	FProcedimiento 

	Procedimiento actGeneral()
		segun reg_m.tipo hacer 
			"I": stockAct:= stockAct + reg_m.cant 
			"E": stockAct:= stockAct - reg_m.cant
		Fsegun
		precioAct:= reg_m.precio   
		depo:= reg_m.nroDepo 
	FProcedimiento

	Procedimiento cargarArreglo()
		i:= reg_m.nroDepo
		j:= reg_m.fechMov.mes 
		A[i,j]:= A[i,j] + (reg_m.cant*reg_m.precio)
		A[5,j]:= A[5,j] + (reg_m.cant*reg_m.precio)
		A[i,13]:= A[i,j] + (reg_m.cant*reg_m.precio)
		A[5,13]:= A[5,13] + (reg_m.cant*reg_m.precio)
	FProcedimiento

	Procedimiento init()
		stock:= 0; precioAct:= 0; nroDepo:= 0
		i:= 0; j:= 0
		res_ins:= reg_m.codIns; res_rub:= reg_m.codRub 
		Para i:= 1 hasta 5 hacer 
			Para j:= 1 hasta 13 hacer 	
				A[i,j]:= 0 
			FPara 
		FPara
	FProcedimiento

	Proceso
		Abrir E/(arch_m); Leer(arch_m,reg_m)
		Abrir E/S(arch_d)
		
		init()

		Mientras NFDA(arch_m) hacer 
			si reg_m.codRub <> res_rub entonces 
				actRub()
			Fsi 
			si reg_m.codIns <> res_ins entonces 
				actIns()
			Fsi 

			actGeneral()
			cargarArreglo()
			Leer(arch_m,reg_m)
		FM 
		Cerrar(arch_m); Cerrar(arch_d)

		Para i:= 1 hasta 5 hacer 
			Para j:= 1 hasta 13 hacer 
				si i <> 5 y j <> 13 hacer 
					Esc("Nro deposito:",i,"Mes:",j,"Comprado:",A[i,j],"$")
				sino 
					si i = 5 y j <> 13 entonces 
						Esc("Mes:",j,"Comprado:",A[i,j],"$")
					sino 
						si j = 13 y i <> 5 entonces 
							Esc("Deposito:",i,"Comprado:",A[i,j],"$")
						Fsi 
					Fsi 
				Fsi 
			FPara 
		FPara

		Esc("Total general comprado:",A[5,13],"$")

	FProceso 
FAccion

