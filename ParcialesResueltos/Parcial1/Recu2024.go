Accion Ejercicio1 es 
 Ambiente
 histo: secuencia de caracteres 
 v1: caracter 
 info: secuencia de caracteres 
 v2: caracter 
 salida: secuencia de caracteres
 historial, vocales: logico 
 vocal: ("A","E","I","O","U")
 c_mayor, c_menor, total, edad, edad1: entero 
 Procedimiento mayor_menor() es 
		Si edad > edad1 entonces 
		 c_mayor:= c_mayor + 1 
		Sino 
		 c_menor:= c_menor + 1 
		FinSi
	FinProcedimiento
 Proceso 
	 Arr(histo)
	 Arr(info)
	 Crear(salida)
	 Avz(histo,v1)
	 Avz(info,v2)
	 c_mayor:=0; c_menor:=0; total:=0; edad:=0; edad1:=0 
	 historial:= falso; vocales:= falso 
	 Mientras NFDS(info) Y NFDS(histo) entonces 
		 	Si v1 = "P" entonces 
			 historial:= verdadero 
			Sino 
			 historial:= falso 
			FinSi
			Avz(histo,v1)
		 Mientras v <> "@" hacer 
	 			Si v2 EN vocal entonces 
				 vocales:= verdadero
				Sino 
				 vocales: falso 
				FinSi
				 Avz(info,v2)
				 Avz(info,v2)
		 	 	Si historial entonces 
				 	Si vocales entonces 
					 Mientras v1 <> "#" hacer 
					 	 Grabar(salida, v1)
						 Avz(histo, v1)
						FinMientras
					 Grabar(salida,"#") 
					 edad1:= conv_entero(v2)*10; Avz(histo,v2)
					 edad1:= edad1 + conv_entero(v2); Avz(histo,v2)
					 mayor_menor()  
					 Mientras v2 <> "#" hacer 
					 	 Grabar(salida,v2)
						 Avz(info,v2)
						FinMientras
					 Grabar(salida,"%")
					 Avz(info,v2)
					Sino 
					 edad1:= conv_entero(v2)*10; Avz(histo,v2)
					 edad1:= edad1 + conv_entero(v2); Avz(histo,v2)
					 mayor_menor()  
					 Mientras v2 <> "#" hacer 
						 Avz(info,v2)
						FinMientras
					FinSi 
				Sino 
				 Mientras v1 <> "#" hacer 
				  	 Avz(histo,v1)
					FinMientras
				 Repetir 
				 	 Avz(info,v2)
					Hasta que v2 = "@"
				FinSi





Accion Ejercicio2 es 
 Ambiente 
 Fecha = registro 
	 dia: N(2)
	 mes: N(2)
	 anio: N(4)
	FinRegistro
 ATLETAS = registro 
	 clave = registro 
	 	 deleg: AN(30)
		 depor: AN(3)
		 nro_credencial: N(4)
		 apeynom: AN(30)
		FinRegistro 
	 partic: N(3)
	 fecha_nac: Fecha 
	FinRegistro
 arch: archivo de ATLETAS ordenado por clave
 reg_ ATLETAS
 SALIDA = registro 
	 deleg: AN(30)
	 c_men25: N(5)
	 c_may25: N(5)
	FinRegistro
 A_salida: archivo de SALIDA
 sal: SALIDA
 t_depor, t_deleg, t_may_depor, t_men_depor, t_may_deleg, t_men_deleg, t_may, t_men, total: entero 
 res_depor, res_deleg: AN 
 Procedimiento corte_deleg() es 
	 corte_depor()
	 Esc("En la delegacion",res_deleg,"hay",t_deleg,"atletas")
	 sal.delegacion:= res_deleg
	 sal.c_men25:= t_men_deleg
	 sal.c_may25:= t_may_deleg
	 Grabar(salida,sal)
	 t_may:= t_may + t_may_deleg
	 t_men:= t_men + t_men_deleg
	 t_may_deleg:= 0
	 t_men_deleg:= 0 
	 t_deleg:= 0 
	 res_deleg:= reg.deleg
	FinProcedimiento
 Procedimiento corte_depor() es 
	 Esc("En la delegacion",res_deleg,"en el deporte",res_depor,"hay",t_depor,"atletas")
	 t_deleg:= t_deleg + t_depor
	 t_may_deleg:= t_may_deleg + t_may_depor
	 t_men_deleg:= t_men_deleg + t_men_depor
	 t_depor:= 0 
	 t_may_depor:= 0 
	 t_men_depor:= 0 
	 res_depor:= reg.clave.depor  
	FinProcedimiento
 Proceso 
	 Abrir E/(arch)
	 Abrir S/(salida)
	 Leer(arch,reg)
	 t_depor:=0; t_deleg:=0; t_may_deleg:=0; t_men_deleg:=0; t_may_depor:=0; t_men_depor:=0; t_may:=0; t_men:=0; total:=0
	 res_depor:= reg.clave.depor 
	 res_deleg:= reg.clave.deleg
	 Mientras NFDA(arhc) hacer 
	 	 	Si res_deleg <> reg.clave.deleg entonces 
			 corte_deleg()
			Sino 
			 	Si corte_depor() <> reg.clave.depor entonces
				 corte_depor()
				FinSi
			FinSi
		 t_depor:= t_depor + 1 
		 	Si calcular_edad(reg.fecha_nac) > 25 entonces 
			 t_may_depor:= t_may_depor + 1 
			Sino 
			 t_men_depor:= t_men_depor + 1 
			FinSi
		 Leer(arch,reg)
		FinMientras
	 total:= t_may + t_men
	 Esc("Jugadores menores",t_men)
	 Esc("Jugadores mayores",t_may)
	 Esc("Total",total)
	 Cerrar(arch)
	 Cerrar(salida)
	FinProceso 
FinAccion


			  
				 
				 
					
					 




 
