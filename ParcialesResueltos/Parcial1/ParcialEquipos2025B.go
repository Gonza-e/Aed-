Accion ej1 es 
	Ambiente 
		secJ,secE,sal: secuencia de caracteres 
		vJ,vE: caracter 

		goles,tAmarillas,asist,aConf: entero 
		bandera1: booleano

		Procedimiento tratarSecE()
			Para i:= 1 hasta 2 hacer 
				goles:= goles*10 + conv(vE)
				Avz(secE,vE)
			FPara 

			Para i:= 1 hasta 2 hacer 
				asist:= asist*10 + conv(vE)
				Avz(secE,vE)
			FPara 

			Para i:= 1 hasta 4 hacer 
				Avz(secE,vE)
			FPara 

			tAmarillas:= conv(vE)
			Para i:= 1 hasta 3 hacer 
				Avz(secE,vE)
			FPara 

			bandera1:= falso 
			si asist > 0 y goles > 0 entonces 
				bandera1:= verdadero 
			Fsi 

			si tAmarillas > 3 entonces 
				aConf:= aConf+1 
			Fsi 

		FP 

	Proceso 
		init()

		Mientras NFDS(secE) y NFDS(secJ) hacer 
			Mientras vJ <> "-" hacer 
				Avz(secJ,vJ)
			FM 
			Avz(secJ,vJ)

			Mientras vJ <> "$" hacer 
				Mientras vJ <> "-" hacer 
					Esc(sal,vJ)
					Avz(secJ,vJ)
				FM 
				Avz(secJ,vJ); Esc(sal,"/")

				Mientras vJ <> "%" hacer 

					jugador:= jugador+1
					tratarSecE()

					Mientras vJ <> "*" hacer 
						si bandera1 entonces 
							Esc(sal,vJ)
						Fsi 
						Avz(secJ,vJ)
					FM 
					si bandera1 entonces 
						Esc(sal,"_")
					Fsi 
					Avz(secJ,vJ)

					Para i:= 1 hasta 2 hacer 
						edad:= edad*10 + conv(vJ)
						Avz(secJ,vJ)
					FPara

					edadAcum:= edadAcum+edad 
					edad:= 0

					Mientras vJ <> "#" hacer 
						Avz(secJ,vJ)
					FM
					Avz(secJ,vJ)
				FM 
				Esc(sal,"%"); Avz(secJ,vJ)
				prom:= edadAcum/jugador 
				Esc("El promedio de edad es:",prom); edadAcum:=0; jugador:=0
			FM 
			Avz(secJ,vJ)
			Esc("La cantidad de jugadores con tarjetas amarillas es de:",aConf); aConf:=0
		FM
		Cerrar(secJ); Cerrar(secE); Cerrar(sal)
	FProceso 
FAccion 

		







