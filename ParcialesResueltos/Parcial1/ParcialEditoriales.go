Accion ej1 es 
	Ambiente	
		sec1,sec2,sal: secuencia de caracteres //sec1: secuencia de editoriales, sec2: secuencia de libros
		v1,v2: caracter 

		cLibros,cPalabras,pag,prov,provMayor,cLibrosMayor: entero 
		bandera: booleano 
		ant,tema: AN

		Funcion conv(x:caracter):entero 
			segun x hacer 
				"0": conv:= 0
				"1": conv:= 1
				"2": conv:= 2
				"3": conv:= 3
				"4": conv:= 4
				"5": conv:= 5
				"6": conv:= 6
				"7": conv:= 7
				"8": conv:= 8
				"9": conv:= 9
			Fsegun 
		FFuncion 

		Procedimiento contarPalabras()
			Mientras v2 <> "#" hacer 
				si ant = "" y v2 <> "" entonces 
					cPalabras:= cPalabras+1
				Fsi 
				ant:= v2; Avz(sec2,v2)
			FM 
			Avz(sec2,v2)
		FP

		Procedimiento init()
			cLibros:=0; cPalabras:=0; pag:=0; prov:=0
			ant:= ""; bandera:= falso; tema:= ""
		FP

	Proceso 
		Arr(sec1); Arr(sec2); Crear(sal)
		Avz(sec1,v1); Avz(sec2,v2)

		init()

		Mientras NFDS(sec1) y NFDS(sec2) hacer 
			repetir 
				Avz(sec1,v1)
			Hasta (v1 = "&")

			Para i:= 1 hasta 2 hacer 
				prov:= prov*10 + conv(v1)
				Avz(sec1,v1)
			FPara 

			Mientras v2 <> "@" hacer 
				cLibros:= cLibros+1 
				tema:= v2
				Avz(sec2,v2)

				Para i:= 1 hasta 3 hacer 
					pag:= pag*10 + conv(v2)
					Avz(sec2,v2)
				FPara 

				bandera:= falso 
				si tema = "L" y (pag > 200 y pag <= 300) entonces 
					bandera:= verdadero 
				Fsi 

				Mientras v2 <> "&" hacer 
					si bandera entonces 
						Grabar(sal,v2)
					Fsi 
					Avz(sec2,v2)
				FM 
				si bandera entonces 
					Grabar(sal,"$")
				Fsi 
				Avz(sec2,v2)

				contarPalabras()
				ant:= ""
			FM

			Esc("Longitud promedio de los resumenes:",cPalabras/cLibros,"palabras")
			si cLibros > cLibrosMayor entonces 
				cLibrosMayor:= cLibros 
				provMayor:= prov 
			Fsi 
			prov:=0; cLibros:=0; cPalabras:=0
			
		FM 
		Cerrar(sec1); Cerrar(sec2); Cerrar(sal)

		Esc("Provincia con la editorial con mayor cantidad de libros:",provMayor,"Cantidad de libros:",cLibrosMayor)
	FProceso 
FAccion 





