Accion Ejercicio1 (pokemon: arreglo[1...151] de AN) es
 Ambiente 
 Fecha = registro 
	 dia: N(2)
	 mes: N(2)
	 anio: N(4)
	FinRegistro 
 USUARIOS = registro
	 clave = registro 
		 cod_region: N(4)
		 cod_usu: N(10)
		FinRegistro
	 correo: AN(50)
	 experiencia: N(7)
	FinRegistro
 arch_u, salida: archivo de USUARIOS ordenado por clave 
 reg_u, sal: USUARIOS
 CAPTURAS = registro 
	 clave = registro
		 cod_region: N(4)
		 cod_usu: N(10)
		 cod_pokemon: 1...151
		FinRegistro
	 puntos_exp: N(5)
	 f_captura_ Fecha
	 estado_pok: AN(1)
	 estado_usu: AN(1)
	FinRegistro
 arch_c: archivo de CAPTURAS ordenado por clave 
 reg_c: CAPTURAS
 poke: arreglo[1...151] de enteros 
 i, mayor, poke_may: enteros
 Procedimiento sumar_puntos() es 
		Si reg_c.estado_pok = "E" entonces 
		 aux.experiencia:= aux.experiencia + (reg_c.puntos_exp)*2
		Sino 
		 aux.experiencia:= aux.experiencia + reg_c.puntos_exp
		FinSi 
	FinProcedimiento 
 Procedimiento alta_usuario() es 
		Si estado_usu = "A" entonces
		 aux.clave.cod_region:= reg_c.clave.cod_region
		 aux.clave.cod_usu:= reg_c.clave.cod_usu
		 aux.correo:= " "
		 aux.experiencia:= sal.experiencia + reg_c.puntos_exp
		 Grabar(salida,sal)
		FinSi
	FinProcedimiento
 Procedimiento proceso_igual() es 
		Si estado_usu = "S" entonces
		 Leer(arch_c,reg_c)
		FinSi
	FinProcedimiento
 Procedimiento leer_usu() es 
	 Leer(arch_u,reg_u)
	 	Si FDA(arch_u) entonces
		 reg_u.clave:= HV 
		FinSi
	FinProcedimiento
 Procedimiento leer_cap() es 
	 Leer(arch_c,reg_c)
	 	Si FDA(arch_c) entonces
		 reg_c.clave:= HV
		FinSi
	FinProcedimiento
 Procedimiento contar_poke() es 
	 	Si estado_pok = "D" entonces
		 i:= reg_c.cod_pokemon 
		 poke[i]:= poke[i] + 1 
		FinSi
	FinProcedimiento
 Proceso 
	 Abrir E/(arch_c)
	 Abrir E/(arch_u)
	 leer_cap()
	 leer_usu()
	 mayor:= 0; poke_may:=0 
	 Para i:= 1 a 151 hacer 
	 	 poke[i]:= 0 
		FinPara 
	 Mientras (reg_c.clave <> HV) Y (reg_u.clave <> HV) hacer 
	 	 	Si reg_u.clave < reg_c.clave entonces 
			 sal:= reg_u
			 Grabar(salida,sal)
			 leer_usu()
			Sino 
				Si reg_u.clave = reg_c.clave entonces 
				 aux:= reg_u
				 Mientras aux.clave = reg_c.clave hacer 
				 	 proceso_igual()
					 sumar_puntos()
					 leer_cap()
					FinMientras 
				 sal:= aux 
				 contar_poke()
				 Grabar(salida,sal)
				 leer_usu()
				Sino 
					Si reg_u.clave > reg_c.clave entonces 
					 alta_usuario()
					 aux:= reg_u
					 Mientras aux.clave = reg_c.clave hacer 
					 	 proceso_igual()
						 sumar_puntos()
						 leer_cap()
						FinMientras
					 sal:= aux
					 contar_poke()
					 Grabar(salida,sal)
					 leer_usu()
					FinSi
				FinSi
			FinSi
		FinMientras
	 Para i:= 1 a 151 hacer 
	 	 	Si mayor < poke[i] entonces
			 mayor:= poke[i] 
			 poke_may:= i 
			FinSi
		FinPara 	
	 Esc("El pokemon que mas entrenadores tiene es",A[poke_may],"con",mayor,"entrenadores")
	 Cerrar(arch_c)
	 Cerrar(arch_u)
	 Cerrar(salida)
	FinProceso 
FinAccion


		  


	
 