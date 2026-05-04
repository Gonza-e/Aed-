Accion ej1 es 
	Ambiente 
		sec1,sec2,sal: secuencia de caracteres 
		v1,v2: caracter

		cTP,cR,cP,cE,cF,cCanciones,cMayor,acumCanciones: entero
		pMayor: AN 
		esRock: booleano

	Proceso 
		Arr(sec1); Arr(sec2); Crear(sal)
		Avz(sec1,v1); Avz(sec2,v2)

		cTP:=0; cR:=0; cP:=0; cE:=0; cCanciones:=0; cMayor:=0; acumCanciones:=0
		gMayor:= ""; genero:= ""
		
		Mientras NFDS(sec1) y NFDS(sec2) hacer 
			segun v1 hacer 
				"R": cR:= cR+1
				"P": cP:= cP+1
				"E": cE:= cE+1
				"F": cF:= cF+1
			Fsegun 
			
			genero:= v1
			esRock:= falso 
			si v1 = "R" entonces 
				esRock:= verdadero 
			Fsi 
			Avz(sec1,v1); Avz(sec1,v1)

			Mientras v1 <> "+" hacer 
				si esRock entonces 
					Grabar(sal,v1)
				Fsi 
				Avz(sec1,v1)
			FM 

			Para i:= 1 hasta 3 hacer 
				Mientras v1 <> "+" hacer 
					Avz(sec1,v1)
				FM 
			FPara 

			Mientras v1 <> "?" hacer 
				cCanciones:= cCanciones*10 + conv(v1)
				Avz(sec1,v1)
			FM 
			Avz(sec1,v1)
		
			si cCanciones > cMayor entonces 
				cMayor:= cCanciones
				pMayor:= genero 
			Fsi 

			Para i:= 1 hasta cCanciones hacer 
				Mientras v2 <> "#" hacer 
					si esRock entonces 
						Grabar(sal,v2)
					Fsi 
					Avz(sec2,v2)
				FM 

				Para i:= 1 hasta 2 hacer 
					Mientras v2 <> "#" hacer 
						Avz(sec2,v2)
					FM 
					Avz(sec2,v2)
				FPara 

				Mientras v2 <> "/" hacer 
					si esRock entonces 
						Grabar(sal,v2)
					Fsi 
					Avz(sec2,v2)
				FM 
				Avz(sec2,v2)

				si esRock entonces 
					Grabar(sal,"+")
				Fsi 

				acumCanciones:= acumCanciones+cCanciones; cCanciones:= 0
			FPara 

			si esRock entonces 
				Grabar(sal,"#")
			Fsi 
		FM
		Cerrar(sec1); Cerrar(sec2); Cerrar(sal)

		cTP:= cR+cE+cF+cP
		Esc("Porcentaje de playlists de genero Rock:",(cR/cTP)*100,"%")
		Esc("Porcentaje de playlists de genero Folklore:",(cF/cTP)*100,"%")
		Esc("Porcentaje de playlists de genero Electronica:",(cE/cTP)*100,"%")
		Esc("Porcentaje de playlists de genero Pop:",(cP/cTP)*100,"%")

		Esc("Promedio de canciones por playlist:",acumCanciones/cTP)

		Esc("El genero de la playlist con mayor cantidad de canciones es:",gMayor)

	FProceso 
FAccion 

Accion ej2 es 
	Ambiente 
		Fecha = registro 
			dia: N(2)
			mes: N(2)
			año: N(4)
		Freg
		CANCIONES = registro 
			gen: AN(20)
			artist: AN(70)
			album: AN(70)
			nombre: AN(12)
			cod: N(15)
			f_publicacion: Fecha
			c_reproducciones: N(20)
		Freg
		arch: archivo de CANCIONES ordenado por genero, artist, album, nombre y cod 
		reg: CANCIONES

		SALIDA = registro 
			artista: AN(70)
			cantCanciones: N(3)
		Freg 
		sal: archivo de SALIDA 
		regSal: SALIDA 

		resGenero, resArtista, artista: AN 
		repArtista, repGenero, cMayor, repTotal, cCanciones: entero 

		Procedimiento tratarReg()
			repArtista:= repArtista+reg.c_reproducciones
			cCanciones:= cCanciones+1 
		FP 

		Procedimiento corteArtista()
			si repArtista > cMayor entonces 
				cMayor:= repArtista
				artista:= resArtista
			Fsi 

			regSal.artista:= resArtista; regSal.cantCanciones:= cCanciones
			Grabar(sal,regSal)

			repGenero:= repGenero+repArtista 
			repArtista:= 0 
			cCanciones:= 0 
			resArtista:= reg.artist 
		FP 

		Procedimiento corteGenero()
			corteArtista()
			Esc("Genero: ",resGenero," Cantidad de reproducciones: ",repGenero)
			repTotal:= repTotal+repGenero 
			repGenero:= 0 
			resGenero:= reg.genero 
		FP 

	Proceso 
		Abrir E/(arch); Abrir S/(sal)
		Leer(arch,reg)

		cMayor:=0; repArtista:=0; repGenero:=0; repTotal:=0; cCanciones:=0
		resArtista:= reg.artist; resGenero:= reg.genero 
		artista:= ""

		Mientras NFDA(arch) hacer 
			si resGenero <> reg.genero entonces 
				corteGenero()
			sino 
				si resArtista <> reg.artist entonces 
					corteArtista()
				Fsi 
			Fsi 

			tratarReg()
			Leer(arch,reg)
		FM 
		corteGenero()
		Cerrar(sal); Cerrar(arch)

		Esc("Cantidad total de reproducciones: ",repTotal)
		Esc("Artista con mayor cantidad de reproducciones: ",artista)
	FProceso 
FAccion 

