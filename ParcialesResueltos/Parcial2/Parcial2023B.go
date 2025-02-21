Accion Ejercicio1 (A: arreglo[1...2,1...6] ) es 
 Ambiente
 Fecha = registro 
	 dia: N(2)
	 mes: N(2)
	 anio: N(4)
	FinRegistro 
 BICICLETAS = registro 
	 clave = registro 
		 nro_serie: N(5)
		 modelo: AN(5)
		FinRegistro
	 f_adquisicion: Fecha
	 f_mantenimiento: Fecha
	FinRegistro
 arch, arch_act: archivo de BICICLETAS ordenado por clave 
 reg, reg_act, aux: BICICLETAS
 NOVEDADES = registro 
	 clave = registro 
	 	 nro_serie: N(5)
		 modelo: AN(5)
		FinRegistro
	 t_novedad: (1,2,3,4)
	 f_novedad: Fecha 
	 h_inicio: N(2)
	 h_fin: N(2)
	 circuito_nro: N(1)
	 id_usuario: AN(10)
	FinRegistro
 arch_nov: archivo de NOVEDADES ordenado por clave 
 reg_nov: NOVEDADES
 Procedimiento leer_bici() es 
	 Leer(arch,reg)
	 	Si FDA(arch) hacer
		 reg.clave:= HV
		FinSi 
	FinProcedimiento 
 Procedimiento leer_nov() es 
	 Leer(arch_nov,reg_nov)
	 	Si FDA(arch_nov) entonces 
		 reg_nov.clave:= HV
		FinSi
	FinProcedimiento
 cir_usu, paseo_usu, diferenciar, valor: entero 
 usuario: arreglo[1...2,1...6] de Proceso  
 Proceso = registro 
	 monto: N(6)
	 paseos: N(6)
	FinRegistro
 Procedimiento procesos_iguales() es 
	 	Si reg_nov.t_novedad = 1 entonces 
		 Esc("a")
		Sino 
			Si reg_nov.t_novedad = 2 entonces 
			 diferenciar:= diff_horas(reg_nov.h_inicio, reg_nov.h_fin)
			 j:= reg_nov.circuito_nro
			 Segun diferenciar hacer 
			 	 > 6: i:= 1; valor:= 1500
				 < 6: i:= 2; valor:= 1000
				FinSegun 
			 usuario[i,j].monto:= usuario[i,j] + (valor*A[i,j])
			 usuario[i,j].paseos:= usuario[i,j].paseos + 1  
			Sino 
				Si reg_nov.t_novedad = 3 entonces
				 aux.f_mantenimiento:= fecha_actual()
				Sino 
				 Esc("Baja de la unidad")
				FinSi
			FinSi
		FinSi
	FinProcedimiento
 Proceso 
	 Abrir E/(arch)
	 Abrir E/(arch_nov)
	 Abrir S/(arch_act)
	 leer_bici()
	 leer_nov()
	 Para i:= 1 a 2 hacer
	  	 Para j:= 1 a 6 hacer
			 usuario[i,j]:= 0 
			FinPara
		FinPara
	 diferenciar:=0; valor:=0; cir_usu:=0; paseo_usu:=0
	 Esc("Ingrese circuito y paseo")
	 Leer(cir_usu); Leer(paseo_usu)
	 Mientras (reg.clave <> HV) o (reg_nov.clave <> HV) hacer 
	 	 	Si reg.clave < reg_nov.clave entonces 
			 reg_act:= reg 
			 Grabar(arch_act,reg_act)
			 leer_bici()
			Sino 
				Si reg.clave = reg_nov.clave entonces 
				 aux:= reg 
				 Mientras aux.clave = reg_nov.clave entonces 
				 	 procesos_iguales()
					 leer_nov()
					FinMientras 
				 reg_act:= aux 
				 Grabar(arch_act,reg_act)
				 leer_bici()
				Sino
					Si reg_nov.t_novedad = 1 entonces
					 aux.nro_serie:= reg_nov.nro_serie
					 aux.modelo:= reg_nov.modelo
					 aux.f_adquisicion:= " "
					 aux.f_mantenimiento:= " "
					Sino 
					 Esc("Error")
					FinSi
				 leer_nov()
				 Mientras aux.clave = reg_nov.clave hacer
				 	 procesos_iguales()
					 leer_nov()
					FinMientras
				 reg_act:= aux
				 Grabar(arch_act,reg_act)
				FinSi
			FinSi
		FinMientras
	 i:= paseo_usu
	 j:= cir_usu 
	 Esc("El total de prestamos es",usuario[i,j].paseos,"y el monto es",usuario[i,j].monto)
	 Cerrar(arch)
	 Cerrar(arch_act)
	 Cerrar(arch_nov)
	FinProceso 
FinAccion

Accion Ejercicio2 es 
 Ambiente
 NOVEDADES = registro 
 	 clave = registro 
		  nro_serie: N(5)
		 modelo: AN(5)
		FinRegistro
	 t_novedad: (1,2,3,4)
	 f_novedad: Fecha 
	 h_inicio: N(2)
	 h_fin: N(2)
	 circuito_nro: N(1)
	 id_usuario: AN(10)
	FinRegistro
 arch_nov: archivo de NOVEDADES ordenado por clave 
 reg_nov: NOVEDADES
 USUARIOS = registro 
	 id_usuario: AN(10)
	 dni: N(8)
	 sexo: ("M", "F")
	 apeynom: AN(30)
	 domicilio: AN(30)
	 localidad: AN(30)
	 provincia: AN(30)
	 edad: N(2)
	FinRegistro
 Funcion obtener_rango(x: entero): AN 
	 Segun x hacer 
		 1: Esc("menor 18")
		 2:	Esc("entre 18 y 35")
		 3: Esc("entre 35 y 75")
		 4: Esc("mayor a 75")		
		FinSegun
	FinFuncion
	Procedimiento edad() es 
		Si reg_usu.edad < 18 entonces 
	 	 j:=1
		Sino 
			Si (reg_usu.edad > 18) Y (reg_usu < 35) entonces
			 j:=2
			Sino 
				Si (reg_usu.edad > 35) Y (reg_usu.edad < 75) entonces
				 j:=3
				Sino 
				 j:=4
				FinSi
			FinSi
		FinSi
	FinProcedimiento
 arch_usu: archivo de USUARIOS indexado por id_usuario
 reg_usu: USUARIOS
 etario: arreglo[1...3,1...5] de enteros
 mayor, rango_may, i, j, diferenciar: entero
 Proceso 
	 Abrir E/(arch_nov)
	 Abrir E/(arch_usu)
	 Leer(arch_nov,reg_nov)
	 Para i:= 1 a 3 hacer 	
	 	 Para j:= 1 a 4 hacer 
		 	 usuario[i,j]:= 0 
			FinPara
		FinPara
	 mayor:=0; rango_may:=0
	 Mientras NFDA(arch_nov) hacer
		 bandera:= falso
	 		Si reg_nov.t_novedad = 2 entonces 
	 		 reg_usu.id_usuario:= reg_nov.id_usuario
			 Leer(arch_usu,reg_usu)
			 	Si EXISTE entonces
				 edad() 
				 bandera:= verdadero
				FinSi
			 	Si bandera entonces 
				 diferenciar:= diff_horas(reg_nov.h_inicio,reg_nov.h_fin)
				 Segun diferenciar hacer 
				 	 > 6: i:=1
					 < 6: i:=2
					FinSegun
				 etario[i,j]:= etario[i,j] + 1 
				 etario[3,j]:= etario[3,j] + 1 
				 etario[i,5]:= etario[i,5] + 1
				FinSi
			FinSi
		 Leer(arch_nov,reg_nov)
		FinMientras
	 Para i:= 1 a 2 hacer 
	     Para j:= 1 a 5 hacer
		  		Si etario[2,j] > mayor entonces  
				 mayor:= etario[2,j]
				 rango_may:= j 
				FinSi
			FinPara
		FinPara
	 Esc("La cantidad de prestamos intensivos para mayores de 75 años es",etario[3,5])
	 Esc("El rango etario que mas paseos recreativos realizo es",obtener_rango(rango_may))
	 Cerrar(arch_usu)
	 Cerrar(arch_nov)
	FinProceso
FinAccion

				 
			
			 
	 	 

	

 


 
			 
		 
		

