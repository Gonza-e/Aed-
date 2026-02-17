Accion ej es 
	Ambiente 

	sec,sal: secuencia de caracteres 
	s: caracter 

	letras: {"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
	numeros: {"0","1","2","3","4","5","6","7","8","9"}
	vocales: {"A","E","I","O","U"}

	voc: arreglo[1...5] de AN 
	cant: arreglo[1...3] de entero //1: total, 2: validos, 3: no validos
	prom: real
	let,ant: AN 
	pos,num,pal: entero 

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
	FFuncion 

	Funcion indice(A: arreglo[1...5] de AN; let: AN)
		Para i:= 1 hasta 5 hacer 
			si A[i] = let entonces 
				indice:= i 
			Fsi 
		FPara 
	FFuncion 

	Procedimiento cargaVoc()
		voc[1]:= "A"; voc[2]:= "E"; voc[3]:= "I"; voc[4]:= "O"; voc[5]:= "U"
	FProcedimiento 

	Procedimiento init()
		Para i:= 1 hasta 5 hacer 
			cant[i]:= 0
		FPara 
		prom:= 0; pos:= 0; num:= 0; pal:= 0
		ant:= " "
	FProcedimiento
	
	Procedimiento verificar()
		si s = " " entonces 
			Esc(sal,s)
			cant[1]:= cant[1]-1
		Fsi 

		si s EN letras entonces 
			si s EN vocales entonces 
				pos:= (indice(voc,s)-2) mod 5
				let:= voc[pos] 
			sino 
				let:= "1"
			Fsi 
			Esc(sal,let)
			cant[2]:= cant[2]+1
		sino 
			si s EN numeros entonces 
				si conv(s) mod 2 = 0 entonces 
					num:= conv(s) * 2 
				sino 
					num:= conv(s) ** 3
				Fsi 
				Esc(sal,num)
				cant[2]:= cant[2]+1
			sino 
				Esc(sal,"@")
				cant[3]:= cant[3]+1
			Fsi 
		Fsi 
	FProcedimiento

	Procedimiento contarPalabras()
		si ant = " " y s <> " " entonces 
			pal:= pal+1
		Fsi 
	FProcedimiento

	Proceso 
		Arr(sec); Avz(sec,s)
		init()
		cargaVoc()

		Mientras NFDS(sec) hacer 
			contarPalabras()
			verificar()
			Avz(sec,s)
		FM 
		Cerrar(sec)

		Esc("Cantidad de palabras: ",pal)
		prom:= (cant[3]/cant[1]) * 100
		Esc("Porcentaje de caracteres invalidos por sobre el total: ",prom,"%")

		si cant[3] < cant[2]/2 entonces
			Esc("Mensaje curado")
		sino 
			si cant[3] = cant[2] entonces 
				Esc("Mensaje alerta")
			sino 
				Esc("Mensaje corrupto")
			Fsi 
		Fsi 

	FProceso 
FAccion

