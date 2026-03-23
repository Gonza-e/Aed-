Accion ej1 es 
	Ambiente 

		fec,dat,sal: secuencia de caracteres 
		f,d: caracter 

		mesUsu,mes,cV,c1,c2,c3,c4,c5,monto,cat: entero 
		p1,p2,p3,p4,p5: AN 
		bandera1,bandera2: booleano
		desc: real 

		Funcion conv(x:AN): entero
			segun x hacer  
				"1": conv:= 1
				"2": conv:= 2
				"3": conv:= 3
				"4": conv:= 4
				"5": conv:= 5
				"6": conv:= 6
				"7": conv:= 7
				"8": conv:= 8
				"9": conv:= 9
				"0": conv:= 0
			Fsegun 
		FF 

		Procedimiento tratarFecha()
			Para i:= 1 hasta 2 hacer 
				Avz(fec,f)
			FPara 

			Para i:= 1 hasta 2 hacer 
				mes:= mes*10 + conv(f)
				Avz(fec,f)
			FPara 

			Para i:= 1 hasta 5 hacer 
				Avz(fec,f)
			FPara
			
			si mesUsu = mes entonces 
				bandera1:= verdadero 
			Fsi 
		FP 

		Procedimiento obtPatente()
			Para i:= 1 hasta 5 hacer 
				segun i hacer 
					1: p1:= d 
					2: p2:= d 
					3: p3:= d 
					4: p4:= d 
					5: p5:= d 
				Fsegun 
				Avz(dat,d)
			FPara 
		FP 

		Procedimiento init()
			mes:= 0; cV:= 0; c1:= 0; c2:= 0; c3:= 0; c4:= 0; c5:= 0; monto:= 0
			p1:= "", p2:= "", p3:= "", p4:= "", p5:= "", cat:= 0
			bandera1:= falso; bandera2:= falso 
		FProcedimiento

	Proceso 
		Arr(fec); Arr(dat); Avz(fec,f); Avz(dat,d)
		Crear(sal)

		init()

		Esc("Ingrese un mes:"); Leer(mesUsu)

		Mientras NFDS(fec) y NFDS(dat) hacer 
			tratarFecha()
			Mientras d <> "-" hacer 
				cV:= cV+1; cat:= conv(d); monto:= obtenerMonto(d)
				Avz(dat,d)

				si d = "S" entonces 
					bandera2:= verdadero
				Fsi 
				Avz(dat,d)

				Para i:= 1 hasta 4 hacer 
					Avz(dat,d)
				FPara 

				obtPatente()

				si bandera1 y bandera2 entonces 
					Esc(sal,p1); Esc(sal,p2); Esc(sal,p3); Esc(sal,p4); Esc(sal,p5)
					Esc(sal,"?")
				Fsi 

				si bandera2 entonces 
					monto:= monto - (monto * obtenerDescuento(p1,p2,p3,p4,p5))
				Fsi 

				segun cat hacer 
					1: c1:= c1+monto
					2: c2:= c2+monto
					3: c3:= c3+monto
					4: c4:= c4+monto
					5: c5:= c5+monto
				Fsegun 

				bandera1:= falso; bandera2:= falso 
			FM
			mes:= 0
			Avz(dat,d)

			Esc("Categoria 1:",c1); c1:= 0 
			Esc("Categoria 2:",c2); c2:= 0
			Esc("Categoria 3:",c3); c3:= 0
			Esc("Categoria 4:",c4); c4:= 0 
			Esc("Categoria 5:",c5); c5:= 0 
		FM 
		Cerrar(dat); Cerrar(fec); Cerrar(sal)

		Esc("Cantidad total de vehiculos que pasaron:",cV)
	FProceso 
FAccion 

Accion ej2 es 
	Ambiente 
		PEAJE = registro 
			año: N(4)
			mes: N(2)
			dia: N(2)
			cat: N(1)
			patente: AN(5)
			origen: ("Rcia","Ctes")
		Freg 
		arch: archivo de PEAJE ordenado por año,mes,dia,cat y patente 
		reg: PEAJE 

		SALIDA = registro 
			año: N(4)
			mes: N(2)
			cRcia: N(5)
			cCtes: N(5)
		Freg 
		archSal: archivo de SALIDA 
		sal: SALIDA 

		cMesR,cMesC,cAñoR,cAñoC,cT: entero 
		resAño,resMes: entero 

		Procedimiento tratarReg()
			segun reg.origen hacer 
				"Rcia": cMesR:= cMesR+1
				"Ctes": cMesC:= cMesC+1
			Fsegun
		FP 

		Procedimiento corteMes()
			Esc("Mes:",resMes,"Resistencia:",cMesR,"Corrientes:",cMesC)
			cAñoC:= cAñoC+cMesC 
			cAñoR:= cAñoR+cMesR
			si (cMesR+cMesC) > 100000 entonces 
				sal.año:= resAño; sal.mes:= resMes; sal.cRcia:= cMesR; sal.cCtes:= cMesC 
				Grabar(archSal,sal)
			Fsi 
			cMesC:= 0; cMesR:= 0 
			resMes:= reg.mes 
		FP 

		Procedimiento corteAño()
			corteMes()
			Esc("Año:",resAño,"Resistencia:",cAñoR,"Corrientes:",cAñoC)
			cT:= cT+(cAñoR+cAñoC)
			cAñoR:= 0; cAñoC:= 0 
			resAño:= reg.año 
		FP 

		Procedimiento init()	
			cT:= 0; cMesC:= 0; cMesR:= 0; cAñoC:= 0; cAñoR:= 0
			resAño:= reg.año; resMes:= reg.mes 
		FP 

	Proceso 
		Abrir E/(arch); Abrir S/(archSal); Leer(arch,reg)

		init()

		Mientras NFDA(arch) hacer 
			si resAño <> reg.año entonces 
				corteAño()
			sino 
				si resMes <> reg.mes entonces 
					corteMes()
				Fsi 
			Fsi 

			tratarReg()
			Leer(arch,reg)
		FM 
		corteAño(); Cerrar(arch); Cerrar(archSal)

		Esc("La cantidad total de vehiculos que pasaron es:",cT)
	FProceso 
FAccion 








