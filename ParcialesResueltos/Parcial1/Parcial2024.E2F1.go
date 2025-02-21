Accion Ejercicio1 es 
 Ambiente 
 sec1: secuencia de caracteres 
 v1: caracter 
 sec2: secuencia de caracteres
 v2: caracter
 salida: secuencia de caracteres
 p1, p2, p3, p4, p5, cat: AN 
 c1, c2, c3, c4, c5, mes, mes1, c_vehiculos: entero 
 descuento: logico
 Funcion conv_entero(v: caracter): entero 
		Segun v hacer 
		 "1": conv_entero:= 1
		 "2": conv_entero:= 2
		 "3": conv_entero:= 3
		 "4": conv_entero:= 4
		 "5": conv_entero:= 5
		 "6": conv_entero:= 6
		 "7": conv_entero:= 7
		 "8": conv_entero:= 8
		 "9": conv_entero:= 9
		 "0": conv_entero:= 0
		FinSegun 
	FinFuncion 
 Procedimiento conseguir_valor() es 
	 	Segun v hacer 
		 "1": c1:= c1 + obtener_valor(cat) 
		 "2": c2:= c2 + obtener_valor(cat)
		 "3": c3:= c3 + obtener_valor(cat)
		 "4": c4:= c4 + obtener_valor(cat)
		 "5": c5:= c5 + obtener_valor(cat)  
		FinSegun
	FinProcedimiento
 Proceso 
	 Arr(sec1)
	 Arr(sec2)
	 Avz(sec1,v1)
	 Avz(sec2,v2)
	 Crear(salida)
	 c1:= 0; c2:= 0; c3:=0; c4:=0; c5:=0; mes:=0; c_vehiculos:=0; mes1:=0
	 descuento:= falso 
	 Esc("Ingrese mes")
	 Leer(mes)
	 Mientras NFDS(sec1) Y NFDS(sec2) hacer 
	 	 Avz(sec1,v1)
		 Avz(sec1,v1)
		 mes1:= mes1 + conv_entero(v)*10
		 Avz(sec1,v1)
		 mes1:= mes1 + conv_entero(v)
		 Para i:= 1 a 5 hacer
		 	 Avz(sec1,v1)
			FinPara 
		 Mientras v2 <> "-" hacer 
		 	 cat:= v 
			 c_vehiculos:= c_vehiculos + 1
			 Avz(sec2)
			 	Si v2 = "S" entonces 
				 descuento:= verdadero 
				FinSi
			 Para i:= 1 a 4 hacer 
			 	 Avz(sec2,v2)
				FinPara 
			 	Si descuento Y (mes = mes1) entonces 
			 	 p1:= v2 
				 Avz(sec2,v2)
				 Grabar(salida, p1)
				 p2:= v2
				 Avz(sec2,v2)
				 Grabar(salida, p2)
				 p3:= v2
				 Avz(sec2,v2)
				 Grabar(salida, p3)
				 p4:= v2
				 Avz(sec2,v2)
				 Grabar(salida, p4)
				 p5:= v2
				 Avz(sec2,v2)
				 Grabar(salida, p5)
				FinSi
			 conseguir_valor()
			 	Si descuento Y (mes = mes1) entonces 
				 Grabar(salida,"?")
				FinSi
			 mes:= 0 
			 mes1:= 0
			 Avz(sec2,v2)
			FinMientras 
		FinMientras
	 Esc("El total recaudado por categoria es de"); Esc("Categoria 1:",c1,"2:",c2,"3:",c3,"4:",c4,"5:",c5)
	 Esc("La cantidad de vehiculos que pasaron en total es de",c_vehiculos)
	FinProceso 
FinAccion


Accion Ejercicio2 es 
 Ambiente 
 PEAJE = registro 
	 anio: N(4)
	 mes: N(2)
	 dia: N(2)
	 categoria: AN(1)
	 patente: AN(5)
	 origen: ("Rcia" o "Ctes")
	FinRegistro 
 Salida = registro
	 anio: N(4)
	 mes: N(2)
	 c_resistencia: N(4)
	 c_corrientes: N(4)
	FinRegistro
 arch: archivo de PEAJE 
 reg: PEAJE 
 salida: archivo de Salida 
 sal: Salida 
 resg_anio, resg_mes: AN 
 c_mes_resis, c_mes_corr, c_anio_resis, c_anio_corr, total: entero 
 Procedimiento corte_anio() es 
	 corte_mes()
	 Esc("Desde resistencia en el anio",resg_anio,"es de",c_anio_resis)
	 Esc("Desde resistencia en el anio",resg_anio,"es de",c_anio_corr)
	 c_anio_corr:= 0 
	 c_anio_resis:= 0 
	 resg_anio:= reg.anio
	FinProcedimiento
 Procedimiento corte_mes() es 
	 Esc("Desde resistencia en el mes",resg_mes,"es de",c_mes_resis,)
	 Esc("Desde Corrientes en el mes",resg_mes,"es de",c_mes_corr)
	 c_anio_corr:= c_anio_corr + c_mes_corr 
	 c_anio_resis:= c_anio_resis + c_mes_resis
	 total:= total + (c_mes_corr + c_mes_resis)
	 	Si total > 100000 entonces 
		 sal.anio:= resg_anio
		 sal.mes:= resg_mes
		 sal.c_resistencia:= c_mes_resis
		 sal.c_corrientes:= c_mes_corr
		FinSi
	 c_mes_corr:= 0 
	 c_mes_resis:= 0 
	 total:= 0 
	 resg_mes:= reg.mes
	FinProcedimiento
 Proceso
	 Abrir E/(arch)
	 Leer(arch,reg)
	 Abrir S/(salida)
	 resg_anio:= reg.anio 
	 resg_mes:= reg.mes
	 c_anio_corr:=0; c_mes_corr:=0; c_mes_resis:=0; c_mes_corr:=0; total:=0
	 Mientras NFDA(arch) hacer 
	 	 	Si resg_anio <> reg.anio entonces
			 corte_anio()
			Sino 
				Si resg_mes <> reg.mes entonces 
				 corte_mes()
				FinSi
			FinSi
		 	Si reg.origen = "Rcia" entonces 
			 c_mes_resis:= c_mes_resis + 1 
			Sino 
			 c_mes_corr:= c_mes_corr + 1 
			FinSi
		 Leer(arch,reg)
		FinMientras
	 Cerrar(arch)
	 Cerrar(salida)
	FinProceso 
FinAccion




	 

	

	