Accion Ejercicio1 es 
 Ambiente 
 sec: secuencia de caracteres 
 v: caracter 
 salida: secuencia de caracteres
 t_recaudado, p_c1, p_c2, p_c3, p_c4, p_c5, c_vehiculos: entero 
 categoria, feriado: logico 
 cat, cat1: AN 
 Procedimiento total() es
	 	Segun cat1 hacer 
		 "1": t_recaudado:= t_recaudado + monto_peaje("1"); p_c1:= p_c1 + 1
		 "2": t_recaudado:= t_recaudado + monto_peaje("2"); p_c2:= p_c2 + 1
		 "3": t_recaudado:= t_recaudado + monto_peaje("3"); p_c3:= p_c3 + 1
		 "4": t_recaudado:= t_recaudado + monto_peaje("4"); p_c4:= p_c4 + 1
		 "5": t_recaudado:= t_recaudado + monto_peaje("5"); p_c5:= p_c5 + 1
		FinSegun 
	FinProcedimiento
 Procedimiento porcentaje() es
	 p_c1:= (p_c1/c_vehiculos) * 100
	 p_c2:= (p_c2/c_vehiculos) * 100
	 p_c3:= (p_c3/c_vehiculos) * 100
	 p_c4:= (p_c4/c_vehiculos) * 100
	 p_c5:= (p_c5/c_vehiculos) * 100
	FinProcedimiento
 Proceso 
	 Arr(sec)
	 Avz(sec,v)
	 Crear(salida)
	 t_recaudado:=0; c_vehiculos:=0; p_c1:=0; p_c2:=0; p_c3:=0; p_c4:=0; p_c5:=0
	 categoria:= falso; feriado:= falso 
	 Esc("Ingrese categoria")
	 Leer(cat)
	 Mientras NFDS(sec) hacer 
	 	 Avz(sec,v)
		 Avz(sec,v)
		 Mientras v <> "!" hacer
		 	 c_vehiculos:= c_vehiculos + 1 
		 	 	Si v = "S" entonces 
				 feriado:= verdadero 
				Sino 
				 feriado:= falso
				FinSi
			 Avz(sec,v)
			 cat1:= v 
			 	Si cat = cat1 entonces 
				 categoria:= verdadero
				Sino 
				 categoria:= falso 
				FinSi
			 total()
			 	Si (categoria) Y (feriado) entonces
				 Para i:= 1 a 11 hacer 
				  	 Grabar(salida,v)
					 Avz(sec,v)
					FinPara 
				Sino 
				 Para i:= 1 a 11 hacer 
				  	 Avz(sec,v)
					FinPara 
				FinSi 
			 	Si (categoria) Y (feriado) entoces 
				 Grabar(salida,"#")
				FinSi
			FinMientras
		 Esc("La cantidad en el dia es de",t_recaudado,"$")
		FinMientras
	 Cerrar(sec)
	 Cerrar(salida)
	 porcentaje()
	 Esc("Porcetajes"); Esc("c1/c2/c3/c4/c5"); Esc(p_c1,"/",p_c2,"/",p_c3,"/",p_c4,"/",p_c5)
	FinProceso 
FinAccion 


Accion Ejercicio2 es 
 Ambiente 
 PEAJE = registro 
	 anio: N(4)
	 mes: N(2)
	 dia: N(2)
	 cat: N(1)
	 patente: AN(5)
	 c_pases: N(4)
	FinRegistro 
 SALIDA = registro 
	 anio: N(4)
	 mes: N(2)
	 dia: N(2)
	 c_pases: N(4)
	FinRegistro 
 arch: archivo de PEAJE ordenado por anio, mes, dia, cat y patente
 salida: archivo de SALIDA 
 reg: PEAJE 
 sal: SALIDA 
 p_mes, p_anio, p_total, mayor: entero
 resg_anio, resg_mes, resg_dia, aaaa: AN 
 Procedimiento corte_anio() es
	 corte_mes()
	 Esc("Los pases en el año",resg_anio,"son de",p_anio)
	 p_total:= p_total + p_anio
		Si p_anio > mayor entonces 
		 mayor:= p_anio
		 aaaa:= resg_anio
		FinSi
	 p_anio:= 0 
	 resg_anio:= reg.anio 	
	FinProcedimiento
 Procedimiento corte_mes() es
	 corte_dia()
	 Esc("Los pases en el mes",resg_mes,"son de",p_mes)
	 p_anio:= p_anio + p_mes
	 p_mes:= 0 
	 resg_mes:= reg.mes 
	FinProcedimiento
 Procedimiento corte_dia() es 
	 sal.anio:= resg_anio
	 sal.mes:= resg_mes
	 sal.dia:= resg_dia
	FinProcedimiento
 Proceso 
	 Abrir E/(arch)
	 Abrir S/(salida)
	 Leer(arch,reg)
	 p_mes:=0; p_anio:=0; p_total:=0, mayor:=LV
	 resg_anio:= reg.anio 
	 resg_mes:= reg.mes 
	 resg_dia:= reg.dia 
	 Mientras NFDA(arch) hacer 
	 	 	Si resg_anio <> reg.anio entonces 
			 corte_anio()
			Sino 
				Si resg_mes <> reg.mes entonces 
				 corte_mes()
				Sino 
					Si resg_dia <> reg.dia entonces 
					 corte_dia()
					FinSi
				FinSi
			FinSi
		 p_mes:= p_mes + reg.c_pases
		 Leer(arch,reg)
		FinMientras
	 Esc("El total de pases es de",p_total)
	 Esc("El año en el que más pases hubo es de", aaaa)
	FinProceso
FinAccion

			  
			  



