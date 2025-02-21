Accion ejercicio1 es 
 Ambiente 
 res: secuencia de caracteres 
 v1: caracter 
 adop: secuencia de caracteres
 v2: caracter
 salida: secuencia de caracteres
 numeros: secuencia de enteros 
 Funcion conv_entero(v:caracter): entero 
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
		FinSegun
	FinFuncion
 canino, adoptado: logico 
 p_noadop, c_noadop, noadop, id1, id2: entero
 tipo: AN
 Proceso 
	 Arr(res)
	 Arr(adop)
	 Avz(res,v1)
	 Avz(adop,v2)
	 canino:=falso; adoptado:=falso
	 p_noadop:=0; c_noadop:=0; noadop:=0; id1:=0; id2:=0
	 Mientras NFDS(res) Y NFDS(adop) hacer 
	 	 id1:= conv_entero(v1)*10; Avz(res,v1)
		 id1:= id1 + conv_entero(v1); Avz(res,v1)
		 id2:= conv_entero(v2)*10; Avz(adop,v2)
		 id2:= id2 + conv_entero(v2)
		 tipo:= v1
		 	Si (id1 = id2) entonces
			 adoptado:= verdadero 
			Sino 
			 adoptado:= falso 
			FinSi
		 Mientras (v1 <> "*") Y (v2 <> "*") hacer 
		  		Si (adoptado) entonces 
					Si tipo = "C" entonces 
					 Mientras v1 <> "#" hacer 
					 	 Grabar(salida,v1)
						 Avz(res,v1)
						FinMientras 
					 Para i:= 1 hasta 3 hacer 
					 	 Avz(res,v1)
						FinPara
					 Mientras v2 <> "%" hacer 
					 	 Avz(adop,v2)
						FinMientras
					 Avz(adop,v2)
					 Para i:= 1 hasta 10 hacer 
					  	 Grabar(salida,v2)
						FinPara 
					 Grabar(salida,"$")
					FinSi
				Sino 
					Si tipo = "C" entonces 
					 p_noadop:= p_noadop + 1 
					Sino 
					 c_noadop:= c_noadop + 1 
					FinSi
				 noadop:= noadop + (c_noadop + p_noadop)
				 Grabar(numero,v1)
				 Mientras (v1 <> "#") Y (v2 <> "%") hacer 
				 	 Avz(res,v1)
					 Avz(adop,v2)
					FinMientras
				 Para i:= 1 hasta 3 hacer 
				 	 Avz(res,v1)
					FinPara 
				 Para i:= 1 hasta 10 hacer 
				 	 Avz(adop,v2)
					FinPara
				FinSi
			FinMientras
		FinMientras
	 p_noadop:= (p_noadop/noadop) * 100
	 c_noadop:= (c_noadop/noadop) * 100
	 Esc("El porcentaje de gatos no adoptados es de",c_noadop,"%")
	 Esc("El porcentaje de perros no adoptados es de",p_noadop,"%")
	FinProceso
FinAccion

Accion ejercicio2 es 
 Ambiente 
 ADOPTADOS = registro 
	 clave = registro 
		 departamento: AN(20)
		 localidad: AN(20)
		 barrio: AN(20)
		 dni: N(8)
		 id_masc: N(2)
		FinRegistro
	 tipo: AN(1)
	 c_anteriores: N(1)
	FinRegistro
 SALIDA = registro 
	 departamento: AN(20)
	 localidad: AN(20)
	 c_adopciones: N(2)
	FinRegistro
 arch: archivo de ADOPTADOS ordenado por clave 
 reg: ADOPTADOS
 salida: archivo de SALIDA 
 sal: SALIDA
 p_barrio, p_localidad, g_barrio, g_localidad, total_g: entero 
 resg_departamento, resg_localidad, resg_barrio: AN
 Procedimiento corte_departamento() es 
	FinProcedimiento 
 Procedimiento corte_localidad() es 
	 Esc("En la localidad",resg_localidad,"la cantidad de perros es",p_localidad)
	 Esc("En la localidad",resg_localidad,"la cantidad de gatos es",g_localidad)
	 total:= p_localidad + g_localidad
	 totaal_g:= totaal_g + g_localidad
	 g_localidad:= 0
	 p_localidad:= 0
	 resg_localidad:= reg.localidad 
	FinProcedimiento
 Procedimiento corte_barrio() es 
	 Esc("En el barrio",resg_barrio,"la cantidad de perros es",p_barrio)
	 Esc("En el barrio",resg_barrio,"la cantidad de gatos es",g_barrio)
	 p_localidad:= p_localidad + p_barrio
	 g_localidad:= g_localidad + g_barrio
	 p_barrio:= 0
	 g_barrio:= 0 
	 resg_barrio:0 reg.barrio
	FinProcedimiento
 Proceso 
	 Abrir E/(arch)
	 Abrir S/(salida)
	 Leer(arch,reg)
	 p_barrio:=0; p_localidad:=0; g_barrio:=0; g_localidad:=0; totaal_g:=0
	 resg_barrio:= reg.barrio
	 resg_localidad:= reg.localidad
	 resg_departamento:= reg.departamento 
	 Mientras NFDA(arch) hacer 
	 	 	Si resg_departamento <> reg.departamento entonces 
			 corte_departamento()
			Sino 
				Si resg_localidad <> reg.localidad entonces 
				 corte_localidad()
				Sino 
				 	Si resg_barrio <> reg.barrio entonces 
					 corte_barrio()
					FinSi
				FinSi
			FinSi
		 	Si reg.tipo = "C" entonces 
			 p_barrio:= p_barrio + 1 
			Sino 
			 g_barrio:= g_barrio + 1 
			FinSi
			Si reg.c_anteriores > 3 entonces 
			 sal.departamento:= resg_departamento
			 sal.localidad:= resg_localidad
			 sal.c_adopciones:=
		FinMientras
	 	
				  



