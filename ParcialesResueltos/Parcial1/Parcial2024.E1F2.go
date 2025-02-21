//Secuencia de ejemplo 
12S22230AC13FE32231AB45HI42231AB21HU!13N50615AD23FH...!

1) Generar una secuencia de salida con la informacion de todos los vehiculos que han circulado los dias viernes, sabado y domingo. La funcion DiaDeLaSemana() recibe como parametro, un numero de dia, un numero de mes y un numero de año y devuelve el dia de la semana correspondiente. Se desea saber "HHMMPatente"
2) Por cada dia informar la cantidad de vehiculos que pasaron por cada categoria 
3) Cantidad de vehiculos que pasaron en dias feriados y total recaudado de esos dias 

Accion parcial es 
 Ambiente 
 sec: secuencia de caracteres 
 v: caracter 
 sal: secuencia de caracteres
 bandera, bandera2: logico
 funcion conv_entero(v:caracter):entero 
	 	segun v hacer 
	 	 ="1": conv_entero:= 1
		 ="2": conv_entero:= 2
		 ="3": conv_entero:= 3
		 ="4": conv_entero:= 4
		 ="5": conv_entero:= 5
		 ="6": conv_entero:= 6
		 ="7": conv_entero:= 7
		 ="8": conv_entero:= 8
		 ="9": conv_entero:= 9
		 ="0": conv_entero:= 0
		FinSegun
	FinFuncion
 Finde: ['viernes', 'sabado', 'domingo']
 c1, c2, c3, c4, c5, cf, total, dia, ddd: entero
 finde2: AN
 Proceso 
	 Arr(sec)
	 Avz(sec,v)
	 c1:= 0; c2:= 0; c3:= 0; c4:= 0; c5:= 0; total:= 0; dia:= 0; ddd:= 0                                                
	 bandera:= falso                                                                   
	 Mientras NFDS(sec) hacer 
		 dia:= dia + conv_entero(v) * 10 
		 Avz(sec,v)
		 dia:= dia + conv_entero(v) 
		 Avz(sec,v)
		 	Si v = "S" entonces 	      
			 bandera:= verdadero
			 cf:= cf + 1 
			FinSi
		 Mientras v <> "!" hacer
		 		Segun v hacer   
				 1: c1:= c1 + 1 
				 2: c2:= c2 + 1 
				 3: c3:= c3 + 1
				 4: c4:= c4 + 1
				 5: c5:= c5 + 1 
				FinSegun
			 	Si bandera entonces
				 	Segun v hacer 
					 1: total:= total + 900
					 2: total:= total + 1800
					 3: total:= total + 2700 
					 4: total:= total + 3600
					 5: total:= total + 4500
					FinSegun 
				FinSi
			 	Si DiaDeLaSemana(dia,mes,anio) EN Finde entonces
				 Para i:= 1 a 10 hacer 
				 	 Grabar(sal,v)
					 Avz(sec,v)
					FinPara
				 Grabar(sal,"-")
				Sino 
				 Para i:= 1 a 10 hacer 
				 	 Avz(sec,v)
					FinPara 
				FinSi
			 Esc("Autos categoria 1",c1)
			 Esc("Autos categoria 2",c2)
			 Esc("Autos categoria 3",c3)
			 Esc("Autos categoria 4",c4)
		     Avz(sec,v)
			FinMientras
		FinMientras
	 Cerrar(sec)
	 Cerrar(sal)
	 Esc("Cantidad de autos en feriado",cf)
	 Esc("Monto recaudado",total)
	FinProceso 
FinAccion
	 

Accion parcial2 es 
 Ambiente 
 PEAJE = registro 
	 clave = registro 
		 anio: N(4)
		 mes: N(2)
		 dia: N(3) 
		 categoria: N(1)
		 patente: AN(5)
		FinRegistro
	 c_pases: N(5)
	FinRegistro 
 Salida = registro 
	 anio: N(4)
	 mes: N(2)
	 dia: N(3)
	 patente: AN(5)
	 c_pases: N(5)
	FinRegistro
 arch: archivo de PEAJE ordenado por clave
 reg: PEAJE
 salida: archivo de Salida
 sal: Salida
 v_dia, v_mes, v_anio, ddd, mmm, aaaa, total: entero 
 res_anio, res_mes, res_dia: AN
 Procedimiento corte_anio() es 
	 corte_mes()
	 Esc("Cantidad en el anio",resg_anio,"es de",v_anio)
	 total:= total + v_anio
	 v_anio:= 0 
	 resg_anio:= reg.anio
	FinProcedimiento 
 Procedimiento corte_mes() es
	 corte_dia()
	 Esc("Cantidad en el mes",resg_mes,"es de",v_mes)
	 v_anio:= v_anio + v_mes 
	 v_mes:= 0 
	 resg_mes:= reg.mes 
	FinProcedimiento
 Procedimiento corte_dia() es 
	 Esc("Cantidad en el dia",res_dia,"es de",v_dia)
	 v_mes:= v_mes + v_dia 
	 v_dia:= 0 
	 resg_dia:= reg.dia
	FinProcedimiento
 Proceso 
	 Abrir E/(arch)
	 Abrir S/(Salida)
	 Leer(arch, reg)
	 v_dia:= 0; v_mes:= 0; v_anio:= 0; ddd:= 0; mmm:= 0; aaaa:= 0; total:= 0
	 res_anio:= reg.anio
	 res_mes:= reg.mes 
	 res_dia:= reg.dia
	 Esc("/dia/ /mes/ /anio/")
	 Leer(ddd, mmm, aaaa)
	 Mientras NFDA(arch) hacer 
	 	 	Si res_anio <> reg.anio entonces 
			 corte_anio()
			Sino 
				Si res_mes <> reg.mes entonces
				 corte_mes()
				Sino 
					Si res_dia <> reg.dia entonces 
					 corte_dia()
					FinSi
				FinSi
			FinSi
		 	Si reg.c_pases > 2 entonces
			 v_dia:= v_dia + reg.c_pases
			FinSi
			Si (ddd = reg.dia) Y (mmm = reg.mes) Y (aaaa = reg.anio) entonces
			 sal.anio:= reg.anio 
			 sal.mes:= reg.mes 
			 sal.dia:= reg.dia 
			 sal.categoria:= reg.categoria
			 sal.patente:= reg.patente 
			 Grabar(Salida, sal)
			FinSi
		 Leer(arch, reg)
		FinMientras
	 Cerrar(arch)
	 Cerrar(Salida)
	 Esc("Total",total)
	FinProceso 
FinAccion
	 

 



 


