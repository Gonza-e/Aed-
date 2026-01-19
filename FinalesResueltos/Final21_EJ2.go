Accion Ejercicio2 es 
	Ambiente

	pal, sal: secuencia de caracteres 
	p: caracter 

	Funcion conv(x: AN): entero 
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

	Funcion convHex(x: entero): AN 
		segun x hacer 
			0: convHex:= "0"
			1: convHex:= "1"
			10: convHex:= "2"
			11: convHex:= "3"
			100: convHex:= "4"
			101: convHex:= "5"
			110: convHex:= "6"
			111: convHex:= "7"
			1000: convHex:= "8"
			1001: convHex:= "9"
			1010: convHex:= "A"
			1011: convHex:= "B"
			1100: convHex:= "C"
			1101: convHex:= "D"
			1110: convHex:= "E"
			1111: convHex:= "F"
		Fsegun 
	FFuncion 

	dig: entero 

	Proceso
		Arr(pal); Avz(pal,p)
		Crear(sal)

		dig:= 0 

		Para i:= 1 hasta 12 hacer 
			Para j:= 1 hasta 4 hacer 
				dig:= dig*10 + conv(p)
				Avz(pal,p)
			FPara 

			Esc(sal,convHex(dig))

			si i mod 8 = 0 entonces 
				Esc(sal,"-")
			fsi 

			dig:= 0
		FPara
		
		Cerrar(sal); Cerrar(pal)

	FProceso 
FAccion
