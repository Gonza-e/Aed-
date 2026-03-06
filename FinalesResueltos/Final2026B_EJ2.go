Accion ej2 es 
	Ambiente 

	DATO = registro 
		uni: AN(30)
		car: AN(30)
		legajo: N(5)
		prom: N(2,2)
		cA: N(3)
		cR: N(3)
		cond: ("R","L")
	Freg 
	arch: archivo de DATO ordenado por uni y car //tomo como un archivo ordenado para poder usar corte de control
	reg: DATO

	may,men,pCar,pUni,porc1,mayUni: real 
	aMay,aMen,cCar,cUni,cReg,cRiesgo,cRiesgoUni,c30: entero 
	universidad,resUni,resCar: AN 

	Procedimiento tratarReg()
		si reg.prom > may entonces 
			may:= reg.prom 
			aMay:= reg.legajo 
		Fsi 
		si reg.prom < men entonces 
			men:= reg.prom 
			aMen:= reg.legajo 
		Fsi 

		cCar:= cCar+1 
		pCar:= pCar+reg.prom 

		si reg.cond = "R" entonces 
			cReg:= cReg+1 
		Fsi 

		si riesgo(reg.prom,reg.cA,reg.cR) entonces 
			cRiesgo:= cRiesgo+1 
		Fsi 

		si reg.cR > reg.cA entonces 
			c30:= c30+1 
		Fsi 
	FP 

	Procedimiento corteCar()
		porc1:= (cReg/cCar)*100
		Esc("Carrera:",resCar,"Porcentaje de alumnos regulares",porc1,"%")
		Esc("Estudiante con mejor promedio:",aMay)
		Esc("Estudiante con peor promedio:",aMen)
		Esc("Carrera:",resCar,"Estudiantes en situacion de riesgo:",cRiesgo)
		si c30 > (cCar*0.3) entonces 
			Esc("Universidad:",resUni,"Carrera:",resCar,"Con mas porcentaje de alumnos con desaprobados que aprobados")
		Fsi 

		cRiesgoUni:= cRiesgoUni+cRiesgo
		cUni:= cUni+cCar; pUni:= pUni+pCar
		cCar:= 0; pCar:= 0; c30:= 0; cRiesgo:= 0; cReg:= 0; aMen:= 0; aMay:= 0
		men:= HV; may:= 0

		resCar:= reg.car 
	FP 

	Procedimiento corteUni()
		corteCar()
		Esc("Universidad:",resUni,"Estudiantes en situacion de riesgo:",cRiesgoUni)
		pUni:= (pCar/cCar)
		si pUni > mayUni entonces 
			universidad:= resUni 
			mayUni:= pUni 
		Fsi 

		pUni:= 0; cRiesgoUni:= 0; cUni:= 0
		resUni:= reg.uni 
	FP 

	Funcion riesgo(x:real; m1,m2:entero): booleano 
		si (x<6) y (m2>m1) entonces 
			riesgo:= verdadero 
		Fsi 
	FF 

	Procedimiento init()
		cCar:= 0; cUni:= 0; pUni:= 0; pCar:= 0; porc1:= 0
		c30:= 0; cRiesgo:= 0; cRiesgoUni:= 0; cReg:= 0 
		men:= HV; may:= 0; mayUni:= 0
		aMay:= 0; aMen:= 0
		resUni:= reg.uni; resCar:= reg.car 
	FP 

	Proceso 
		Abrir E/(arch); Leer(arch,reg)
		init()

		Mientras NFDA(arch) hacer 
			si resUni<>reg.uni entonces 
				corteUni()
			sino 
				si resCar<>reg.car entonces 
					corteCar()
				Fsi 
			Fsi 

			tratarReg()
			Leer(arch,reg)
		FM 
		corteUni(); Cerrar(arch)

		Esc("La facultad con mejor promedio general es:",universidad)
	FProceso 
FAccion 
	

	
	

