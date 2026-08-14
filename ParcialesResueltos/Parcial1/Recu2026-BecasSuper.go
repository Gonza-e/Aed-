Accion ej1 es 
	Ambiente 
		datos,sal: secuencia de caracter
		d: caracter 
		ing: secuencia de enteros
		n,d1,d2,cF,cM,total: entero 

		cumple: booleano 

		Procedimiento primerosDos()
			segun sexo hacer 
				"F": Esc(sal,"27-"); cF:= cF+1
				"M": Esc(sal,"20-"); cM:= cM+1
			Fsegun 
		FP 

		Funcion codigo(digito1,digito2:entero): AN 
			si digito1 mod 2 = 0 y digito2 mod 2 = 0 entonces 
				codigo:= "-1"
			sino 
				si digito1 mod 2 <> 0 y digito2 mod 2 = 0 entonces 
					codigo:= "-2"
				sino 
					si digito1 mod 2 = 0 y digito2 mod 2 <> 0 entonces 
						codigo:= "-3"
					sino 
						codigo:= "-4"
					Fsi 
				Fsi 
			Fsi 
		FFuncion  
		
		Procedimiento init()
			d1:= 0; d2:= 0; cF:= 0; cM:= 0; total:= 0
		FP 

	Proceso 
		Arr(datos); Arr(ing); Avz(datos,d); Avz(ing,n)
		Crear(sal)
		init()

		Mientras d <> "*" hacer 
			cumple:= falso 
			si n <= 10000 entonces 
				cumple:= verdadero 
			Fsi 
			Avz(ing,n)

			Mientras d <> "$" hacer 
				Repetir
					si cumple entonces 
						Esc(sal,d)
					Fsi
					Avz(datos,d)
				Hasta (d = "|")

				sexo:= d; Avz(datos,d); Avz(datos,d)

				si cumple entonces 
					primerosDos()
				Fsi 

				Para i:= 1 hasta 8 hacer 
					segun i hacer 
						1: d1:= convEntero(d)
						8: d2:= convEntero(d)
					Fsegun 
					si cumple entonces 
						Esc(sal,d)
					Fsi 
					Avz(datos,d)
				FPara 

				si cumple entonces 
					Esc(sal,codigo(d1,d2))
				Fsi 
				Avz(datos,d); Esc(sal,d); Avz(datos,d)

				Mientras d <> "|" hacer 
					si cumple entonces 
						Esc(sal,d)
					Fsi 
					Avz(datos,d)
				FM 
				Avz(datos,d)

				Mientras d <> "$" hacer 
					Avz(datos,d)
				FM 
			FM 
			Avz(datos,d)
		FM 
		Cerrar(sal); Cerrar(ing); Cerrar(datos)
		total:= cF+cM 
		Esc("Porcentaje de masculinos:",(cM/total)*100,"%")
		Esc("Porcentaje de femeninos: ",(cF/total)*100,"%")
	FProceso
FAccion 

Accion ej2 es 
	Ambiente 
		total,invalidos,opcion: entero 
		

		Funcion longitud(n:entero): entero 
			Ambiente 
				num,cont: entero 
			Proceso
				num:= n; cont:= 0   
				Mientras num > 0 hacer 
					cont:= cont+1
					num:= num div 10 
				FM 
				longitud:= cont 
			FProceso 
		FFuncion 

		Funcion escaneo(x:entero): entero 
			Ambiente 
				d,num,par,impar,long: entero 
			Proceso 
				res:= x; num:= x; long:= longitud(x); par:= 0; impar:= 0; d:= 0
				si long <> 12 entonces 
					escaneo:= 0 
				Fsi 

				Mientras num > 0 hacer 
					si long mod 2 = 0 entonces 
						par:= par + num mod 10 
					sino 
						impar:= impar + num mod 10 
					Fsi 
					num:= num div 10 
					long:= long-1
				FM 

				par:= par*3
				par:= par + impar 
				d:= par mod 10 

				si d = 10 entonces 
					escaneo:= res*10 
				sino 
					escaneo:= res*10 + (10-d) 
				Fsi 
			FProceso 
		FFuncion 

		Procedimiento ingresar()
			Ambiente 
				precio,cod,codigo,cant: entero 
			Proceso 
				Esc("Ingrese el codigo: "); Leer(codigo)
				si escaneo(codigo) = 0 entonces 
					Esc("Ingrese otro codigo"); invalidos:= invalidos+1
				sino 
					cod:= escaneo(codigo) 
					precio:= ObtenerPrecio(cod)
					Esc("Ingrese la cantidad de productos: "); Leer(cant)
					si cant > 0 entonces 
						total:= total + (precio*cant)
					Fsi 
				Fsi 
			FProceso 
		FP 
	Proceso 
		total:= 0; invalidos:= 0
		Repetir 
			Esc("Ingrese una opcion: ")
			Esc("1: Ingresar código y calcular importe | 0: Finalizar")
			Leer(opcion)
			si opcion = 1 entonces 
				ingresar()
			Fsi 
		Hasta (opcion = 0)

		si total > 15000 entonces 
			total:= total - REDOND(total*0.15)
		Fsi 
		Esc("Monto final a pagar (redondeado): $",total)
		Esc("Total de codigos invalidos: ",invalidos)
	FProceso
FAccion 