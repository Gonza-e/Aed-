Accion ej1 es 
	Ambiente 
		sec1,sec2,sal: secuencia de caracteres 
		v1,v2: caracter 

		goles,asist,min,tRojas,cJugadores,cConfederacion: entero 
		bandera: booleano 

		Procedimiento tratarSec2()
			Para i:= 1 hasta 2 hacer 
				goles:= goles*10+conv(v2)
				Avz(sec2,v2)
			FPara 

			Para i:= 1 hasta 2 hacer 
				asist:= asist*10+conv(v2)
				Avz(sec2,v2)
			FPara 

			Para i:= 1 hasta 4 hacer 
				min:= min*10+conv(v2)
				Avz(sec2,v2)
			FPara 
 
			Avz(sec2,v2)
			tRojas:= conv(v2)
			Avz(sec2,v2)
			Avz(sec2,v2)
		FP 

	Proceso 
		Arr(sec1); Arr(sec2); Avz(sec1,v1); Avz(sec2,v2); Crear(sal)

		cConfederacion:=0; cJugadores:=0

		Mientras NFDS(sec1) y NFDS(sec2) hacer 
			Mientras v1 <> "$" hacer 
				Mientras v1 <> "-" hacer 
					Avz(sec1,v1)
				FM 
				Avz(sec1,v1)

				Mientras v1 <> "%" hacer 
					cJugadores:= cJugadores+1
					Mientras v1 <> "-" hacer 
						Grabar(sal,v1)
						Avz(sec1,v1)
					FM 
					Avz(sec1,v1); Grabar(sal,"/")

					goles:=0; aist:=0; min:=0; tRojas:=0
					tratarSec2()

					bandera:= falso 
					si goles > (asist+tRojas)*3 entonces 
						bandera:= verdadero 
					Fsi 

					Mientras v1 <> "*" hacer 
						si bandera entonces 
							Grabar(sal,v1)
						Fsi 
						Avz(sec1,v1)
					FM 
					Avz(sec1,v1)

					si bandera entonces 
						Grabar(sal,"_")
					Fsi 

					Para i:= 1 hasta 5 hacer 
						Avz(sec1,v1)
					FPara 

					si min > 1000 entonces 
						cConfederacion:= cConfederacion+1
					Fsi 
				FM 

				si bandera entonces 
					Grabar(sal,"%")
				Fsi 

				Avz(sec1,v1)
				Esc("Cantidad de jugadores:",cJugadores)
				cJugadores:=0
			FM  
			Avz(sec1,v1)
			Esc("Jugadores que superaron los 1000 minutos:",cConfederacion)
			cConfederacion:=0
		FM 
		Cerrar(sec1); Cerrar(sec2); Cerrar(sal)

	FProceso 
FAccion 


Accion ej2 es 
	Ambiente 
		PARTICIPACIONES = registro 
			conf: AN(20)
			equipo: AN(30)
			pos: AN(20)
			nombre: AN(40)
			asist: N(2)
			goles: N(2)
			edad: N(2)
		Freg 
		Arch: archivo de PARTICIPACIONES ordenado por conf, equipo y pos 
		reg: PARTICIPACIONES

		SALIDA = registro 
			equipo: AN(30)
			pos: AN(20)
			prom: N(2)
		Freg 
		equipazos: archivo de SALIDA 
		sal: SALIDA 

		golesConf,golesEqui,golesPos,masde20,cPos: entero 
		resConf,resEquipo,resPos: AN 

		Procedimiento tratarReg()
			golesPos:= golesPos+reg.goles 
			cPos:= cPos+1
		FP 

		Procedimiento cortePos()
			sal.equipo:= resEquipo 
			sal.pos:= resPos 
			sal.prom:= golesPos DIV cPos 
			Grabar(equipazos,sal)
			goelsEqui:= golesEqui+golesPos
			golesPos:=0; cPos:=0
			resPos:= reg.pos 
		FP  

		Procedimiento corteEquipo()
			cortePos()
			Esc("Equipo:",resEquipo,"Goles:",golesEqui)

			si golesEqui > 20 entonces 
				masde20:= masde20+1 
			Fsi 

			golesConf:= golesConf+golesEqui 
			golesEqui:=0
			resEquipo:= reg.equipo 
		FP

		Procedimiento corteConf()
			corteEquipo()
			Esc("Confederacion:",resConf,"Goles:",golesConf)
			Esc("Equipos con mas de 20 goles:",masde20)

			golesConf:=0; masde20:=0
			resConf:= reg.conf
		FP 

		Procedimiento init()
			masde20:=0; golesEqui:=0; golesConf:=0; golesPos:=0; cPos:=0
			resConf:= reg.conf; resEquipo:= reg.equipo; resPos:= reg.pos
		FP  
	Proceso
		Abrir E/(arch); Abrir S/(equipazos)
		Leer(arch,reg)

		init()

		Mientras NFDA(arch) hacer 
			si resConf <> reg.conf entonces 
				corteConf()
			sino 
				si resEquipo <> reg.equipo entonces 
					corteEquipo()
				sino 
					si resPos <> reg.pos entonces 
						cortePos()
					Fsi 
				Fsi 
			Fsi 

			tratarReg()
			Leer(arch,reg)
		FM 
		corteConf()
		Cerrar(arch); Cerrar(equipazos)
	FProceso 
FAccion 



