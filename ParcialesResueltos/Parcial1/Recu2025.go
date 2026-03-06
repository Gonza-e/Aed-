Accion ej1 es 
	Ambiente 
	
	hist,datos,sal: secuencia de caracteres 
	h,d: caracter
	edad,c1,c2: entero 
	bandera1,bandera2: booleano 
	vocal = ("A","E","I","O","U")

	Funcion conv(x:AN):entero 
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
	FF 

	Procedimiento tratar()
		si h = "O" entonces 
			bandera1:= verdadero 
		Fsi 
		
		Mientras h<>"#" hacer 
			si bandera1 entonces 
				Esc(sal,h)
			Fsi 
			Avz(hist,h)
		FM 
		Esc(sal,"#")
		Avz(hist,h)
	FP 

	Procedimiento init()
		edad:= 0; c1:= 0; c2:= 0; bandera1:= falso; bandera2:= falso 
	FP 

	Proceso 
		Arr(hist); Arr(datos); Avz(hist,h); Avz(datos,d)
		init()

		Mientras NFDS(hist) y NFDS(datos) hacer 
			tratar()

			Mientras d<>"@" hacer 
				si d EN vocal entonces 
					bandera2:= verdadero
				Fsi 

				Para i:= 1 hasta 3 hacer 
					si bandera1 y bandera2 hacer 
						Esc(sal,d)
					Fsi 
					Avz(datos,d)
				FPara 

				Para i:= 1 hasta 2 hacer 
					edad:= edad*10 + conv(d)
					Avz(datos,d)
				FPara 

				Mientras d<>"#" hacer 
					si bandera1 y bandera2 entonces
						Esc(sal,d)
					Fsi 
					Avz(datos,d)
				FM
				
				si edad >= edadUsu entonces 
					c1:= c1+1
				Fsi 

				bandera2:= falso 
			FM
			bandera1:= falso 
			Esc("Cantidad de atletas que igualan o superan la edad:",c1); c1:= 0
			Avz(datos,d)
		FM 
	FProceso 
FAccion 
					
				
